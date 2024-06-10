package scanner

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Bedrock-Technology/regen3/contracts/DelegationManager"
	"github.com/Bedrock-Technology/regen3/contracts/EigenPod"
	"github.com/Bedrock-Technology/regen3/contracts/Restaking"
	"github.com/Bedrock-Technology/regen3/models"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"math/big"
)

func (s *Scanner) InitQueueWithdraw() error {
	blockNow, err := s.findLatestBlockNumber()
	if err != nil {
		return err
	}
	s.BlockTimer.NewJob(blockNow, s.Config.CheckQueueWithdraw.IntervalBlock,
		s.Config.CheckQueueWithdraw.FirstRun,
		&QueueWithdrawRun{
			scanner: s,
		})
	logrus.Infof("Add QueueWithdrawRun timer blockNow:%d, firstRun:%d, interval:%d",
		blockNow, s.Config.CheckQueueWithdraw.FirstRun, s.Config.CheckQueueWithdraw.IntervalBlock)
	return nil
}

type QueueWithdrawRun struct {
	scanner *Scanner
}

func (v *QueueWithdrawRun) JobRun() {
	// todo will rewrite when checkpoint system launch
	logrus.Infof("Start QueueWithdrawRun")
	//get not complete
	for _, pod := range v.scanner.Pods {
		queueWithdrawalsNotCompleted := make([]models.QueueWithdrawals, 0)
		rest := v.scanner.DBEngine.Model(&models.QueueWithdrawals{}).Where("completed = ?", 0).
			Where("pod = ?", pod.Address).Limit(1).Find(&queueWithdrawalsNotCompleted)
		if rest.Error != nil {
			logrus.Errorf("Get QueueWithdrawals failed: %v", rest.Error)
			return
		}
		if len(queueWithdrawalsNotCompleted) == 0 {
			logrus.Infof("pod %s no queueWithdrawals found", pod.Address)
			//todo simple, need modify on M3
			podContract, err := EigenPod.NewEigenPod(common.HexToAddress(pod.Address), v.scanner.EthClient)
			if err != nil {
				logrus.Errorf("Get pod %s podContract: %v", pod.Address, err)
				continue
			}
			shares, err := podContract.WithdrawableRestakedExecutionLayerGwei(&bind.CallOpts{})
			if err != nil {
				logrus.Errorf("Get pod %s WithdrawableRestakedExecutionLayerGwei: %v", pod.Address, err)
				continue
			}
			if shares == 0 {
				logrus.Infof("WithdrawableRestakedExecutionLayerGwei 0")
				continue
			}
			sharesInwei := big.NewInt(0).Mul(big.NewInt(int64(shares)), big.NewInt(1000000000))
			realTx, err := v.scanner.sendQueueWithdrawals(sharesInwei, pod)
			if err != nil {
				if errors.Is(err, errBaseFeeTooHigh) {
					return
				}
				logrus.Errorf("sendQueueWithdrawals index %v error:%v", pod.Address, err)
				panic("sendVerifyWithdrawProof error")
			}
			logrus.Infoln("waiting sendQueueWithdrawals tx:", realTx.Hash())
			txReceipt, err := bind.WaitMined(context.Background(), v.scanner.EthClient, realTx)
			if err != nil {
				logrus.Errorf("waiting sendQueueWithdrawals index %v error:%v", pod.Address, err)
				panic("waiting error")
			}
			logrus.WithField("Report", "true").Infof("sendQueueWithdrawals tx:%s", txReceipt.TxHash)
			//write to db
			fee := big.NewInt(0).Mul(txReceipt.EffectiveGasPrice, big.NewInt(int64(txReceipt.GasUsed)))
			txRecord := models.Transaction{
				TxHash: txReceipt.TxHash.String(),
				Status: txReceipt.Status,
				TxType: TxQueueWithdrawals,
				Fee:    fee.String(),
			}
			rest := v.scanner.DBEngine.Create(&txRecord)
			if rest.Error != nil {
				logrus.Errorf("Insert txRecord[%s] err:%v", txReceipt.TxHash.String(), rest.Error)
			}
			if txReceipt.Status != 1 {
				logrus.Errorf("sendQueueWithdrawals tx: %v status failed:%v", txReceipt.TxHash, txRecord.Status)
				panic("sendQueueWithdrawals")
			}
			err = checkIfWithdrawalQueuedContained(txReceipt.Logs, pod.Owner, v.scanner)
			if err != nil {
				logrus.Errorln("checkIfWithdrawalQueuedContained error:", err)
				panic("checkIfWithdrawalQueuedContained")
			}
		} else {
			logrus.Infof("pod %s has notcompleted queueWithdrawals, try complete", pod.Address)
			queue := queueWithdrawalsNotCompleted[0]
			withdrawal := DelegationManagerWithdrawalQueued{}
			err := json.Unmarshal([]byte(queue.Withdrawal), &withdrawal.Withdrawal)
			if err != nil {
				logrus.Errorf("Unmarshal QueueWithdrawals failed: %v", err)
				continue
			}
			block, err := v.scanner.EthClient.BlockNumber(context.Background())
			if err != nil {
				logrus.Errorf("Get BlockNumber failed: %v", err)
				continue
			}
			if uint64(withdrawal.Withdrawal.StartBlock)+v.scanner.Config.MinWithdrawalDelayBlocks > block {
				logrus.Infof("minWithdrawalDelayBlocks period has not yet passed")
				continue
			}
			logrus.Infof("Withdrawer: %s, owner: %s", withdrawal.Withdrawal.Withdrawer.String(), pod.Owner)
			if withdrawal.Withdrawal.Withdrawer.String() == pod.Owner {
				//send completeQueueWithdrawal
				realTx, err := v.scanner.SendCompleteQueuedWithdrawals(pod, withdrawal, true)
				if err != nil {
					if errors.Is(err, errBaseFeeTooHigh) {
						return
					}
					logrus.Errorf("sendCompleteQueuedWithdrawals index %v error:%v", pod.Address, err)
					panic("sendCompleteQueuedWithdrawals error")
				}
				logrus.Infoln("waiting sendCompleteQueuedWithdrawals tx:", realTx.Hash())
				txReceipt, err := bind.WaitMined(context.Background(), v.scanner.EthClient, realTx)
				if err != nil {
					logrus.Errorf("waiting sendCompleteQueuedWithdrawals index %v error:%v", pod.Address, err)
					panic("waiting error")
				}
				logrus.WithField("Report", "true").Infof("sendCompleteQueuedWithdrawals tx:%s", txReceipt.TxHash)
				//write to db
				fee := big.NewInt(0).Mul(txReceipt.EffectiveGasPrice, big.NewInt(int64(txReceipt.GasUsed)))
				txRecord := models.Transaction{
					TxHash: txReceipt.TxHash.String(),
					Status: txReceipt.Status,
					TxType: TxCompleteQueueWithdrawals,
					Fee:    fee.String(),
				}
				rest := v.scanner.DBEngine.Create(&txRecord)
				if rest.Error != nil {
					logrus.Errorf("Insert txRecord[%s] err:%v", txReceipt.TxHash.String(), rest.Error)
				}
				if txReceipt.Status != 1 {
					logrus.Errorf("sendCompleteQueuedWithdrawals tx: %v status failed:%v", txReceipt.TxHash, txRecord.Status)
					panic("sendQueueWithdrawals")
				}
				err = checkIfCompleteWithdrawalQueuedContained(txReceipt.Logs,
					queue.WithdrawalRoot, v.scanner)
				if err != nil {
					logrus.Errorln("checkIfCompleteWithdrawalQueuedContained error:", err)
					panic("checkIfCompleteWithdrawalQueuedContained")
				}
			}
		}
	}
}

