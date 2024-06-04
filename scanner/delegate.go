package scanner

import (
	"context"
	"github.com/Bedrock-Technology/regen3/contracts/DelegationManager"
	"github.com/Bedrock-Technology/regen3/contracts/Restaking"
	"github.com/Bedrock-Technology/regen3/models"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"math/big"
)

func (s *Scanner) DelegateTo(podIndex int64, operator string) {
	restakingAbi, err := Restaking.RestakingMetaData.GetAbi()
	if err != nil {
		return
	}
	isswe := DelegationManager.ISignatureUtilsSignatureWithExpiry{
		Signature: []byte{},
		Expiry:    big.NewInt(0),
	}
	approverSalt := [32]byte{0}
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
	header, err := s.EthClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		logrus.Errorln("HeaderByNumber err:", err)
		return
	}

	gasTipCap, err := s.EthClient.SuggestGasTipCap(context.Background())
	if err != nil {
		logrus.Errorln("HeaderByNumber err:", err)
		return
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
		logrus.Errorln("EstimateGas err:", err)
		return
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
		logrus.Errorln("RawTransact err:", err)
		return
	}
	logrus.Infoln("waiting delegateTo tx:", realTx.Hash())
	txReceipt, err := bind.WaitMined(context.Background(), s.EthClient, realTx)
	if err != nil {
		logrus.Errorf("waiting delegateTo index %v error:%v", podIndex, err)
		panic("waiting error")
	}
	logrus.WithField("Report", "true").Infof("delegateTo tx:%s", txReceipt.TxHash)
	//write to db
	fee := big.NewInt(0).Mul(txReceipt.EffectiveGasPrice, big.NewInt(int64(txReceipt.GasUsed)))
	txRecord := models.Transaction{
		TxHash: txReceipt.TxHash.String(),
		Status: txReceipt.Status,
		TxType: TxDelegateTo,
		Fee:    fee.String(),
	}
	rest := s.DBEngine.Create(&txRecord)
	if rest.Error != nil {
		logrus.Errorf("Insert txRecord[%s] err:%v", txReceipt.TxHash.String(), rest.Error)
	}
	if txReceipt.Status != 1 {
		logrus.Errorf("delegateTo tx: %v status failed:%v", txReceipt.TxHash, txRecord.Status)
		panic("delegateTo")
	}
	return
}

const TxDelegateTo = "delegateTo"
