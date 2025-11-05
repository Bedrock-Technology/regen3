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
	"github.com/Bedrock-Technology/regen3/contracts/PodManager"
	"github.com/Bedrock-Technology/regen3/contracts/Restaking"
	"github.com/Bedrock-Technology/regen3/models"
	eigenpodproofs "github.com/Layr-Labs/eigenpod-proofs-generation"
	"github.com/attestantio/go-eth2-client/api"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/samber/lo"
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
			logrus.Infof("pod[%d][%s] has active checkpoint or err %v", pod.PodIndex, s.scanner.restakingVersion(pod.Restaking), err)
			continue
		}
		if !s.NeedDoCheckPoint(pod.Address, pod.Owner, pod.Restaking, pod.PodIndex) {
			continue
		}
		// send startCheckPoint
		timestamp, err := s.scanner.SendCheckPoint(big.NewInt(int64(pod.PodIndex)), pod.Address, pod.Restaking)
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
		time.Sleep(1 * time.Second)
	}
}

func (s *Scanner) SendCheckPoint(podId *big.Int, podAddress, restaking string) (timestamp uint64, err error) {
	restakingAbi, _ := Restaking.RestakingMetaData.GetAbi()
	input, err := restakingAbi.Pack("startCheckPoint", podId, true)
	if err != nil {
		return 0, err
	}
	realTx, err := s.sendRawTransaction(input, restaking, podId.Uint64(), TxStartCheckPoints)
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
	// hold on senconds for rpc sync? maybe balance to another one
	time.Sleep(time.Minute)
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

func (s *StartCheckPointRun) NeedDoCheckPoint(podAddress, podOwner, restaking string, podIndex uint64) bool {
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
	logrus.Infof("pod[%d][%s], podBalanceGwei:%s, executionLayerGwei:%s, minus:%s, threshold:%s", podIndex,
		s.scanner.restakingVersion(restaking),
		decimal.NewFromUint64(podBalanceGwei).Mul(decimal.New(1, -9)),
		decimal.NewFromUint64(executionLayerGwei).Mul(decimal.New(1, -9)),
		decimal.NewFromUint64(podBalanceGwei-executionLayerGwei).Mul(decimal.New(1, -9)),
		s.getCheckPointThreshold(podIndex))

	switch s.scanner.restakingVersion(restaking) {
	case "N":
		if (podBalanceGwei-executionLayerGwei >= s.scanner.Config.CheckPointThreshold && podIndex != 0) ||
			(podBalanceGwei-executionLayerGwei >= s.scanner.Config.Pod0CheckPointThreshold && podIndex == 0) {
			return s.ifCheckPointDuration(podAddress, podIndex, restaking)
		}
	case "P":
		if podBalanceGwei-executionLayerGwei >= s.scanner.Config.CheckPointThreshold {
			return s.ifCheckPointDuration(podAddress, podIndex, restaking)
		}
		if s.checkSharesLessThan(podOwner, podAddress, restaking, podIndex) {
			return s.ifCheckPointDuration(podAddress, podIndex, restaking)
		}
	default:
		return false
	}

	return false
}

func (s *StartCheckPointRun) ifCheckPointDuration(podAddress string, podIndex uint64, restaking string) bool {
	if s.scanner.Config.Network != "mainnet" {
		logrus.Infof("not mainnet:%v", s.scanner.Config.Network)
		return true
	}
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
		logrus.Warnf("pod[%d][%s], latest cp at %v", podIndex, s.scanner.restakingVersion(restaking), latestCp[0].UpdatedAt)
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

func (s *StartCheckPointRun) checkSharesLessThan(podOwner, podAddress, restaking string, podIndex uint64) bool {
	eigenPodManager, _ := PodManager.NewPodManager(common.HexToAddress(s.scanner.Config.EigenPodManager), s.scanner.EthClient)
	shares, err := eigenPodManager.StakerDepositShares(nil, common.HexToAddress(podOwner), common.HexToAddress("0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0"))
	if err != nil {
		logrus.Errorln("Get Staker Shares error:", err)
		return false
	}
	var validators []models.Validator
	rest := s.scanner.DBEngine.Model(&models.Validator{}).Where("pod_address = ?", podAddress).
		Where("credential_verified != ?", 0).Where("withdrawn_on_chain = ?", 0).Where("withdrawn_on_pod = ?", 0).
		Find(&validators)
	if rest.Error != nil {
		logrus.Errorln("Get validators error:", rest.Error)
		return false
	}
	chunks := lo.Chunk(validators, 100)
	totalBalance := uint64(0)
	for _, v := range chunks {
		pubKeys := lo.Map(v, func(item models.Validator, index int) string {
			return item.PubKey
		})
		validators, err := s.scanner.BeaconClient.ValidatorsByPubkeys(pubKeys)
		if err != nil {
			logrus.Errorln("ValidatorsByPubkeys error:", err)
			return false
		}
		for _, validator := range validators.Data {
			totalBalance = totalBalance + uint64(validator.Balance)
		}
	}
	sharesGwei := shares.Div(shares, big.NewInt(1e9)).Uint64()
	logrus.Infof("pod[%d][%s], depositShares:%s, totalBalance:%s, minus:%s, threshold:%s", podIndex,
		s.scanner.restakingVersion(restaking),
		decimal.NewFromUint64(sharesGwei).Mul(decimal.New(1, -9)),
		decimal.NewFromUint64(totalBalance).Mul(decimal.New(1, -9)),
		decimal.NewFromInt(int64(totalBalance-sharesGwei)).Mul(decimal.New(1, -9)),
		decimal.NewFromUint64(s.scanner.Config.CheckPointThreshold).Mul(decimal.New(1, -9)))
	return sharesGwei+s.scanner.Config.CheckPointThreshold < totalBalance
}
