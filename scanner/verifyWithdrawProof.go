package scanner

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Bedrock-Technology/regen3/contracts/EigenPod"
	"github.com/Bedrock-Technology/regen3/contracts/Restaking"
	"github.com/Bedrock-Technology/regen3/models"
	"github.com/Bedrock-Technology/regen3/proofgen"
	eigenpodproofs "github.com/Layr-Labs/eigenpod-proofs-generation"
	txsubmitter "github.com/Layr-Labs/eigenpod-proofs-generation/tx_submitoor/tx_submitter"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"math/big"
)

func (s *Scanner) InitVerifyWithdrawProof() error {
	blockNow, err := s.findLatestBlockNumber()
	if err != nil {
		return err
	}
	s.BlockTimer.NewJob(blockNow, s.Config.CheckVerifyWithdrawProof.IntervalBlock,
		s.Config.CheckVerifyWithdrawProof.FirstRun,
		&VerifyWithdrawProofRun{
			scanner: s,
		})
	logrus.Infof("Add VerifyWithdrawProof timer blockNow:%d, firstRun:%d, interval:%d",
		blockNow, s.Config.CheckVerifyWithdrawProof.FirstRun, s.Config.CheckVerifyWithdrawProof.IntervalBlock)
	return nil
}

type VerifyWithdrawProofRun struct {
	scanner *Scanner
}

func checkIfHistoricalSlot(validators []models.Validator, cursorSlot uint64) error {
	for _, v := range validators {
		if historicalSummaryStateSlot(v.WithdrawnOnChain) > cursorSlot {
			logrus.Infof("summary state slot:%v , sursorSlot: %v",
				historicalSummaryStateSlot(v.WithdrawnOnChain), cursorSlot)
			return errors.New("slot not contained")
		}
	}
	return nil
}

