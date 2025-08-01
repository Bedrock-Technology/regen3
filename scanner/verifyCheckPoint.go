package scanner

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/Bedrock-Technology/regen3/contracts/EigenPod"
	"github.com/Bedrock-Technology/regen3/models"
	eigenpodproofs "github.com/Layr-Labs/eigenpod-proofs-generation"
	eigenutils "github.com/Layr-Labs/eigenpod-proofs-generation/cli/core/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (s *Scanner) InitVerifyCheckPoint() {
	s.BlockTimer.NewJob(s.Config.CheckVerifyCheckPoint.IntervalBlock,
		s.Config.CheckVerifyCheckPoint.FirstRun,
		&VerifyCheckPointRun{scanner: s})
	logrus.Infof("Add CheckVerifyCheckPoint timer blockNow:%d, firstRun:%d, interval:%d",
		s.BlockTimer.BlockNow(), s.Config.CheckVerifyCheckPoint.FirstRun, s.Config.CheckVerifyCheckPoint.IntervalBlock)
}

type VerifyCheckPointRun struct {
	scanner *Scanner
}

func (s *VerifyCheckPointRun) JobRun() {
	logrus.Info("VerifyCheckPointRun")
	for _, pod := range s.scanner.Pods {
		eigenPod, _ := EigenPod.NewEigenPod(common.HexToAddress(pod.Address), s.scanner.EthClient)
		currentTimestamp, err := eigenPod.CurrentCheckpointTimestamp(nil)
		if err != nil || currentTimestamp == 0 {
			logrus.Infof("pod[%d] no need verify checkpoint", pod.PodIndex)
			continue
		}

		var checkPointInDb []models.CheckPoint
		rest := s.scanner.DBEngine.Model(&models.CheckPoint{}).Where("pod = ?", pod.Address).
			Where("checkpoint_finalized = ?", 0).Where("proofs != ''").Find(&checkPointInDb)
		if rest.Error != nil || len(checkPointInDb) != 1 || currentTimestamp != checkPointInDb[0].CheckpointTimestamp {
			logrus.Errorf("pod[%d] checkpoint error", pod.PodIndex)
			panic("checkpoint error")
		}
		logrus.Infof("pod[%d] need verify checkPoint", pod.PodIndex)
		proof := eigenpodproofs.VerifyCheckpointProofsCallParams{}
		if err = json.Unmarshal([]byte(checkPointInDb[0].Proofs), &proof); err != nil {
			logrus.Errorln("json.Unmarshal proofs error", err)
			panic("verify checkpoint error")
		}
		var proofedSlice []uint64
		if err = json.Unmarshal([]byte(checkPointInDb[0].Proofed), &proofedSlice); err != nil {
			logrus.Errorln("json.Unmarshal proofed error", err)
			panic("verify checkpoint error")
		}

		allProofChunks := chunk(proof.BalanceProofs, checkPointInDb[0].BatchSize)
		latestChunk := len(proofedSlice) - 1
		balanceProofs := allProofChunks[latestChunk+1]
		eigenAbi, _ := EigenPod.EigenPodMetaData.GetAbi()
		input, _ := eigenAbi.Pack("verifyCheckpointProofs", EigenPod.BeaconChainProofsBalanceContainerProof{
			BalanceContainerRoot: proof.ValidatorBalancesRootProof.ValidatorBalancesRoot,
			Proof:                proof.ValidatorBalancesRootProof.Proof.ToByteSlice(),
		}, eigenutils.CastBalanceProofs(balanceProofs))

		realTx, err := s.scanner.sendRawTransaction(input, pod.Address, pod.PodIndex, TxVerifyCheckPoints)
		if err != nil {
			if errors.Is(err, errBaseFeeTooHigh) {
				logrus.Warnf("sendRawTransaction pod[%d] error:%v", pod.PodIndex, errBaseFeeTooHigh)
				return
			}
			logrus.Errorf("send VerifyCheckPointRun pod[%d] error:%v", pod.PodIndex, err)
			panic("VerifyCheckPointRun error")
		}
		logrus.Infof("waiting %s pod[%d] tx:%s", TxVerifyCheckPoints, pod.PodIndex, realTx.Hash())
		txReceipt, err := bind.WaitMined(context.Background(), s.scanner.EthClient, realTx)
		if err != nil {
			logrus.Errorf("wait sendVerifyCheckPointRun pod %v error:%v", pod.Address, err)
			panic("waiting error")
		}
		logrus.WithField("Report", "true").Infof("%s pod[%d] vcount[%d] %d/%d tx:%s", TxVerifyCheckPoints,
			pod.PodIndex, len(balanceProofs), latestChunk+2, len(allProofChunks), txReceipt.TxHash)

		if err = writeTransaction(s.scanner.DBEngine, txReceipt, TxVerifyCheckPoints); err != nil {
			logrus.Errorln("writeTransaction err:", err)
			panic("writeTransaction error")
		}

		matched, finalized, withdrawn := s.scanner.checkIfCheckPointContained(txReceipt.Logs, len(balanceProofs), pod.Address)
		if !matched {
			logrus.Errorln("checkIfCheckPointContained")
			panic("checkIfCheckPointContained error")
		}
		proofedSlice = append(proofedSlice, uint64(latestChunk+1))

		finalizedBlock := uint64(0)
		if finalized {
			finalizedBlock = txReceipt.BlockNumber.Uint64()
			if len(proofedSlice) != len(allProofChunks) {
				logrus.Errorln("proofedSlice !=allProofChunks err:", err)
				panic("proofedSlice != allProofChunks")
			}
		}
		logrus.Infof("pod[%d] timestamp:%d, finalized:%v, withdrawn:%v", pod.PodIndex, checkPointInDb[0].CheckpointTimestamp,
			finalized, withdrawn)
		if err := updateCheckPoint(s.scanner.DBEngine, proofedSlice, finalizedBlock, txReceipt.BlockNumber.Uint64(), pod.Address,
			checkPointInDb[0].CheckpointTimestamp, withdrawn); err != nil {
			panic("updateCheckPoint")
		}
		time.Sleep(1 * time.Second)
	}
}

