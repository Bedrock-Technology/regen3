package scanner

import (
	"context"
	"errors"
	"math/big"

	"github.com/Bedrock-Technology/regen3/contracts/EigenPod"
	"github.com/Bedrock-Technology/regen3/models"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	TxVerifyWithdrawalCredentials = "verifyWithdrawalCredentials"
	TxVerifyCheckPoints           = "verifyCheckPoints"
	TxStartCheckPoints            = "startCheckPoints"
	TxQueueWithdrawals            = "queueWithdrawals"
	TxCompleteQueueWithdrawals    = "completeQueueWithdrawals"
	TxDelegateTo                  = "delegateTo"
)

func addGasBuffer(gasLimit uint64) uint64 {
	return 6 * gasLimit / 5 // add 20% buffer to gas limit
}

func chunk[T any](arr []T, chunkSize uint64) [][]T {
	// Validate the chunkSize to ensure it's positive
	if chunkSize <= 0 {
		panic("chunkSize must be greater than 0")
	}

	// Create a slice to hold the chunks
	var chunks [][]T

	// Loop through the input slice and create chunks
	arrLen := uint64(len(arr))
	for i := uint64(0); i < arrLen; i += chunkSize {
		end := uint64(i + chunkSize)
		if end > arrLen {
			end = arrLen
		}
		chunks = append(chunks, arr[i:end])
	}

	return chunks
}

func writeTransaction(db *gorm.DB, txReceipt *types.Receipt, txType string) error {
	fee := new(big.Int).Mul(txReceipt.EffectiveGasPrice, big.NewInt(int64(txReceipt.GasUsed)))
	txRecord := models.Transaction{
		TxHash: txReceipt.TxHash.String(),
		Status: txReceipt.Status,
		TxType: txType,
		Block:  txReceipt.BlockNumber.Uint64(),
		Fee:    fee.String(),
	}
	if err := db.Create(&txRecord).Error; err != nil {
		return err
	}
	if txReceipt.Status != 1 {
		return errTxReceiptStatusFail
	}
	return nil
}

// Min( Max fee - Base fee, Max priority fee)
// gasPrice = Min(Base fee + Max priority fee, GasFeeCap)
func (s *Scanner) sendRawTransaction(input []byte, toAddress string) (*types.Transaction, error) {
	header, err := s.EthClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, errBaseFeeTooHigh
	}
	if header.BaseFee.Cmp(big.NewInt(10e9)) > 0 {
		logrus.Warnf("Base fee bigger than 10gwei:%s", header.BaseFee)
		return nil, errBaseFeeTooHigh
	}
	gasTipCap, err := s.EthClient.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, errBaseFeeTooHigh
	}
	if gasTipCap.Cmp(big.NewInt(1e9)) > 0 {
		gasTipCap = big.NewInt(1e9)
	}
	to := common.HexToAddress(toAddress)
	gasFeeCap := new(big.Int).Add(new(big.Int).Mul(header.BaseFee, big.NewInt(2)), gasTipCap)
	gasLimit, err := s.EthClient.EstimateGas(context.Background(), ethereum.CallMsg{
		From:      common.HexToAddress(s.Config.KeyAgent.Address),
		To:        &to,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Data:      input,
	})
	if err != nil {
		return nil, err
	}
	opts, err := s.signWithChainIDFromKeyAgent(common.HexToAddress(s.Config.KeyAgent.Address), big.NewInt(int64(s.Config.ChainId)))
	if err != nil {
		return nil, err
	}
	opts.GasTipCap = gasTipCap
	opts.GasFeeCap = gasFeeCap
	opts.GasLimit = addGasBuffer(gasLimit)
	return bind.NewBoundContract(to, abi.ABI{}, s.EthClient, s.EthClient, s.EthClient).RawTransact(opts, input)
}

func (s *Scanner) hasActiveCheckPoint(podAddress string) (bool, error) {
	eigenPod, _ := EigenPod.NewEigenPod(common.HexToAddress(podAddress), s.EthClient)
	currentTimestamp, err := eigenPod.CurrentCheckpointTimestamp(nil)
	if err != nil {
		return false, err
	}
	return currentTimestamp != 0, nil
}

func (s *Scanner) baseFeeBiggerThan() (bool, error) {
	header, err := s.EthClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return false, err
	}
	return header.BaseFee.Cmp(big.NewInt(10e9)) > 0, nil
}

var (
	errBaseFeeTooHigh      = errors.New("base fee too high")
	errTxReceiptStatusFail = errors.New("tx status fail")
)
