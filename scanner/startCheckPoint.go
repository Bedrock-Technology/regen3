package scanner

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/Bedrock-Technology/regen3/beaconClient"
	"github.com/Bedrock-Technology/regen3/contracts/EigenPod"
	"github.com/Bedrock-Technology/regen3/contracts/Restaking"
	"github.com/Bedrock-Technology/regen3/models"
	eigenpodproofs "github.com/Layr-Labs/eigenpod-proofs-generation"
	"github.com/attestantio/go-eth2-client/api"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func (s *Scanner) InitStartCheckPoint() {
	s.BlockTimer.NewJob(s.Config.CheckStartCheckPoint.IntervalBlock,
		s.Config.CheckStartCheckPoint.FirstRun,
		&StartCheckPointRun{
			scanner: s,
		})
	logrus.Infof("Add StartCheckPoint timer blockNow:%d, firstRun:%d, interval:%d",
		s.BlockTimer.BlockNow(), s.Config.CheckStartCheckPoint.FirstRun, s.Config.CheckStartCheckPoint.IntervalBlock)
}

type StartCheckPointRun struct {
	scanner *Scanner
}

func (s *StartCheckPointRun) JobRun() {
	logrus.Info("StartCheckPointRun")
	for _, pod := range s.scanner.Pods {
		// condition
		if active, err := s.scanner.hasActiveCheckPoint(pod.Address); err != nil || active {
			logrus.Infof("pod[%d] has active checkpoint or err %v", pod.PodIndex, err)
			continue
		}
		if !s.NeedDoCheckPoint(pod.Address, pod.PodIndex) {
			continue
		}
		// send startCheckPoint
		timestamp, err := s.scanner.SendCheckPoint(big.NewInt(int64(pod.PodIndex)), pod.Address)
		if err != nil {
			if errors.Is(err, errBaseFeeTooHigh) {
				logrus.Warnf("%s pod[%d] error:%v", TxStartCheckPoints, pod.PodIndex, errBaseFeeTooHigh)
				return
			}
			logrus.Errorf("send checkpoint pod[%d] error:%v", pod.PodIndex, err)
			panic("send checkpoint err")
		}
		if timestamp != 0 {
			logrus.Infof("need do FillProofs pod[%d] timestamp:%v", pod.PodIndex, timestamp)
			proofs, err := s.scanner.FillProofs(pod.Address, timestamp)
			if err != nil {
				logrus.Errorf("FillProofs pod[%d] timestamp:%v", pod.PodIndex, timestamp)
				panic("FillProofs")
			}
			// write to db
			rest := s.scanner.DBEngine.Model(&models.CheckPoint{}).Where("pod = ?", pod.Address).
				Where("checkpoint_timestamp = ?", timestamp).Where("checkpoint_finalized = ?", 0).
				Update("proofs", string(proofs))
			if rest.Error != nil {
				logrus.Errorf("update pod[%d] timestamp:%v", pod.PodIndex, timestamp)
				panic("update")
			}
		}
	}
}