func (v *VerifyWithdrawProofRun) JobRun() {
	logrus.Info("VerifyWithdrawProofRun")
	for _, pod := range v.scanner.Pods {
		//get pod's validators that need to verify
		validators := make([]models.Validator, 0, batchSizeProof)
		rest := v.scanner.DBEngine.Model(&models.Validator{}).Where("pod_address = ?", pod.Address).
			Where("withdrawn_on_chain != ?", 0).Where("withdrawn_on_pod = ?", 0).
			Limit(batchSizeProof).Find(&validators)
		if rest.Error != nil {
			logrus.Errorln("Get pod's[%s] validators that need to proof error: %v", pod.Address, rest.Error)
			return
		}
		if len(validators) != 0 {
			logrus.Infof("need send[%v] to proof", validators)
			cursor, err := models.GetCursor(v.scanner.DBEngine, models.EigenOracle)
			if err != nil {
				logrus.Errorln("GetCursor:", err)
				return
			}
			err = checkIfHistoricalSlot(validators, cursor.Slot)
			if err != nil {
				logrus.Infof("checkIfHistoricalSlot")
				continue
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
			logrus.Infof("getWithdrawalProofTx tx: %v", hex.EncodeToString(tx.Data()))
			realTx, err := v.scanner.sendVerifyWithdrawProof(tx, big.NewInt(int64(pod.PodIndex)))
			if err != nil {
				if errors.Is(err, errBaseFeeTooHigh) {
					return
				}
				logrus.Errorf("sendVerifyWithdrawProof index %v error:%v", validators, err)
				panic("sendVerifyWithdrawProof error")
			}
			logrus.Infoln("waiting sendVerifyWithdrawProof tx:", realTx.Hash())
			txReceipt, err := bind.WaitMined(context.Background(), v.scanner.EthClient, realTx)
			if err != nil {
				logrus.Errorf("waiting sendVerifyWithdrawProof index %v error:%v", validators, err)
				panic("waiting error")
			}
			logrus.WithField("Report", "true").Infof("sendVerifyWithdrawProof tx:%s", txReceipt.TxHash)
			//write to db
			fee := big.NewInt(0).Mul(txReceipt.EffectiveGasPrice, big.NewInt(int64(txReceipt.GasUsed)))
			txRecord := models.Transaction{
				TxHash: txReceipt.TxHash.String(),
				Status: txReceipt.Status,
				TxType: TxVerifyWithdrawalProof,
				Fee:    fee.String(),
			}
			rest := v.scanner.DBEngine.Create(&txRecord)
			if rest.Error != nil {
				logrus.Errorf("Insert txRecord[%s] err:%v", txReceipt.TxHash.String(), rest.Error)
			}
			if txReceipt.Status != 1 {
				logrus.Errorf("sendVerifyWithdrawProof tx: %v status failed:%v", validators, txRecord.Status)
				panic("sendVerifyWithdrawProof")
			}
			needCheck := make([]uint64, 0)
			for _, validator := range validators {
				needCheck = append(needCheck, validator.ValidatorIndex)
			}
			err = checkIfFullWithdrawalContained(txReceipt.Logs, needCheck, pod.Address, v.scanner)
			if err != nil {
				logrus.Errorln("checkIfEventContained error:", err)
				panic("checkIfEventContained")
			}
		}
	}
}

func checkIfFullWithdrawalContained(logs []*types.Log, needCheck []uint64, podAddress string, s *Scanner) error {
	egAbi, err := EigenPod.EigenPodMetaData.GetAbi()
	if err != nil {
		return err
	}
	contract, err := EigenPod.NewEigenPod(common.HexToAddress(podAddress), s.EthClient)
	if err != nil {
		return err
	}
	validatorContains := make([]uint64, 0)
	for _, log := range logs {
		if log.Address == common.HexToAddress(podAddress) {
			e, err := egAbi.EventByID(log.Topics[0])
			if err != nil {
				return err
			}
			if e.Name == "FullWithdrawalRedeemed" {
				r, err := contract.ParseFullWithdrawalRedeemed(*log)
				if err != nil {
					return err
				}
				validatorContains = append(validatorContains, r.ValidatorIndex.Uint64())
			}
		}
	}
	if len(needCheck) != len(validatorContains) {
		return fmt.Errorf(" not equal")
	}
	for _, n := range needCheck {
		find := false
		for _, c := range validatorContains {
			if n == c {
				find = true
			}
		}
		if find == false {
			return fmt.Errorf("not equal")
		}
	}
	return nil
}

const SlotsPerHistoricalRoot = 8192

func historicalSummaryStateSlot(withdrawalSlot uint64) uint64 {
	return SlotsPerHistoricalRoot * ((withdrawalSlot / SlotsPerHistoricalRoot) + 1)
}

const withdrawalHeaderFormat = "%s_withdrawal_header_%d.json"
const withdrawalBodyFormat = "%s_withdrawal_body_%d.json"
const historicalSummaryStateFormat = "%s_historical_summary_state_%d.json"

func (s *Scanner) getWithdrawalProofTx(proof proofgen.WithdrawalProof, podOwner string) (tx *types.Transaction, err error) {
	proofBytes, err := json.Marshal(proof)
	if err != nil {
		return nil, err
	}
	defer func() {
		if info := recover(); info != nil {
			err = errors.New("getWithdrawalProofTx recover")
		}
	}()
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
	tx, err = proofgen.VerifyAndProcessWithdrawalsGen(submitter, proof)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (s *Scanner) sendVerifyWithdrawProof(tx *types.Transaction, podId *big.Int) (*types.Transaction, error) {
	//parse data to param
	eigenPodAbi, err := EigenPod.EigenPodMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	m := map[string]interface{}{}
	err = eigenPodAbi.Methods["verifyAndProcessWithdrawals"].Inputs.UnpackIntoMap(m, tx.Data()[4:])
	if err != nil {
		return nil, err
	}
	restakingAbi, err := Restaking.RestakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	input, err := restakingAbi.Pack("verifyAndProcessWithdrawals", podId, m["oracleTimestamp"],
		m["stateRootProof"], m["withdrawalProofs"], m["validatorFieldsProofs"], m["validatorFields"],
		m["withdrawalFields"])
	if err != nil {
		return nil, err
	}
	header, err := s.EthClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	if header.BaseFee.Cmp(big.NewInt(20000000000)) > 0 {
		logrus.Warnf("Base fee bigger than 20gwei:%s", header.BaseFee)
		return nil, errBaseFeeTooHigh
	}
	//gasTipCap := big.NewInt(150000000) //0.15gwei
	gasTipCap, err := s.EthClient.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, err
	}
	//max 1gwei
	if gasTipCap.Cmp(big.NewInt(1000000000)) > 0 {
		gasTipCap = big.NewInt(1000000000)
	}
	gasFeeCap := new(big.Int).Add(header.BaseFee.Mul(header.BaseFee, big.NewInt(2)), gasTipCap)
	restakingContractAddress := common.HexToAddress(s.Config.RestakingContract)
	gasLimit, err := s.EthClient.EstimateGas(context.Background(), ethereum.CallMsg{
		From:      common.HexToAddress(s.Config.KeyAgent.Address),
		To:        &restakingContractAddress,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Data:      input,
	})
	if err != nil {
		return nil, err
	}
	opts, err := s.signWithChainIDFromKeyAgent(common.HexToAddress(s.Config.KeyAgent.Address),
		big.NewInt(int64(s.Config.ChainId)))
	opts.GasTipCap = gasTipCap
	opts.GasFeeCap = gasFeeCap
	opts.GasLimit = addGasBuffer(gasLimit)
	//Min( Max fee - Base fee, Max priority fee)
	//gasPrice = Min(Base fee + Max priority fee, GasFeeCap)
	realTx, err := bind.NewBoundContract(common.HexToAddress(s.Config.RestakingContract), abi.ABI{}, s.EthClient, s.EthClient,
		s.EthClient).RawTransact(opts, input)
	if err != nil {
		return nil, err
	}
	return realTx, nil
}
