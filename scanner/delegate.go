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

func (s *Scanner) DelegateTo(podId int64, operator, restakingContract string) {
	restakingAbi, _ := Restaking.RestakingMetaData.GetAbi()
	isswe := DelegationManager.ISignatureUtilsMixinTypesSignatureWithExpiry{
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

	input, err := restakingAbi.Pack("callDelegationManager", big.NewInt(podId), output)
	if err != nil {
		logrus.Errorln("Pack err:", err)
		return
	}
	realTx, err := s.sendRawTransaction(input, restakingContract, uint64(podId), TxDelegateTo)
	if err != nil {
		logrus.Errorf("delegateTo pod %v error:%v", podId, err)
		return
	}
	logrus.Infoln("waiting delegateTo tx:", realTx.Hash())
	txReceipt, err := bind.WaitMined(context.Background(), s.EthClient, realTx)
	if err != nil {
		logrus.Errorf("waiting delegateTo index %v error:%v", podId, err)
		panic("waiting error")
	}
	logrus.WithField("Report", "true").Infof("pod[%d] %s tx:%s", podId, TxDelegateTo, txReceipt.TxHash)
	if err := writeTransaction(s.DBEngine, txReceipt, TxDelegateTo); err != nil {
		logrus.Errorf("writeTransaction err:%v", err)
		return
	}
	if rest := s.DBEngine.Model(&models.Pod{}).Where("id = ?", podId).Update("delegate_to", operator); rest.Error != nil {
		logrus.Errorln("rest.Error:", rest.Error)
		return
	}
	logrus.Infof("staker:%v delegate to %v", podId, operator)
}
