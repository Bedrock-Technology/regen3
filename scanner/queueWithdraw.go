package scanner

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/Bedrock-Technology/regen3/contracts/DelegationManager"
	"github.com/Bedrock-Technology/regen3/contracts/EigenPod"
	"github.com/Bedrock-Technology/regen3/contracts/Restaking"
	"github.com/Bedrock-Technology/regen3/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func (s *Scanner) InitQueueWithdraw() {
	s.BlockTimer.NewJob(s.Config.CheckQueueWithdraw.IntervalBlock,
		s.Config.CheckQueueWithdraw.FirstRun,
		&QueueWithdrawRun{scanner: s})
	logrus.Infof("Add QueueWithdrawRun timer blockNow:%d, firstRun:%d, interval:%d",
		s.BlockTimer.BlockNow(), s.Config.CheckQueueWithdraw.FirstRun, s.Config.CheckQueueWithdraw.IntervalBlock)
}

type QueueWithdrawRun struct {
	scanner *Scanner
}

func (v *QueueWithdrawRun) JobRun() {
	logrus.Infof("Start QueueWithdrawRun")
	for _, pod := range v.scanner.Pods {
		if active, err := v.scanner.hasActiveCheckPoint(pod.Address); err != nil || active {
			logrus.Infof("pod[%d] has active checkpoint or err:%v", pod.PodIndex, err)
			continue
		}
		var queueWithdrawalsNotCompleted []models.QueueWithdrawals
		rest := v.scanner.DBEngine.Model(&models.QueueWithdrawals{}).Where("completed = ?", 0).
			Where("pod = ?", pod.Address).Order("start_block asc").Limit(1).Find(&queueWithdrawalsNotCompleted)
		if rest.Error != nil {
			logrus.Errorf("Get QueueWithdrawals failed: %v", rest.Error)
			return
		}
		if len(queueWithdrawalsNotCompleted) == 1 {
			logrus.Infof("pod[%d] tryCompleteQueue root:%s", pod.PodIndex, queueWithdrawalsNotCompleted[0].WithdrawalRoot)
			_ = v.scanner.tryCompleteQueue(queueWithdrawalsNotCompleted[0], pod)
		}
		sumUncompleteGwei, err := v.scanner.GetWithdrawalUncompletedGwei(pod.Address, pod.PodIndex)
		if err != nil {
			logrus.Errorf("GetWithdrawalUncompletedGwei pod[%d] error:%v", pod.PodIndex, err)
			return
		}
		podContract, _ := EigenPod.NewEigenPod(common.HexToAddress(pod.Address), v.scanner.EthClient)
		shares, err := podContract.WithdrawableRestakedExecutionLayerGwei(&bind.CallOpts{})
		if err != nil {
			logrus.Errorf("Get pod[%d] WithdrawableRestakedExecutionLayerGwei: %v", pod.PodIndex, err)
			return
		}
		queueGwei := shares - sumUncompleteGwei
		logrus.Infof("pod[%d] shares:%s, sumUncompleteGwei:%s, minus:%s", pod.PodIndex,
			decimal.NewFromUint64(shares).Mul(decimal.New(1, -9)),
			decimal.NewFromUint64(sumUncompleteGwei).Mul(decimal.New(1, -9)),
			decimal.NewFromUint64(queueGwei).Mul(decimal.New(1, -9)))
		if queueGwei > 0 {
			sharesWei := big.NewInt(0).Mul(big.NewInt(int64(queueGwei)), big.NewInt(1e9))
			err := v.scanner.sendQueueWithdrawals(sharesWei, pod)
			if err != nil {
				if errors.Is(err, errBaseFeeTooHigh) {
					logrus.Warnf("sendRawTransaction pod[%d] error:%v", pod.PodIndex, errBaseFeeTooHigh)
					return
				}
				logrus.Errorf("end sendQueueWithdrawals pod[%d] error:%v", pod.PodIndex, err)
				panic("send sendQueueWithdrawals err")
			}
		} else {
			logrus.Infof("pod[%d] no need QueueWithdraw", pod.PodIndex)
		}
	}
}

type DelegationManagerWithdrawalQueued struct {
	Withdrawal     DelegationManager.IDelegationManagerWithdrawal
	WithdrawalRoot [32]byte
}

func (s *Scanner) sendQueueWithdrawals(shares *big.Int, pod models.Pod) error {
	dm, _ := DelegationManager.DelegationManagerMetaData.GetAbi()
	params := []DelegationManager.IDelegationManagerQueuedWithdrawalParams{
		{
			Strategies: []common.Address{common.HexToAddress("0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0")},
			Shares:     []*big.Int{shares},
			Withdrawer: common.HexToAddress(pod.Owner),
		},
	}
	data, err := dm.Pack("queueWithdrawals", params)
	if err != nil {
		return err
	}
	restakingAbi, _ := Restaking.RestakingMetaData.GetAbi()
	input, err := restakingAbi.Pack("callDelegationManager", big.NewInt(int64(pod.PodIndex)), data)
	if err != nil {
		return err
	}
	realTx, err := s.sendRawTransaction(input, s.Config.RestakingContract)
	if err != nil {
		return err
	}
	txReceipt, err := bind.WaitMined(context.Background(), s.EthClient, realTx)
	if err != nil {
		logrus.Errorf("waiting %s podId[%d], error:%v", TxQueueWithdrawals, pod.PodIndex, err)
		return err
	}
	logrus.WithField("Report", "true").Infof("%s pod[%d] shares:%s tx:%s", TxQueueWithdrawals, pod.PodIndex,
		decimal.NewFromBigInt(shares, -18).Truncate(9), txReceipt.TxHash)
	err = writeTransaction(s.DBEngine, txReceipt, TxQueueWithdrawals)
	if err != nil {
		logrus.Errorln("writeTransaction err:", err)
		return err
	}
	return checkIfWithdrawalQueuedContained(txReceipt.Logs, pod.Owner, pod.Address, s)
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
	return s.sendRawTransaction(input, s.Config.RestakingContract)
}

