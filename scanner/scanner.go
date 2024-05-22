package scanner

import (
	"errors"
	"fmt"
	"github.com/Bedrock-Technology/regen3/beaconClient"
	"github.com/Bedrock-Technology/regen3/blockTimer"
	"github.com/Bedrock-Technology/regen3/config"
	"github.com/Bedrock-Technology/regen3/models"
	txsubmitter "github.com/Layr-Labs/eigenpod-proofs-generation/tx_submitoor/tx_submitter"
	"github.com/attestantio/go-eth2-client/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Scanner struct {
	Config        *config.Config
	DBEngine      *gorm.DB
	EthClient     *ethclient.Client
	BeaconClient  *beaconClient.Client
	BlockTimer    *blockTimer.BlockTimer
	Pods          map[string]models.Pod
	FilterAddress []common.Address
	Submitter     *txsubmitter.EigenPodProofTxSubmitter
	Quit          chan struct{}
}

func New(config *config.Config, quit chan struct{}) *Scanner {
	scanner := &Scanner{
		Quit: quit,
	}
	scanner.Config = config
	//Init db
	gormDb, err := gorm.Open(mysql.Open(config.MysqlDsn))
	if err != nil {
		panic(fmt.Sprintf("Open Mysql err:%v", err))
	}
	scanner.DBEngine = gormDb
	//Init EthClient
	scanner.EthClient, err = ethclient.Dial(config.EthClient)
	if err != nil {
		panic(fmt.Sprintf("Open EthClient err:%v", err))
	}
	//Init BeaconClient
	scanner.BeaconClient, err = beaconClient.NewClient(config.BeaconClient)
	if err != nil {
		panic(fmt.Sprintf("Open BeaconClient err:%v", err))
	}
	//Init Pods
	pods, err := models.GetAllPods(scanner.DBEngine)
	if err != nil {
		panic(fmt.Sprintf("Get all Pods err:%v", err))
	}
	scanner.Pods = pods
	logrus.Info("get pods length:", len(pods))
	//Init FilterAddress
	scanner.FilterAddress = scanner.fillFilterAddress()
	logrus.Info("get filter address: ", scanner.FilterAddress)
	//Init blockTimer
	scanner.BlockTimer = blockTimer.NewBlockTimer()
	//Init timer
	err = scanner.InitVerifyWithdrawCredential()
	if err != nil {
		panic(fmt.Sprintf("InitVerifyWithdrawCredential err:%v", err))
	}

	return scanner
}

func (s *Scanner) fillFilterAddress() []common.Address {
	address := make([]common.Address, 0)
	for _, pod := range s.Pods {
		address = append(address, common.HexToAddress(pod.Address))
	}
	address = append(address,
		common.HexToAddress(s.Config.EigenDelegationManagerContract),
		common.HexToAddress(s.Config.EigenOracleContract),
		common.HexToAddress(s.Config.RestakingContract),
		common.HexToAddress(s.Config.StakingContract),
	)

	return address
}

func (s *Scanner) Scan() {
	//scan slot
	scanTicker := time.NewTicker(12 * time.Second)
	defer scanTicker.Stop()
	for {
		select {
		case <-scanTicker.C:
			s.scan()
		case <-s.Quit:
			logrus.Info("Quit")
			beaconClient.Cancel()
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
	start := cursor.Slot
	logrus.Infof("Scan [%d:%d)", start, end)
	realEnd, err := s.scanSlotAndBlock(start, end)
	if err != nil {
		logrus.Errorf("scanSlotAndBlock [%d:%d),error:%v", start, end, err)
	}
	cursor.Slot = realEnd
	err = models.SaveCursor(s.DBEngine, cursor)
	if err != nil {
		logrus.Errorln("UpdateCursor error:", err)
		return
	}
}

func (s *Scanner) scanSlotAndBlock(start, end uint64) (realEnd uint64, err error) {
LOOP:
	for start < end {
		logrus.Infof("Scan Slot[%d]", start)
		slotBody, err := s.BeaconClient.SignedBeaconBlock(beaconClient.CTX, &api.SignedBeaconBlockOpts{
			Block: fmt.Sprintf("%d", start),
		})
		if err != nil {
			var apiErr *api.Error
			if errors.As(err, &apiErr) {
				switch apiErr.StatusCode {
				case 404:
					logrus.Warnf("empty slot[%d]", start)
					start++
					continue
				default:
					logrus.Errorf("error[%d]:%v", apiErr.StatusCode, err)
					return start, err
				}
			}
		}
		executionBlockNumber, err := slotBody.Data.ExecutionBlockNumber()
		if err != nil {
			return start, err
		}
		logrus.Infof("slotBody[%d]: %v", start, executionBlockNumber)
		err = s.processBeacon(slotBody, s.DBEngine)
		if err != nil {
			return start, err
		}
		err = s.processBlock(executionBlockNumber, s.DBEngine)
		if err != nil {
			return start, err
		}

		s.BlockTimer.InvokeTimer(executionBlockNumber)

		start++
		select {
		case _, b := <-s.Quit:
			if !b {
				logrus.Infoln("Get close chain Quit  scan")
				break LOOP
			}
		default:
		}
	}
	return start, nil
}