type DelegationManagerWithdrawalQueued struct {
	WithdrawalRoot [32]byte
	Withdrawal     DelegationManager.IDelegationManagerWithdrawal
}

func (s *Scanner) sendQueueWithdrawals(shares *big.Int, pod models.Pod) (*types.Transaction, error) {
	dm, err := DelegationManager.DelegationManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	params := []DelegationManager.IDelegationManagerQueuedWithdrawalParams{
		{
			Strategies: []common.Address{common.HexToAddress("0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0")},
			Shares:     []*big.Int{shares},
			Withdrawer: common.HexToAddress(pod.Owner),
		},
	}
	data, err := dm.Pack("queueWithdrawals", params)
	if err != nil {
		return nil, err
	}
	restakingAbi, err := Restaking.RestakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	input, err := restakingAbi.Pack("callDelegationManager", big.NewInt(int64(pod.PodIndex)), data)
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

func (s *Scanner) SendCompleteQueuedWithdrawals(pod models.Pod, withdrawalQueued DelegationManagerWithdrawalQueued, asToken bool) (*types.Transaction, error) {
	dm, err := DelegationManager.DelegationManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	data, err := dm.Pack("completeQueuedWithdrawals",
		[]DelegationManager.IDelegationManagerWithdrawal{withdrawalQueued.Withdrawal},
		[][]common.Address{{common.HexToAddress("0x0000000000000000000000000000000000000000")}},
		[]*big.Int{big.NewInt(0)},
		[]bool{asToken},
	)
	if err != nil {
		return nil, err
	}
	restakingAbi, err := Restaking.RestakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	input, err := restakingAbi.Pack("callDelegationManager", big.NewInt(int64(pod.PodIndex)), data)
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

func checkIfWithdrawalQueuedContained(logs []*types.Log, podOwner string, s *Scanner) error {
	dmAbi, err := DelegationManager.DelegationManagerMetaData.GetAbi()
	if err != nil {
		return err
	}
	contract, err := DelegationManager.NewDelegationManager(common.HexToAddress(s.Config.EigenDelegationManagerContract), s.EthClient)
	if err != nil {
		return err
	}
	for _, log := range logs {
		if log.Address == common.HexToAddress(s.Config.EigenDelegationManagerContract) {
			e, err := dmAbi.EventByID(log.Topics[0])
			if err != nil {
				return err
			}
			if e.Name == "WithdrawalQueued" {
				r, err := contract.ParseWithdrawalQueued(*log)
				if err != nil {
					return err
				}
				if r.Withdrawal.Withdrawer.String() == podOwner {
					return nil
				}
			}
		}
	}
	return fmt.Errorf("not found")
}

func checkIfCompleteWithdrawalQueuedContained(logs []*types.Log, withdrawalRoot string, s *Scanner) error {
	dmAbi, err := DelegationManager.DelegationManagerMetaData.GetAbi()
	if err != nil {
		return err
	}
	contract, err := DelegationManager.NewDelegationManager(common.HexToAddress(s.Config.EigenDelegationManagerContract), s.EthClient)
	if err != nil {
		return err
	}
	for _, log := range logs {
		if log.Address == common.HexToAddress(s.Config.EigenDelegationManagerContract) {
			e, err := dmAbi.EventByID(log.Topics[0])
			if err != nil {
				return err
			}
			if e.Name == "WithdrawalCompleted" {
				r, err := contract.ParseWithdrawalCompleted(*log)
				if err != nil {
					return err
				}
				if base64.StdEncoding.EncodeToString(r.WithdrawalRoot[:]) == withdrawalRoot {
					return nil
				}
			}
		}
	}
	return fmt.Errorf("not found")
}
