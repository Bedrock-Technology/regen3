package scanner

import (
	"context"
	"fmt"
	"github.com/Bedrock-Technology/regen3/contracts/DelegationManager"
	"github.com/Bedrock-Technology/regen3/contracts/EigenLayerBeaconOracle"
	"github.com/Bedrock-Technology/regen3/contracts/EigenPod"
	"github.com/Bedrock-Technology/regen3/models"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math/big"
	"os"
)

func (s *Scanner) processBlock(blockNumber uint64, orm *gorm.DB) error {
	//get event log
	logs, err := s.EthClient.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(blockNumber)),
		ToBlock:   big.NewInt(int64(blockNumber)),
		Addresses: s.FilterAddress,
	})
	if err != nil {
		return err
	}
	for _, log := range logs {
		logrus.Info("topic:", log.Topics[0].Hex())
		logrus.Info("txHash:", log.TxHash.String())
		logrus.Info("address:", log.Address.String())

		switch log.Address.String() {
		case s.Config.EigenOracleContract:
			err := s.processEigenOracle(log.TxHash, log, orm)
			if err != nil {
				return err
			}
		case s.Config.EigenDelegationManagerContract:
			err := s.processDelegationManager(log.TxHash, log, orm)
			if err != nil {
				return err
			}
		}
		//pods events
		if _, b := s.Pods[log.Address.String()]; b {
			err := s.processEigenPods(log.TxHash, log, orm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Scanner) processEigenOracle(txHash common.Hash, log types.Log, orm *gorm.DB) error {
	//log's status
	txReceipt, err := s.EthClient.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return err
	}
	if txReceipt.Status == 1 { //OK
		eigenOracleAbi, err := EigenLayerBeaconOracle.EigenLayerBeaconOracleMetaData.GetAbi()
		if err != nil {
			return err
		}
		event, err := eigenOracleAbi.EventByID(log.Topics[0])
		if err != nil {
			return err
		}
		contract, err := EigenLayerBeaconOracle.NewEigenLayerBeaconOracle(
			common.HexToAddress(s.Config.EigenOracleContract),
			s.EthClient)
		if err != nil {
			return err
		}
		switch event.Name {
		case "EigenLayerBeaconOracleUpdate":
			beaconOracleUpdate, err := contract.ParseEigenLayerBeaconOracleUpdate(log)
			if err != nil {
				return err
			}
			cursor, err := models.GetCursor(orm, models.EigenOracle)
			if err != nil {
				return err
			}
			lastSlot := cursor.Slot
			cursor.Slot = beaconOracleUpdate.Slot.Uint64()
			rest := orm.Save(&cursor)
			if rest.Error != nil {
				return rest.Error
			}
			logrus.Infof("Update EigenBeaconOracle slot before[%d], now[%d]:", lastSlot, cursor.Slot)
			//remove statefile, headerfile
			statefileName := fmt.Sprintf(beaconStateFormat, s.Config.Network, lastSlot)
			statefilePath := fmt.Sprintf("%s/%s", s.Config.DataPath, statefileName)
			headerfileName := fmt.Sprintf(beaconHeaderFormat, s.Config.Network, lastSlot)
			headerfilePath := fmt.Sprintf("%s/%s", s.Config.DataPath, headerfileName)
			err = os.Remove(statefilePath)
			if err != nil {
				logrus.Warnf("remove state file[%s] failed:%v", statefilePath, err)
			}
			logrus.Infof("remove state file[%s]", statefilePath)
			err = os.Remove(headerfilePath)
			if err != nil {
				logrus.Warnf("remove header file[%s] failed:%v", headerfilePath, err)
			}
			logrus.Infof("remove header file[%s]", headerfilePath)
		}
	}
	return nil
}

func (s *Scanner) processEigenPods(txHash common.Hash, log types.Log, orm *gorm.DB) error {
	//log's status
	txReceipt, err := s.EthClient.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return err
	}
	if txReceipt.Status == 1 {
		eigenPodAbi, err := EigenPod.EigenPodMetaData.GetAbi()
		if err != nil {
			return err
		}
		event, err := eigenPodAbi.EventByID(log.Topics[0])
		if err != nil {
			return err
		}
		contract, err := EigenPod.NewEigenPod(
			log.Address, s.EthClient)
		if err != nil {
			return err
		}
		switch event.Name {
		case "ValidatorRestaked":
			eigenPodRestaked, err := contract.ParseValidatorRestaked(log)
			if err != nil {
				return err
			}
			rest := orm.Model(&models.Validator{}).
				Where("validator_index = ?", eigenPodRestaked.ValidatorIndex.Uint64()).
				Where("pod_address = ?", log.Address.String()).
				UpdateColumn("credential_verified", txReceipt.BlockNumber.Uint64())
			logrus.Infof("ValidatorRestaked %v", eigenPodRestaked.ValidatorIndex)
			return rest.Error
		case "FullWithdrawalRedeemed":
			eigenPodFullWithdrawalRedeemed, err := contract.ParseFullWithdrawalRedeemed(log)
			if err != nil {
				return err
			}
			//del data
			validator := models.Validator{}
			rest := orm.Where("validator_index = ?", eigenPodFullWithdrawalRedeemed.ValidatorIndex.Uint64()).
				Where("pod_address = ?", log.Address.String()).
				Take(&validator)
			if rest.Error != nil {
				return rest.Error
			}
			rest = orm.Model(&models.Validator{}).
				Where("validator_index = ?", validator.ValidatorIndex).
				Where("pod_address = ?", log.Address.String()).
				UpdateColumn("withdrawn_on_pod", txReceipt.BlockNumber.Uint64())
			if rest.Error != nil {
				return rest.Error
			}
			withdrawalHeaderName := fmt.Sprintf(withdrawalHeaderFormat, s.Config.Network, validator.WithdrawnOnChain)
			withdrawalHeaderPath := fmt.Sprintf("%s/%s", s.Config.DataPath, withdrawalHeaderName)
			withdrawalBodyName := fmt.Sprintf(withdrawalBodyFormat, s.Config.Network, validator.WithdrawnOnChain)
			withdrawalBodyPath := fmt.Sprintf("%s/%s", s.Config.DataPath, withdrawalBodyName)
			historicalSummaryStateName := fmt.Sprintf(historicalSummaryStateFormat, s.Config.Network,
				historicalSummaryStateSlot(validator.WithdrawnOnChain))
			historicalSummaryStatePath := fmt.Sprintf("%s/%s", s.Config.DataPath, historicalSummaryStateName)
			err = os.Remove(withdrawalHeaderPath)
			if err != nil {
				logrus.Errorf("remove state file[%s] failed:%v", withdrawalHeaderPath, err)
			}
			logrus.Infof("remove state file[%s]", withdrawalHeaderPath)
			err = os.Remove(withdrawalBodyPath)
			if err != nil {
				logrus.Errorf("remove state file[%s] failed:%v", withdrawalBodyPath, err)
			}
			logrus.Infof("remove state file[%s]", withdrawalBodyPath)
			err = os.Remove(historicalSummaryStatePath)
			if err != nil {
				logrus.Errorf("remove state file[%s] failed:%v", historicalSummaryStatePath, err)
			}
			logrus.Infof("remove state file[%s]", historicalSummaryStatePath)
		}
	}
	return nil
}

func (s *Scanner) processDelegationManager(txHash common.Hash, log types.Log, orm *gorm.DB) error {
	//log's status
	txReceipt, err := s.EthClient.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return err
	}
	if txReceipt.Status == 1 {
		DelegationManagerAbi, err := DelegationManager.DelegationManagerMetaData.GetAbi()
		if err != nil {
			return err
		}
		event, err := DelegationManagerAbi.EventByID(log.Topics[0])
		if err != nil {
			return err
		}
		contract, err := DelegationManager.NewDelegationManager(
			log.Address, s.EthClient)
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
					//our pods
					rest := orm.Model(&models.Pod{}).
						Where("owner = ?", pod.Owner).
						UpdateColumn("delegate_to", eigenPodRestaked.Operator.String())
					logrus.Infof("staker:%v delegate to %v", eigenPodRestaked.Staker.String(), eigenPodRestaked.Operator)
					return rest.Error
				}
			}
		}
	}
	return nil
}
