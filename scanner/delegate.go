package scanner

import (
	"context"
	"math/big"

	"github.com/Bedrock-Technology/regen3/contracts/DelegationManager"
	"github.com/Bedrock-Technology/regen3/contracts/Restaking"
	"github.com/Bedrock-Technology/regen3/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

func (s *Scanner) DelegateTo(podIndex int64, operator string) {
	restakingAbi, _ := Restaking.RestakingMetaData.GetAbi()
	isswe := DelegationManager.ISignatureUtilsSignatureWithExpiry{
		Signature: []byte{},
		Expiry:    big.NewInt(0),
	}
	approverSalt := [32]byte{}
	delegationManagerAbi, err := DelegationManager.DelegationManagerMetaData.GetAbi()
	if err != nil {
		return
	}
	output, err := delegationManagerAbi.Pack("delegateTo", common.HexToAddress(operator), isswe, approverSalt)
	if err != nil {
		return
	}

	input, err := restakingAbi.Pack("callDelegationManager", big.NewInt(podIndex), output)
	if err != nil {
		logrus.Errorln("Pack err:", err)
		return
	}
	realTx, err := s.sendRawTransaction(input, s.Config.RestakingContract, uint64(podIndex), TxDelegateTo)
	if err != nil {
		logrus.Errorf("delegateTo pod %v error:%v", podIndex, err)
		return
	}
	logrus.Infoln("waiting delegateTo tx:", realTx.Hash())
	txReceipt, err := bind.WaitMined(context.Background(), s.EthClient, realTx)
	if err != nil {
		logrus.Errorf("waiting delegateTo index %v error:%v", podIndex, err)
		panic("waiting error")
	}
	logrus.WithField("Report", "true").Infof("pod[%d] %s tx:%s", podIndex, TxDelegateTo, txReceipt.TxHash)
	if err := writeTransaction(s.DBEngine, txReceipt, TxDelegateTo); err != nil {
		logrus.Errorf("writeTransaction err:%v", err)
		return
	}
	if rest := s.DBEngine.Model(&models.Pod{}).Where("pod_index = ?", podIndex).Update("delegate_to", operator); rest.Error != nil {
		logrus.Errorln("rest.Error:", rest.Error)
		return
	}
	logrus.Infof("staker:%v delegate to %v", podIndex, operator)
}
