package scanner

import (
	"context"
	"github.com/Bedrock-Technology/regen3/contracts/EigenLayerBeaconOracle"
	"github.com/Bedrock-Technology/regen3/models"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math/big"
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
			rest := orm.Model(&models.Cursor{}).Where("meme = ?", models.EigenOracle).
				UpdateColumn("slot", beaconOracleUpdate.Slot.Uint64())
			if rest.Error != nil {
				return rest.Error
			}
			logrus.Info("Update EigenBeaconOracle slot:", beaconOracleUpdate.Slot)
		}
	}
	return nil
}
