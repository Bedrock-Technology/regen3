package scanner

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (s *Scanner) signWithChainIDFromKeyAgent(keyAddr common.Address, chainID *big.Int) (*bind.TransactOpts, error) {
	signer := types.LatestSignerForChainID(chainID)
	return &bind.TransactOpts{
		From: keyAddr,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != keyAddr {
				return nil, fmt.Errorf("address not equal[%s], [%s]", keyAddr.String(), address.String())
			}
			sig, err := s.KeyAgentClient.Eth1SignTransaction(signer.Hash(tx).Bytes())
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, sig)
		},
		Context: context.Background(),
	}, nil
}
