package EigenPod

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum"
	"math/big"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var provider *ethclient.Client
var contract *EigenPod

var (
	RpcHost = ""
	ChainId = int64(0)
)

var (
	ContractAddress = ""
	PodOwner        = ""
	Pod             = ""
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	RpcHost = os.Getenv("HOLESKY_URL")
	ChainId, _ = strconv.ParseInt(os.Getenv("HOLESKY_CHAIN_ID"), 0, 64)
	ContractAddress = os.Getenv("HOLESKY_POD")
	PodOwner = os.Getenv("HOLESKY_POD_WONER")
	Pod = os.Getenv("HOLESKY_POD")

	client, err := ethclient.Dial(RpcHost)
	if err != nil {
		panic(err)
	}
	d, err := NewEigenPod(common.HexToAddress(ContractAddress), client)
	if err != nil {
		panic(err)
	}
	provider = client
	contract = d
}

type ValidatorStatus uint8

func (vs ValidatorStatus) String() string {
	switch vs {
	case ValidatorInactive:
		return "INACTIVE"
	case ValidatorActive:
		return "ACTIVE"
	case ValidatorWithdrawn:
		return "WITHDRAWN"
	}
	return "UNKNOWN"
}

const (
	ValidatorInactive  ValidatorStatus = 0
	ValidatorActive    ValidatorStatus = 1
	ValidatorWithdrawn ValidatorStatus = 2
)

func TestEigenPodCaller_ValidatorPubkeyToInfo(t *testing.T) {
	pubKey := "b78d2c8969f9ab62bc52a8612bc6c1217751819d42c25865048d5003d440c7986c04f730313ad84f2aff2dc256d66df4"
	pubKeyByte, _ := hex.DecodeString(pubKey)
	validatorInfo, err := contract.ValidatorPubkeyToInfo(&bind.CallOpts{}, pubKeyByte)
	if err != nil {
		t.Fatalf("failed to call ValidatorPubkeyToInfo: %v", err)
	}
	t.Logf("ValidatorPubkeyToInfo: %v", validatorInfo)
	if ValidatorStatus(validatorInfo.Status) == ValidatorActive {
		t.Logf("ValidatorPubkeyToInfo status: %v", "Active")
	}
	t.Logf("ValidatorPubkeyToInfo Status: %v", ValidatorStatus(validatorInfo.Status))
}

func TestFilter(t *testing.T) {
	logs, err := provider.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(1362421)),
		ToBlock:   big.NewInt(int64(1362421)),
		Addresses: []common.Address{common.HexToAddress(Pod)},
	})
	if err != nil {
		t.Logf("err:%v", err)
		return
	}
	out, _ := json.Marshal(&logs)
	t.Logf("logs:%v", string(out))
}
