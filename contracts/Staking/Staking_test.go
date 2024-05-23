package Staking

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"math/big"
	"os"
	"strconv"
	"testing"
)

var provider *ethclient.Client
var contract *Staking

var (
	RpcHost = ""
	ChainId = int64(0)
)

var (
	ContractAddress = ""
	PodOwner        = ""
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	RpcHost = os.Getenv("HOLESKY_URL")
	ChainId, _ = strconv.ParseInt(os.Getenv("HOLESKY_CHAIN_ID"), 0, 64)
	ContractAddress = "0x0f59BfDEdbB4ECc965be28484BfD968552fD5C67"
	PodOwner = os.Getenv("HOLESKY_POD_WONER")

	client, err := ethclient.Dial(RpcHost)
	if err != nil {
		panic(err)
	}
	d, err := NewStaking(common.HexToAddress(ContractAddress), client)
	if err != nil {
		panic(err)
	}
	provider = client
	contract = d
}

func TestNewStaking(t *testing.T) {
	nextId, err := contract.GetNextValidatorId(&bind.CallOpts{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("nextId:%v", nextId)
	for i := uint64(0); i < nextId.Uint64(); i++ {
		info, err := contract.ValidatorRegistry(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			t.Fatal("err")
		}
		out, _ := json.Marshal(info)
		t.Logf("info:%v", string(out))
	}

}