func (s *Scanner) SendCheckPoint(podId *big.Int, podAddress string) (timestamp uint64, err error) {
	restakingAbi, _ := Restaking.RestakingMetaData.GetAbi()
	input, err := restakingAbi.Pack("startCheckPoint", podId, true)
	if err != nil {
		return 0, err
	}
	realTx, err := s.sendRawTransaction(input, s.Config.RestakingContract)
	if err != nil {
		return 0, err
	}
	txReceipt, err := bind.WaitMined(context.Background(), s.EthClient, realTx)
	if err != nil {
		logrus.Errorf("waiting SendCheckPoint podId %v, error:%v", podId.Uint64(), err)
		return 0, err
	}
	logrus.WithField("Report", "true").Infof("%s pod[%d] tx:%s", TxStartCheckPoints, podId.Uint64(), txReceipt.TxHash)
	// write to db
	err = writeTransaction(s.DBEngine, txReceipt, TxStartCheckPoints)
	if err != nil {
		logrus.Errorln("writeTransaction err:", err)
		return 0, err
	}
	// search the log
	egAbi, _ := EigenPod.EigenPodMetaData.GetAbi()
	contract, err := EigenPod.NewEigenPod(common.HexToAddress(podAddress), s.EthClient)
	if err != nil {
		return 0, err
	}
	for _, log := range txReceipt.Logs {
		if log.Address == common.HexToAddress(podAddress) {
			e, err := egAbi.EventByID(log.Topics[0])
			if err != nil {
				return 0, err
			}
			if e.Name == "CheckpointCreated" {
				r, err := contract.ParseCheckpointCreated(*log)
				if err != nil {
					return 0, err
				}
				eigenPod, _ := EigenPod.NewEigenPod(common.HexToAddress(podAddress), s.EthClient)
				checkPoint, err := eigenPod.CurrentCheckpoint(nil)
				if err != nil {
					return 0, err
				}
				// write to db
				cp := models.CheckPoint{
					Pod:                 podAddress,
					CheckpointTimestamp: r.CheckpointTimestamp,
					BeaconBlockRoot:     hex.EncodeToString(r.BeaconBlockRoot[:]),
					Proofed:             "[]",
					Proofs:              "",
					ActiveNum:           checkPoint.ProofsRemaining.Uint64(),
					BatchSize:           uint64(s.Config.CheckStartCheckPoint.BatchSize),
					CheckpointFinalized: 0,
				}
				timestamp = r.CheckpointTimestamp
				if checkPoint.ProofsRemaining.Uint64() == 0 { // empty checkpoint
					cp.CheckpointFinalized = txReceipt.BlockNumber.Uint64()
					logrus.Infof("CheckpointFinalized, pod[%d] timestamp:%d", podId.Uint64(), timestamp)
					timestamp = 0
				}
				if result := s.DBEngine.Create(&cp); result.Error != nil {
					return 0, result.Error
				}
			}
		}
	}
	return
}

func (s *Scanner) FillProofs(podAddress string, timestamp uint64) ([]byte, error) {
	var validators []uint64
	rest := s.DBEngine.Model(&models.Validator{}).Select("validator_index").Where("pod_address = ?", podAddress).
		Where("credential_verified != ?", 0).Where("withdrawn_on_pod = ?", 0).Find(&validators)
	if rest.Error != nil {
		return nil, rest.Error
	}
	eigenPod, _ := EigenPod.NewEigenPod(common.HexToAddress(podAddress), s.EthClient)
	checkPoint, err := eigenPod.CurrentCheckpoint(nil)
	if err != nil {
		return nil, err
	}
	currentTimestamp, err := eigenPod.CurrentCheckpointTimestamp(nil)
	if err != nil {
		return nil, err
	}
	if currentTimestamp != timestamp || checkPoint.ProofsRemaining.Uint64() != uint64(len(validators)) {
		return nil, fmt.Errorf("currentTimestamp[%v] !=timestamp[%v] or ProofsRemaining[%v] != len(validators)[%v]",
			currentTimestamp, timestamp, checkPoint.ProofsRemaining.Uint64(), len(validators))
	}
	beaconBlockHeader, err := s.BeaconClient.BeaconBlockHeader(beaconClient.CTX, &api.BeaconBlockHeaderOpts{
		Block: "0x" + common.Bytes2Hex(checkPoint.BeaconBlockRoot[:]),
	})
	if err != nil {
		return nil, fmt.Errorf("GetParentBlockRoot err:%v", err)
	}
	beaconBlockState, err := s.BeaconClient.BeaconState(beaconClient.CTX,
		&api.BeaconStateOpts{State: strconv.FormatUint(uint64((*beaconBlockHeader).Data.Header.Message.Slot), 10)})
	if err != nil {
		return nil, fmt.Errorf("BeaconState err:%v", err)
	}
	proofs, err := eigenpodproofs.NewEigenPodProofs(s.Config.ChainId, 0 /* oracleStateCacheExpirySeconds - 5min */)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize prover: %w", err)
	}
	proof, err := proofs.ProveCheckpointProofs(beaconBlockHeader.Data.Header.Message, beaconBlockState.Data, validators)
	if err != nil {
		return nil, fmt.Errorf("failed to prove checkpoint: %w", err)
	}

	return json.Marshal(proof)
}

