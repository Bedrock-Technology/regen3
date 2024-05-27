package scanner

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/Bedrock-Technology/regen3/beaconClient"
	"github.com/Bedrock-Technology/regen3/blockTimer"
	"github.com/Bedrock-Technology/regen3/config"
	"github.com/Bedrock-Technology/regen3/contracts/EigenPod"
	"github.com/Bedrock-Technology/regen3/contracts/Restaking"
	"github.com/Bedrock-Technology/regen3/contracts/Staking"
	"github.com/Bedrock-Technology/regen3/keyAgentRpc"
	"github.com/Bedrock-Technology/regen3/models"
	txsubmitter "github.com/Layr-Labs/eigenpod-proofs-generation/tx_submitoor/tx_submitter"
	"github.com/attestantio/go-eth2-client/api"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/big"
	"strconv"
	"time"
)

type Scanner struct {
	Config         *config.Config
	DBEngine       *gorm.DB
	EthClient      *ethclient.Client
	BeaconClient   *beaconClient.Client
	BlockTimer     *blockTimer.BlockTimer
	Pods           map[string]models.Pod
	FilterAddress  []common.Address
	Submitter      *txsubmitter.EigenPodProofTxSubmitter
	KeyAgentClient *keyAgentRpc.Client
	Quit           chan struct{}
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
	if scanner.Config.CheckVerifyWithdrawCredential.Enable {
		err = scanner.InitVerifyWithdrawCredential()
		if err != nil {
			panic(fmt.Sprintf("InitVerifyWithdrawCredential err:%v", err))
		}
	}
	if scanner.Config.CheckVerifyWithdrawProof.Enable {
		err = scanner.InitVerifyWithdrawProof()
		if err != nil {
			panic(fmt.Sprintf("InitVerifyWithdrawCredential err:%v", err))
		}
	}
	scanner.KeyAgentClient = keyAgentRpc.NewClient(scanner.Config.KeyAgent)
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
	err = models.SaveCursor(s.DBEngine, cursor)
	if err != nil {
		logrus.Errorln("UpdateCursor error:", err)
		DbTransaction.Rollback()
		return
	}
	rest := DbTransaction.Commit()
	if rest.Error != nil {
		logrus.Errorln("DbTransaction error:", rest.Error)
		return
	}
	s.BlockTimer.InvokeTimer(cursor.Slot)
}

func (s *Scanner) scanSlotAndBlock(start, end uint64, DbTrans *gorm.DB) (realEnd uint64, err error) {
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
		err = s.processBeacon(slotBody, DbTrans)
		if err != nil {
			return start, err
		}
		err = s.processBlock(executionBlockNumber, DbTrans)
		if err != nil {
			return start, err
		}

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

func (s *Scanner) Setup(slot string) {
	contract, err := Restaking.NewRestaking(common.HexToAddress(s.Config.RestakingContract), s.EthClient)
	if err != nil {
		logrus.Errorln("NewRestaking err", err)
		return
	}
	podsNum, err := contract.GetTotalPods(&bind.CallOpts{})
	if err != nil {
		logrus.Errorln("GetTotalPods ", err)
	}
	logrus.Infof("podsNum [%v]", podsNum.Uint64())
	pods := make([]models.Pod, 0)
	for i := uint64(0); i < podsNum.Uint64(); i++ {
		pod := models.Pod{}
		pod.PodIndex = i
		podInfo, err := contract.GetPod(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			logrus.Errorln("podInfo err:", podInfo)
			return
		}
		pod.Address = podInfo.String()
		podOwner, err := contract.PodOwners(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			logrus.Errorln("PodOwners err:", podInfo)
			return
		}
		pod.Owner = podInfo.String()
		//if pod active
		podContract, err := EigenPod.NewEigenPod(podInfo, s.EthClient)
		if err != nil {
			logrus.Errorln("NewEigenPod err:", podInfo)
			return
		}
		restaked, err := podContract.HasRestaked(&bind.CallOpts{})
		if err != nil {
			logrus.Errorln("HasRestaked err:", podInfo)
			return
		}
		if restaked {
			pod.IsCredential = 1
		}
		logrus.Infof("pod %d, podAddress: %s, podOwner: %s, hasRestaked: %v", i, podInfo.String(), podOwner.String(),
			restaked)
		pods = append(pods, pod)
	}
	//validatorStatus
	podsMap := make(map[uint64]models.Pod)
	for _, v := range pods {
		podsMap[v.PodIndex] = v
	}
	stakeContract, err := Staking.NewStaking(common.HexToAddress(s.Config.StakingContract), s.EthClient)
	if err != nil {
		logrus.Errorln("NewStaking err:", err)
		return
	}
	nextId, err := stakeContract.GetNextValidatorId(&bind.CallOpts{})
	if err != nil {
		logrus.Errorln("GetNextValidatorId err:")
		return
	}
	logrus.Infof("nextId:%v", nextId.Uint64())
	validators := make([]models.Validator, 0)
	for i := uint64(0); i < nextId.Uint64(); i++ {
		info, err := stakeContract.ValidatorRegistry(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			logrus.Errorln("ValidatorRegistry err:")
			return
		}
		logrus.Infof("pubkey:%v, podAddress: %v", hex.EncodeToString(info.Pubkey), podsMap[uint64(info.Eigenpod)].Address)
		if info.Restaking {
			validator := models.Validator{
				PubKey:             fmt.Sprintf("0x%s", hex.EncodeToString(info.Pubkey)),
				ValidatorIndex:     0,
				PodAddress:         podsMap[uint64(info.Eigenpod)].Address,
				CredentialVerified: 0,
				WithdrawnOnChain:   0,
				WithdrawnOnPod:     0,
				VoluntaryExit:      0,
			}
			//get if validatorIndexed
			pubKeyS := fmt.Sprintf(`"%s"`, validator.PubKey)
			pubKeyBls := phase0.BLSPubKey{}
			_ = pubKeyBls.UnmarshalJSON([]byte(pubKeyS))
			validatorOnbeacon, err := s.BeaconClient.Validators(beaconClient.CTX, &api.ValidatorsOpts{
				State:   slot,
				PubKeys: []phase0.BLSPubKey{pubKeyBls},
			})
			if err != nil {
				logrus.Errorf("get %s error:", validator.PubKey, err)
				return
			}
			if len(validatorOnbeacon.Data) != 0 { //deposited
				for _, v := range validatorOnbeacon.Data {
					if v.Validator.PublicKey.String() == validator.PubKey {
						validator.ValidatorIndex = uint64(v.Index)
						validators = append(validators, validator)
					}
				}
			}
		}
	}
	//write db
	trans := s.DBEngine.Begin()
	rest := trans.CreateInBatches(&pods, 50)
	if rest.Error != nil {
		trans.Rollback()
		logrus.Errorln("Insert error:", rest.Error)
		return
	}
	rest = trans.CreateInBatches(&validators, 50)
	if rest.Error != nil {
		logrus.Errorln("Insert error:", rest.Error)
		trans.Rollback()
		return
	}
	u, _ := strconv.ParseUint(slot, 0, 64)
	rest = trans.Create(&models.Cursor{
		Slot: u + 1,
		Meme: models.Scanner,
	})
	trans.Commit()
	logrus.Infof("setup complete...donot forget eigenOracle cursor")
}