func (s *Scanner) checkIfCheckPointContained(logs []*types.Log, needCheck int, podAddress string) (matched, finalized bool, withdrawnId []uint64) {
	egAbi, _ := EigenPod.EigenPodMetaData.GetAbi()
	eigenPod, _ := EigenPod.NewEigenPod(common.HexToAddress(podAddress), s.EthClient)
	validatorContains := 0
	for _, log := range logs {
		if log.Address == common.HexToAddress(podAddress) {
			e, err := egAbi.EventByID(log.Topics[0])
			if err != nil {
				panic("EventByID not found")
			}
			switch e.Name {
			case "ValidatorCheckpointed":
				validatorContains++
			case "CheckpointFinalized":
				finalized = true
			case "ValidatorWithdrawn":
				ew, err := eigenPod.ParseValidatorWithdrawn(*log)
				if err != nil {
					panic("ParseValidatorWithdrawn error")
				}
				info, err := eigenPod.ValidatorPubkeyHashToInfo(&bind.CallOpts{}, ew.PubkeyHash)
				if err != nil {
					panic("ValidatorPubkeyHashToInfo error")
				}
				withdrawnId = append(withdrawnId, info.ValidatorIndex)
			}
		}
	}
	matched = needCheck == validatorContains
	return
}

func updateCheckPoint(db *gorm.DB, proofed []uint64, finalized, block uint64, podAddress string, timestamp uint64, withdrawal []uint64) error {
	proofedJson, _ := json.Marshal(proofed)
	trans := db.Begin()
	rest := trans.Model(&models.CheckPoint{}).Where("pod = ?", podAddress).
		Where("checkpoint_timestamp = ?", timestamp).Updates(map[string]interface{}{
		"proofed": string(proofedJson), "checkpoint_finalized": finalized,
	})
	if rest.Error != nil {
		trans.Rollback()
		return rest.Error
	}
	if len(withdrawal) != 0 {
		rest := trans.Model(&models.Validator{}).Where("validator_index in ?", withdrawal).Update("withdrawn_on_pod", block)
		if rest.Error != nil {
			trans.Rollback()
			return rest.Error
		}
	}
	return trans.Commit().Error
}
