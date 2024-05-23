package Restaking

import (
	"github.com/Bedrock-Technology/regen3/contracts/EigenPod"
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
var contract *Restaking

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
	ContractAddress = "0xf59fc684Ad69A7F8B1C563D8b9fC4003F841F4Ef"
	PodOwner = os.Getenv("HOLESKY_POD_WONER")

	client, err := ethclient.Dial(RpcHost)
	if err != nil {
		panic(err)
	}
	d, err := NewRestaking(common.HexToAddress(ContractAddress), client)
	if err != nil {
		panic(err)
	}
	provider = client
	contract = d
}

func TestNewRestakingCaller(t *testing.T) {
	podsNum, err := contract.GetTotalPods(&bind.CallOpts{})
	if err != nil {
		t.Fatal("err:", err)
	}
	t.Logf("podsNum [%v]", podsNum.Uint64())
	for i := uint64(0); i < podsNum.Uint64(); i++ {
		podInfo, err := contract.GetPod(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			t.Fatal("podInfo err:", podInfo)
		}
		podOwner, err := contract.PodOwners(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			t.Fatal("podInfo err:", podInfo)
		}
		//if pod active
		podContract, err := EigenPod.NewEigenPod(podInfo, provider)
		if err != nil {
			t.Fatal("NewEigenPod err:", err)
		}
		restaked, err := podContract.HasRestaked(&bind.CallOpts{})
		if err != nil {
			t.Fatal("HasRestaked err:", err)
		}
		t.Logf("pod %d, podAddress: %s, podOwner: %s, hasRestaked: %v", i, podInfo.String(), podOwner.String(),
			restaked)
	}
}