func checkIfWithdrawalQueuedContained(logs []*types.Log, podOwner, podAddress string, s *Scanner) error {
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
					withdrawal, _ := json.Marshal(&r.Withdrawal)
					queue := models.QueueWithdrawals{
						Pod:            podAddress,
						WithdrawalRoot: base64.StdEncoding.EncodeToString(r.WithdrawalRoot[:]),
						Withdrawal:     string(withdrawal),
						StartBlock:     uint64(r.Withdrawal.StartBlock),
						Completed:      0,
					}
					return s.DBEngine.Create(&queue).Error
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
				root := base64.StdEncoding.EncodeToString(r.WithdrawalRoot[:])
				if root == withdrawalRoot {
					return s.DBEngine.Model(&models.QueueWithdrawals{}).
						Where("withdrawal_root = ?", root).
						Update("completed", log.BlockNumber).Error
				}
			}
		}
	}
	return fmt.Errorf("not found")
}

func (s *Scanner) GetWithdrawalUncompletedGwei(podAddress string, podIndex uint64) (uint64, error) {
	queueWithdrawalUnCompletedShareSumWei := big.NewInt(0)
	var qws []models.QueueWithdrawals
	rest := s.DBEngine.Model(&models.QueueWithdrawals{}).Where("pod = ?", podAddress).
		Where("completed = ?", 0).Find(&qws)
	if rest.Error != nil {
		logrus.Errorln("get QueueWithdrawals error:", rest.Error)
		return 0, rest.Error
	}
	logrus.Infof("pod[%d] GetWithdrawalUncompletedGwei len(qws):%v", podIndex, len(qws))
	for _, queueWithdrawal := range qws {
		var qwp DelegationManager.IDelegationManagerQueuedWithdrawalParams
		if err := json.Unmarshal([]byte(queueWithdrawal.Withdrawal), &qwp); err != nil {
			logrus.Errorln("Unmarshal error:", err)
			return 0, err
		}
		for _, share := range qwp.Shares {
			queueWithdrawalUnCompletedShareSumWei.Add(queueWithdrawalUnCompletedShareSumWei, share)
		}
	}
	return queueWithdrawalUnCompletedShareSumWei.Div(queueWithdrawalUnCompletedShareSumWei, big.NewInt(1e9)).Uint64(), nil
}

func (s *Scanner) tryCompleteQueue(queueWithdrawalsNotCompleted models.QueueWithdrawals, pod models.Pod) error {
	var withdrawal DelegationManagerWithdrawalQueued
	if err := json.Unmarshal([]byte(queueWithdrawalsNotCompleted.Withdrawal), &withdrawal.Withdrawal); err != nil {
		logrus.Errorf("Unmarshal QueueWithdrawals failed: %v", err)
		return nil
	}
	block, err := s.EthClient.BlockNumber(context.Background())
	if err != nil {
		logrus.Errorf("Get BlockNumber failed: %v", err)
		return nil
	}
	if uint64(withdrawal.Withdrawal.StartBlock)+s.Config.MinWithdrawalDelayBlocks > block {
		logrus.Infof("pod[%d] minWithdrawalDelayBlocks[%d] startBlock[%d] block[%d] diff[%d] period has not yet passed", pod.PodIndex,
			s.Config.MinWithdrawalDelayBlocks, withdrawal.Withdrawal.StartBlock, block, block-uint64(withdrawal.Withdrawal.StartBlock))
		return nil
	}
	if withdrawal.Withdrawal.Withdrawer.String() == pod.Owner {
		realTx, err := s.SendCompleteQueuedWithdrawals(pod, withdrawal, true)
		if err != nil {
			if errors.Is(err, errBaseFeeTooHigh) {
				logrus.Warnf("%s pod[%d] error:%v", TxCompleteQueueWithdrawals, pod.PodIndex, errBaseFeeTooHigh)
				return nil
			}
			logrus.Errorf("sendCompleteQueuedWithdrawals pod[%d] error:%v", pod.PodIndex, err)
			panic("sendCompleteQueuedWithdrawals error")
		}
		logrus.Infof("waiting %s pod[%d] tx:%s", TxCompleteQueueWithdrawals, pod.PodIndex, realTx.Hash())
		txReceipt, err := bind.WaitMined(context.Background(), s.EthClient, realTx)
		if err != nil {
			logrus.Errorf("waiting sendCompleteQueuedWithdrawals index %v error:%v", pod.Address, err)
			panic("waiting error")
		}
		// big.NewInt(0).Div(withdrawal.Withdrawal.Shares[0], big.NewInt(1e9))
		logrus.WithField("Report", "true").Infof("%s pod[%d] shares:%s tx:%s", TxCompleteQueueWithdrawals,
			pod.PodIndex, decimal.NewFromBigInt(withdrawal.Withdrawal.Shares[0], -18).Truncate(9), txReceipt.TxHash)
		if err := writeTransaction(s.DBEngine, txReceipt, TxCompleteQueueWithdrawals); err != nil {
			logrus.Errorln("writeTransaction err:", err)
			panic("writeTransaction error")
		}
		if err := checkIfCompleteWithdrawalQueuedContained(txReceipt.Logs, queueWithdrawalsNotCompleted.WithdrawalRoot, s); err != nil {
			logrus.Errorln("checkIfCompleteWithdrawalQueuedContained error:", err)
			panic("checkIfCompleteWithdrawalQueuedContained")
		}
	}
	return nil
}
