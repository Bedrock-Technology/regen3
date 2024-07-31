package scanner

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/Bedrock-Technology/regen3/contracts/DelegationManager"
	"github.com/Bedrock-Technology/regen3/contracts/EigenPod"
	"github.com/Bedrock-Technology/regen3/models"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
	"math/big"
)

func (s *Scanner) processBlock(blockNumber uint64, orm *gorm.DB) error {
	logs, err := s.EthClient.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(blockNumber)),
		ToBlock:   big.NewInt(int64(blockNumber)),
		Addresses: s.FilterAddress,
	})
	if err != nil {
		return err
	}
	for _, log := range logs {
		if log.Address.String() == s.Config.EigenDelegationManagerContract {
			if err := s.processDelegationManager(log.TxHash, log, orm); err != nil {
				return err
			}
		}
		if _, exists := s.Pods[log.Address.String()]; exists {
			if err := s.processEigenPods(log.TxHash, log, orm); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Scanner) processEigenPods(txHash common.Hash, log types.Log, orm *gorm.DB) error {
	txReceipt, err := s.EthClient.TransactionReceipt(context.Background(), txHash)
	if err != nil || txReceipt.Status != 1 {
		return err
	}

	eigenPodAbi, err := EigenPod.EigenPodMetaData.GetAbi()
	if err != nil {
		return err
	}
	event, err := eigenPodAbi.EventByID(log.Topics[0])
	if err != nil {
		return err
	}
	contract, err := EigenPod.NewEigenPod(log.Address, s.EthClient)
	if err != nil {
		return err
	}

	switch event.Name {
	case "ValidatorRestaked":
		eigenPodRestaked, err := contract.ParseValidatorRestaked(log)
		if err != nil {
			return err
		}
		return orm.Model(&models.Validator{}).
			Where("validator_index = ?", eigenPodRestaked.ValidatorIndex.Uint64()).
			Where("pod_address = ?", log.Address.String()).
			Update("credential_verified", txReceipt.BlockNumber.Uint64()).Error
	case "FullWithdrawalRedeemed":
		eigenPodFullWithdrawalRedeemed, err := contract.ParseValidatorWithdrawn(log)
		if err != nil {
			return err
		}
		validator := models.Validator{}
		if err := orm.Where("validator_index = ?", eigenPodFullWithdrawalRedeemed.ValidatorIndex.Uint64()).
			Where("pod_address = ?", log.Address.String()).Take(&validator).Error; err != nil {
			return err
		}
		return orm.Model(&models.Validator{}).
			Where("validator_index = ?", validator.ValidatorIndex).
			Where("pod_address = ?", log.Address.String()).
			Update("withdrawn_on_pod", txReceipt.BlockNumber.Uint64()).Error
	}
	return nil
}

func (s *Scanner) processDelegationManager(txHash common.Hash, log types.Log, orm *gorm.DB) error {
	txReceipt, err := s.EthClient.TransactionReceipt(context.Background(), txHash)
	if err != nil || txReceipt.Status != 1 {
		return err
	}

	delegationManagerAbi, err := DelegationManager.DelegationManagerMetaData.GetAbi()
	if err != nil {
		return err
	}
	event, err := delegationManagerAbi.EventByID(log.Topics[0])
	if err != nil {
		return err
	}
	contract, err := DelegationManager.NewDelegationManager(log.Address, s.EthClient)
	if err != nil {
		return err
	}

	switch event.Name {
	case "StakerDelegated":
		eigenPodRestaked, err := contract.ParseStakerDelegated(log)
		if err != nil {
			return err
		}
		for _, pod := range s.Pods {
			if pod.Owner == eigenPodRestaked.Staker.String() {
				return orm.Model(&models.Pod{}).
					Where("owner = ?", pod.Owner).
					Update("delegate_to", eigenPodRestaked.Operator.String()).Error
			}
		}
	case "WithdrawalQueued":
		r, err := contract.ParseWithdrawalQueued(log)
		if err != nil {
			return err
		}
		for _, pod := range s.Pods {
			if pod.Owner == r.Withdrawal.Withdrawer.String() {
				withdrawal, _ := json.Marshal(&r.Withdrawal)
				queue := models.QueueWithdrawals{
					Pod:            pod.Address,
					WithdrawalRoot: base64.StdEncoding.EncodeToString(r.WithdrawalRoot[:]),
					Withdrawal:     string(withdrawal),
					StartBlock:     uint64(r.Withdrawal.StartBlock),
					Completed:      0,
				}
				return orm.Create(&queue).Error
			}
		}
	case "WithdrawalCompleted":
		r, err := contract.ParseWithdrawalCompleted(log)
		if err != nil {
			return err
		}
		root := base64.StdEncoding.EncodeToString(r.WithdrawalRoot[:])
		return orm.Model(&models.QueueWithdrawals{}).
			Where("withdrawal_root = ?", root).
			Update("completed", 1).Error
	}
	return nil
}