func (s *StartCheckPointRun) NeedDoCheckPoint(podAddress string, podIndex uint64) bool {
	eigenPod, _ := EigenPod.NewEigenPod(common.HexToAddress(podAddress), s.scanner.EthClient)
	executionLayerGwei, err := eigenPod.WithdrawableRestakedExecutionLayerGwei(nil)
	if err != nil {
		logrus.Errorln("WithdrawableRestakedExecutionLayerGwei error:", err)
		return false
	}
	podBalance, err := s.scanner.EthClient.BalanceAt(context.Background(), common.HexToAddress(podAddress), nil)
	if err != nil {
		logrus.Errorln("BalanceAt error:", err)
		return false
	}
	podBalanceGwei := podBalance.Div(podBalance, big.NewInt(1e9)).Uint64()
	logrus.Infof("pod[%d], podBalanceGwei:%s, executionLayerGwei:%s, minus:%s, threshold:%s", podIndex,
		decimal.NewFromUint64(podBalanceGwei).Mul(decimal.New(1, -9)),
		decimal.NewFromUint64(executionLayerGwei).Mul(decimal.New(1, -9)),
		decimal.NewFromUint64(podBalanceGwei-executionLayerGwei).Mul(decimal.New(1, -9)),
		s.getCheckPointThreshold(podIndex))

	if (podBalanceGwei-executionLayerGwei >= s.scanner.Config.CheckPointThreshold && podIndex != 0) ||
		(podBalanceGwei-executionLayerGwei >= s.scanner.Config.Pod0CheckPointThreshold && podIndex == 0) {
		return s.ifCheckPointDuration(podAddress, podIndex)
	}
	// get pod's active validator num
	var activeCount int64
	rest := s.scanner.DBEngine.Model(&models.Validator{}).Where("pod = ?", podAddress).Where("voluntary_exit = ?", 0).Count(&activeCount)
	if rest.Error != nil {
		logrus.Errorln("Get activeCount error:", rest.Error)
		return false
	}
	// no validator
	if activeCount == 0 {
		logrus.Infof("podIndex %d, have no validators", podIndex)
		if podBalanceGwei-executionLayerGwei >= 32e9 {
			return s.ifCheckPointDuration(podAddress, podIndex)
		}
		logrus.Infof("podIndex %d, No balance", podIndex)
	}
	return false
}

func (s *StartCheckPointRun) ifCheckPointDuration(podAddress string, podIndex uint64) bool {
	var latestCp []models.CheckPoint
	rest := s.scanner.DBEngine.Where("pod = ?", podAddress).Order("updated_at desc").Limit(1).Find(&latestCp)
	if rest.Error != nil {
		logrus.Errorln("Get latest checkpoint error:", rest.Error)
		return false
	}
	if len(latestCp) != 1 {
		return true
	}
	now := time.Now().UTC()
	if now.Sub(latestCp[0].UpdatedAt) < 24*time.Hour {
		logrus.Warnf("pod[%d], latest cp at %v", podIndex, latestCp[0].UpdatedAt)
		return false
	}
	return true
}

func (s *StartCheckPointRun) getCheckPointThreshold(podIndex uint64) string {
	if podIndex == 0 {
		return decimal.NewFromUint64(s.scanner.Config.Pod0CheckPointThreshold).Mul(decimal.New(1, -9)).String()
	}
	return decimal.NewFromUint64(s.scanner.Config.CheckPointThreshold).Mul(decimal.New(1, -9)).String()
}
