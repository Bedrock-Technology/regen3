package scanner

import (
	"context"
	"fmt"
	"github.com/Bedrock-Technology/regen3/contracts/EigenLayerBeaconOracle"
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
				logrus.Errorf("remove state file[%s] failed:%v", statefilePath, err)
			}
			err = os.Remove(headerfilePath)
			if err != nil {
				logrus.Errorf("remove header file[%s] failed:%v", statefilePath, err)
			}
		}
	}
	return nil
}
