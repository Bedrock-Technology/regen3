package scanner

import (
	"errors"
	"fmt"
	"github.com/Bedrock-Technology/regen3/crons"
	"time"

	"github.com/Bedrock-Technology/regen3/beaconClient"
	"github.com/Bedrock-Technology/regen3/blockTimer"
	"github.com/Bedrock-Technology/regen3/config"
	"github.com/Bedrock-Technology/regen3/keyAgentRpc"
	"github.com/Bedrock-Technology/regen3/models"
	"github.com/attestantio/go-eth2-client/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Scanner struct {
	Config         *config.Config
	DBEngine       *gorm.DB
	EthClient      *ethclient.Client
	BeaconClient   *beaconClient.Client
	BlockTimer     *blockTimer.BlockTimer
	Pods           map[string]models.Pod
	FilterAddress  []common.Address
	KeyAgentClient *keyAgentRpc.Client
	Cron           *crons.ScanCron
	Quit           chan struct{}
}

func New(config *config.Config, quit chan struct{}) *Scanner {
	scanner := &Scanner{Config: config, Quit: quit}

	var err error
	scanner.DBEngine, err = gorm.Open(mysql.Open(config.MysqlDsn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic(fmt.Sprintf("Open Mysql err:%v", err))
	}

	scanner.EthClient, err = ethclient.Dial(config.EthClient)
	if err != nil {
		panic(fmt.Sprintf("Open EthClient err:%v", err))
	}

	scanner.BeaconClient, err = beaconClient.NewClient(config.BeaconClient)
	if err != nil {
		panic(fmt.Sprintf("Open BeaconClient err:%v", err))
	}

	scanner.Pods, err = models.GetAllPods(scanner.DBEngine)
	if err != nil {
		panic(fmt.Sprintf("Get all Pods err:%v", err))
	}
	logrus.Info("get pods length:", len(scanner.Pods))

	scanner.FilterAddress = scanner.fillFilterAddress()
	logrus.Info("get filter address: ", scanner.FilterAddress)

	scanner.initBlockTimers()
	scanner.KeyAgentClient = keyAgentRpc.NewClient(scanner.Config.KeyAgent)

	scanner.Cron = crons.New()
	if err := scanner.initCrontabs(); err != nil {
		panic(fmt.Sprintf("initCrontabs err:%v", err))
	}

	return scanner
}

func (s *Scanner) initBlockTimers() {
	cursor, err := models.GetCursor(s.DBEngine, models.Scanner)
	if err != nil {
		logrus.Errorln("GetCursor:", err)
		//panic("get cursor err")
	}
	s.BlockTimer = blockTimer.NewBlockTimer(cursor.Slot)

	if s.Config.CheckVerifyWithdrawCredential.Enable {
		s.InitVerifyWithdrawCredential()
	}

	if s.Config.CheckQueueWithdraw.Enable {
		s.InitQueueWithdraw()
	}

	if s.Config.CheckStartCheckPoint.Enable {
		s.InitStartCheckPoint()
	}

	if s.Config.CheckVerifyCheckPoint.Enable {
		s.InitVerifyCheckPoint()
	}
}

func (s *Scanner) initCrontabs() error {
	if s.Config.ReportSpec != "" {
		reportSpec := &ReportSpec{Scanner: s, LastReportTime: time.Now().UTC()}
		if _, err := s.Cron.AddSpec(s.Config.ReportSpec, reportSpec); err != nil {
			return err
		}
		logrus.Info("Add ReportSpec")
	}
	return nil
}

func (s *Scanner) fillFilterAddress() []common.Address {
	address := make([]common.Address, 0, len(s.Pods)+2)
	for _, pod := range s.Pods {
		address = append(address, common.HexToAddress(pod.Address))
	}
	address = append(address,
		common.HexToAddress(s.Config.EigenDelegationManagerContract),
	)
	return address
}

func (s *Scanner) Scan() {
	scanTicker := time.NewTicker(12 * time.Second)
	defer scanTicker.Stop()

	s.Cron.Start()

	for {
		select {
		case <-scanTicker.C:
			s.scan()
		case <-s.Quit:
			logrus.Info("Quit")
			beaconClient.Cancel()
			context := s.Cron.Stop()
			select {
			case <-context.Done():
				logrus.Infoln("Quit Done")
			}
			return
		}
	}
}

func (s *Scanner) scan() {
	end, err := s.BeaconClient.GetLatestSlotNumber()
	if err != nil {
		logrus.Errorln("Get Beacon Latest Slot Number err:", err)
		return
	}

	cursor, err := models.GetCursor(s.DBEngine, models.Scanner)
	if err != nil {
		logrus.Errorln("GetCursor:", err)
		return
	}

	DbTransaction := s.DBEngine.Begin()
	start := cursor.Slot
	logrus.Infof("Scan [%d:%d)", start, end)

	realEnd, err := s.scanSlotAndBlock(start, end, DbTransaction)
	if err != nil {
		logrus.Errorf("scanSlotAndBlock [%d:%d),error:%v", start, end, err)
		DbTransaction.Rollback()
		return
	}

	cursor.Slot = realEnd
	if err := models.SaveCursor(DbTransaction, cursor); err != nil {
		logrus.Errorln("UpdateCursor error:", err)
		DbTransaction.Rollback()
		return
	}

	if rest := DbTransaction.Commit(); rest.Error != nil {
		logrus.Errorln("DbTransaction error:", rest.Error)
		return
	}

	s.BlockTimer.InvokeTimer(cursor.Slot)
}

func (s *Scanner) scanSlotAndBlock(start, end uint64, DbTrans *gorm.DB) (uint64, error) {
	for start < end {
		logrus.Infof("Scan Slot[%d]", start)
		slotBody, err := s.BeaconClient.SignedBeaconBlock(beaconClient.CTX, &api.SignedBeaconBlockOpts{
			Block: fmt.Sprintf("%d", start),
		})
		if err != nil {
			var apiErr *api.Error
			if errors.As(err, &apiErr) {
				if apiErr.StatusCode == 404 {
					logrus.Warnf("empty slot[%d]", start)
					start++
					continue
				}
				logrus.Errorf("error[%d]:%v", apiErr.StatusCode, err)
				return start, err
			}
		}

		if slotBody == nil || slotBody.Data == nil {
			logrus.Warnf("slotBody is nil[%d]", start)
			return start, err
		}

		executionBlockNumber, err := slotBody.Data.ExecutionBlockNumber()
		if err != nil {
			return start, err
		}
		logrus.Infof("slotBody[%d]: %v", start, executionBlockNumber)

		if err := s.processBeacon(slotBody, DbTrans); err != nil {
			return start, err
		}

		start++
		select {
		case _, b := <-s.Quit:
			if !b {
				logrus.Infoln("Get close chain Quit scan")
				return start, nil
			}
		default:
		}
	}
	return start, nil
}
