package scanner

import (
	"context"
	"encoding/hex"
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

func (s *Scanner) InitVerifyWithdrawCredential() error {
	blockNow, err := s.findLatestBlockNumber()
	if err != nil {
		return err
	}
	s.BlockTimer.NewJob(blockNow, s.Config.CheckVerifyWithdrawCredential.IntervalBlock,
		s.Config.CheckVerifyWithdrawCredential.FirstRun,
		&VerifyWithdrawCredentialRun{
			scanner: s,
		})
	logrus.Infof("Add VerifyWithdrawCredential timer blockNow:%d, firstRun:%d, interval:%d",
		blockNow, s.Config.CheckVerifyWithdrawCredential.FirstRun, s.Config.CheckVerifyWithdrawCredential.IntervalBlock)
	return nil
}

type VerifyWithdrawCredentialRun struct {
	scanner *Scanner
}

func (v *VerifyWithdrawCredentialRun) JobRun() uint64 {
	logrus.Info("VerifyWithdrawCredentialRun")
	delta := uint64(0)
	for _, pod := range v.scanner.Pods {
		if pod.IsCredential == 1 {
			//get pod's validators that need to verify
			validators := make([]uint64, 0, v.scanner.Config.CheckVerifyWithdrawCredential.BatchSize)
			rest := v.scanner.DBEngine.Model(&models.Validator{}).Select("validator_index").Where("pod_address = ?", pod.Address).
				Where("credential_verified = ?", 0).Where("withdrawn_on_chain = ?", 0).
				Where("withdrawn_on_pod = ?", 0).Limit(v.scanner.Config.CheckVerifyWithdrawCredential.BatchSize).Find(&validators)
			if rest.Error != nil {
				logrus.Infof("Get pod's[%s] validators that need to verify error: %v", pod.Address, rest.Error)
				return delta
			}
			if len(validators) != 0 {
				logrus.Infof("need send[%v] to credential", validators)
				cursor, err := models.GetCursor(v.scanner.DBEngine, models.EigenOracle)
				if err != nil {
					logrus.Errorln("GetCursor:", err)
					return delta
				}

				statefileName := fmt.Sprintf(beaconStateFormat, v.scanner.Config.Network, cursor.Slot)
				statefilePath := fmt.Sprintf("%s/%s", v.scanner.Config.DataPath, statefileName)
				headerfileName := fmt.Sprintf(beaconHeaderFormat, v.scanner.Config.Network, cursor.Slot)
				headerfilePath := fmt.Sprintf("%s/%s", v.scanner.Config.DataPath, headerfileName)

				err = v.scanner.getBeaconStates(statefilePath, cursor.Slot)
				if err != nil {
					logrus.Errorln("getBeaconStates err:", err)
					return delta
				}
				err = v.scanner.getBeaconHeaders(headerfilePath, cursor.Slot)
				if err != nil {
					logrus.Errorln("getBeaconHeaders err:", err)
					return delta
				}
				tx, err := v.scanner.getVerifyWithdrawCredentialTx(statefilePath, headerfilePath,
					common.HexToAddress(pod.Address), pod.Owner, validators)
				if err != nil {
					logrus.Errorln("getVerifyWithdrawCredentialTx err:", err)
					return delta
				}
				logrus.Infof("tx data:%v", hex.EncodeToString(tx.Data()))
				realTx, err := v.scanner.sendVerifyWithdrawCredential(tx, big.NewInt(int64(pod.PodIndex)))
				if err != nil {
					if errors.Is(err, errBaseFeeTooHigh) {
						return delta
					}
					logrus.Errorf("sendVerifyWithdrawCredential index %v error:%v", validators, err)
					panic("sendVerifyWithdrawCredential error")
				}
				logrus.Infoln("waiting sendVerifyWithdrawCredential tx:", realTx.Hash())
				txReceipt, err := bind.WaitMined(context.Background(), v.scanner.EthClient, realTx)
				if err != nil {
					logrus.Errorf("waiting sendVerifyWithdrawCredential index %v error:%v", validators, err)
					panic("waiting error")
				}
				logrus.WithField("Report", "true").Infof("sendVerifyWithdrawCredential tx:%s", txReceipt.TxHash)
				//write to db
				fee := big.NewInt(0).Mul(txReceipt.EffectiveGasPrice, big.NewInt(int64(txReceipt.GasUsed)))
				txRecord := models.Transaction{
					TxHash: txReceipt.TxHash.String(),
					Status: txReceipt.Status,
					TxType: TxVerifyWithdrawalCredentials,
					Fee:    fee.String(),
				}
				rest := v.scanner.DBEngine.Create(&txRecord)
				if rest.Error != nil {
					logrus.Errorf("Insert txRecord[%s] err:%v", txReceipt.TxHash.String(), rest.Error)
				}
				if txReceipt.Status != 1 {
					logrus.Errorf("sendVerifyWithdrawCredential tx: %v status failed:%v", validators, txRecord.Status)
					panic("sendVerifyWithdrawCredential")
				}
				err = checkIfEventContained(txReceipt.Logs, validators, pod.Address, v.scanner)
				if err != nil {
					logrus.Errorln("checkIfEventContained error:", err)
					panic("checkIfEventContained")
				}
				delta = txReceipt.BlockNumber.Uint64()
			}
		}
	}
	return delta
}

func checkIfEventContained(logs []*types.Log, needCheck []uint64, podAddress string, s *Scanner) error {
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
			if e.Name == "ValidatorRestaked" {
				r, err := contract.ParseValidatorRestaked(*log)
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

func (s *Scanner) getVerifyWithdrawCredentialTx(oracleStateFile, oracleHeaderFile string,
	podAddress common.Address, podOwner string,
	validators []uint64) (tx *types.Transaction, err error) {
	defer func() {
		if info := recover(); info != nil {
			err = errors.New("getVerifyWithdrawCredentialTx recover")
		}
	}()
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
	tx, err = proofgen.VerifyWithdrawalCredentialsGen2(submitter, oracleStateFile, oracleHeaderFile, podAddress, validators)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (s *Scanner) sendVerifyWithdrawCredential(tx *types.Transaction, podId *big.Int) (*types.Transaction, error) {
	//parse data to param
	eigenPodAbi, err := EigenPod.EigenPodMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	m := map[string]interface{}{}
	err = eigenPodAbi.Methods["verifyWithdrawalCredentials"].Inputs.UnpackIntoMap(m, tx.Data()[4:])
	if err != nil {
		return nil, err
	}
	restakingAbi, err := Restaking.RestakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	input, err := restakingAbi.Pack("verifyWithdrawalCredentials", podId, m["oracleTimestamp"],
		m["stateRootProof"], m["validatorIndices"], m["validatorFieldsProofs"], m["validatorFields"])
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
