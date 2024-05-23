package scanner

import (
	"encoding/json"
	"fmt"
	"github.com/Bedrock-Technology/regen3/models"
	"github.com/Bedrock-Technology/regen3/proofgen"
	eigenpodproofs "github.com/Layr-Labs/eigenpod-proofs-generation"
	txsubmitter "github.com/Layr-Labs/eigenpod-proofs-generation/tx_submitoor/tx_submitter"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
)

func (s *Scanner) InitVerifyWithdrawProof() error {
	blockNow, err := s.findLatestBlockNumber()
	if err != nil {
		return err
	}
	s.BlockTimer.NewJob(blockNow, s.Config.CheckVerifyWithdrawCredential.IntervalBlock,
		s.Config.CheckVerifyWithdrawCredential.FirstRun,
		&VerifyWithdrawProofRun{
			scanner: s,
		})
	logrus.Infoln("Add VerifyWithdrawProof timer blockNow:", blockNow)
	return nil
}

type VerifyWithdrawProofRun struct {
	scanner *Scanner
}

func (v *VerifyWithdrawProofRun) JobRun() {
	logrus.Info("VerifyWithdrawProofRun")
	for _, pod := range v.scanner.Pods {
		//get pod's validators that need to verify
		validators := make([]models.Validator, 0, batchSize)
		rest := v.scanner.DBEngine.Model(&models.Validator{}).Where("pod_address = ?", pod.Address).
			Where("credential_verified != ?", 0).Where("withdrawn_on_chain != ?", 0).
			Where("withdrawn_on_pod = ?", 0).
			Find(&validators).Limit(batchSize)
		if rest.Error != nil {
			logrus.Infof("Get pod's[%s] validators that need to proof error: %v", pod.Address, rest.Error)
			return
		}
		if len(validators) != 0 {
			logrus.Infof("need send[%v] to proof", validators)
			cursor, err := models.GetCursor(v.scanner.DBEngine, models.EigenOracle)
			if err != nil {
				logrus.Errorln("GetCursor:", err)
				return
			}

			statefileName := fmt.Sprintf(beaconStateFormat, v.scanner.Config.Network, cursor.Slot)
			statefilePath := fmt.Sprintf("%s/%s", v.scanner.Config.DataPath, statefileName)
			headerfileName := fmt.Sprintf(beaconHeaderFormat, v.scanner.Config.Network, cursor.Slot)
			headerfilePath := fmt.Sprintf("%s/%s", v.scanner.Config.DataPath, headerfileName)

			err = v.scanner.getBeaconStates(statefilePath, cursor.Slot)
			if err != nil {
				logrus.Errorln("getBeaconStates err:", err)
				return
			}
			err = v.scanner.getBeaconHeaders(headerfilePath, cursor.Slot)
			if err != nil {
				logrus.Errorln("getBeaconHeaders err:", err)
				return
			}
			proof := proofgen.WithdrawalProof{}
			proof.EigenPodAddress = common.HexToAddress(pod.Address)
			proof.BeaconStateFiles.OracleStateFile = statefilePath
			proof.BeaconStateFiles.OracleBlockHeaderFile = headerfilePath
			proof.WithdrawalDetails.ValidatorIndices = make([]uint64, 0)
			proof.WithdrawalDetails.WithdrawalBlockHeaderFiles = make([]string, 0)
			proof.WithdrawalDetails.WithdrawalBlockBodyFiles = make([]string, 0)
			proof.WithdrawalDetails.HistoricalSummaryStateFiles = make([]string, 0)

			for _, validator := range validators {
				withdrawalHeaderName := fmt.Sprintf(withdrawalHeaderFormat, v.scanner.Config.Network, validator.WithdrawnOnChain)
				withdrawalHeaderPath := fmt.Sprintf("%s/%s", v.scanner.Config.DataPath, withdrawalHeaderName)
				withdrawalBodyName := fmt.Sprintf(withdrawalBodyFormat, v.scanner.Config.Network, validator.WithdrawnOnChain)
				withdrawalBodyPath := fmt.Sprintf("%s/%s", v.scanner.Config.DataPath, withdrawalBodyName)
				historicalSummaryStateName := fmt.Sprintf(historicalSummaryStateFormat, v.scanner.Config.Network,
					historicalSummaryStateSlot(validator.WithdrawnOnChain))
				historicalSummaryStatePath := fmt.Sprintf("%s/%s", v.scanner.Config.DataPath, historicalSummaryStateName)
				//download
				err = v.scanner.getBeaconStates(historicalSummaryStatePath, historicalSummaryStateSlot(validator.WithdrawnOnChain))
				if err != nil {
					logrus.Errorln("getBeaconStates err:", err)
					return
				}
				err = v.scanner.getBeaconHeaders(withdrawalHeaderPath, validator.WithdrawnOnChain)
				if err != nil {
					logrus.Errorln("getBeaconHeaders err:", err)
					return
				}
				err = v.scanner.getBeaconBlocks(withdrawalBodyPath, validator.WithdrawnOnChain)
				if err != nil {
					logrus.Errorln("getBeaconBlocks err:", err)
					return
				}
				proof.WithdrawalDetails.ValidatorIndices = append(proof.WithdrawalDetails.ValidatorIndices, validator.ValidatorIndex)
				proof.WithdrawalDetails.WithdrawalBlockHeaderFiles = append(proof.WithdrawalDetails.WithdrawalBlockHeaderFiles, withdrawalHeaderPath)
				proof.WithdrawalDetails.WithdrawalBlockBodyFiles = append(proof.WithdrawalDetails.WithdrawalBlockBodyFiles, withdrawalBodyPath)
				proof.WithdrawalDetails.HistoricalSummaryStateFiles = append(proof.WithdrawalDetails.HistoricalSummaryStateFiles, historicalSummaryStatePath)
			}
			tx, err := v.scanner.getWithdrawalProofTx(proof, pod.Owner)
			if err != nil {
				logrus.Errorln("getWithdrawalProofTx err:", err)
				return
			}
			logrus.Infof("getWithdrawalProofTx tx: %v", tx.Data())
		}
	}
}

const SlotsPerHistoricalRoot = 8192

func historicalSummaryStateSlot(withdrawalSlot uint64) uint64 {
	return SlotsPerHistoricalRoot * ((withdrawalSlot / SlotsPerHistoricalRoot) + 1)
}

const withdrawalHeaderFormat = "%s_withdrawal_header_%d.json"
const withdrawalBodyFormat = "%s_withdrawal_body_%d.json"
const historicalSummaryStateFormat = "%s_historical_summary_state_%d.json"

func (s *Scanner) getWithdrawalProofTx(proof proofgen.WithdrawalProof, podOwner string) (*types.Transaction, error) {
	proofBytes, err := json.Marshal(proof)
	if err != nil {
		return nil, err
	}
	logrus.Info("getWithdrawalProofTx proof:", string(proofBytes))
	chainClient, err := txsubmitter.NewChainClient(s.EthClient, "", podOwner, 0, 0)
	if err != nil {
		return nil, err
	}
	eigenPodProofs, err := eigenpodproofs.NewEigenPodProofs(s.Config.ChainId, 0)
	if err != nil {
		return nil, err
	}
	submitter := txsubmitter.NewEigenPodProofTxSubmitter(
		*chainClient,
		*eigenPodProofs,
	)
	tx, err := proofgen.VerifyAndProcessWithdrawalsGen(submitter, proof)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
