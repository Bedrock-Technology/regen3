package proofgen

import (
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	eigenpodproofs "github.com/Layr-Labs/eigenpod-proofs-generation"
	txsubmitter "github.com/Layr-Labs/eigenpod-proofs-generation/tx_submitoor/tx_submitter"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Bedrock-Technology/regen3/contracts/EigenPod"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var provider *ethclient.Client
var contract *EigenPod.EigenPod

var (
	RpcHost    = ""
	ChainId    = int64(0)
	BeaconHost = ""
)

var (
	ContractAddress = ""
	PodOwner        = ""
	PodAddress      = ""
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	RpcHost = os.Getenv("HOLESKY_URL")
	ChainId, _ = strconv.ParseInt(os.Getenv("HOLESKY_CHAIN_ID"), 0, 64)
	ContractAddress = os.Getenv("HOLESKY_POD")
	PodOwner = os.Getenv("HOLESKY_POD_WONER")
	BeaconHost = os.Getenv("HOLESKY_BEACON_URL")
	PodAddress = os.Getenv("HOLESKY_POD")

	client, err := ethclient.Dial(RpcHost)
	if err != nil {
		panic(err)
	}
	d, err := EigenPod.NewEigenPod(common.HexToAddress(ContractAddress), client)
	if err != nil {
		panic(err)
	}
	provider = client
	contract = d
}

func TestVerifyWithdrawalCredentialsGen(t *testing.T) {
	//get state
	block := 1603985
	headerUrl := fmt.Sprintf("%s/eth/v1/beacon/headers/%v", BeaconHost, block)
	t.Logf("url: %v", headerUrl)
	headerUrlResponse, err := http.Get(headerUrl)
	if err != nil {
		t.Logf("%v", err)
		return
	}
	t.Logf("status:%v", headerUrlResponse.Status)
	defer headerUrlResponse.Body.Close()
	headerBody, err := io.ReadAll(headerUrlResponse.Body)
	if err != nil {
		t.Logf("Read:%v", err)
		return
	}

	stateUrl := fmt.Sprintf("%s/eth/v2/debug/beacon/states/%v", BeaconHost, block)
	t.Logf("url: %v", stateUrl)
	stateUrlResponse, err := http.Get(stateUrl)
	if err != nil {
		t.Logf("%v", err)
		return
	}
	t.Logf("status:%v", stateUrlResponse.Status)
	defer stateUrlResponse.Body.Close()
	stateBody, err := io.ReadAll(stateUrlResponse.Body)
	if err != nil {
		t.Logf("Read:%v", err)
		return
	}

	t.Log("do verify")
	tx, err := VerifyWithdrawalCredentialsGen(uint64(ChainId), provider, stateBody, headerBody, common.HexToAddress(PodAddress), common.HexToAddress(PodOwner), []uint64{1702059})
	if err != nil {
		t.Logf("%v", err)
		return
	}
	t.Logf("fake:%v", tx.Hash())

	//send to chain
	opts := getTransOpts("HOLESKY_ACCOUNT_0")
	caller := bind.NewBoundContract(*tx.To(), abi.ABI{}, provider, provider, provider)
	realTx, err := caller.RawTransact(opts, tx.Data())
	if err != nil {
		t.Logf("RawTransact err:%v", err)
		return
	}
	t.Logf("raw tx:%v", realTx.Hash())
	/*
	   proofgen_test.go:90: status:200 OK
	   proofgen_test.go:98: do verify
	   proofgen_test.go:104: fake:0x39037cf8d6456c8dbfcaa84e84b97b73b06b14d4411d52f9b78709c3c1a73b49
	   proofgen_test.go:113: raw tx:0x57f8588d3e24d12ceb5196dd214e892b321e1462ae1480afbe76a8143d66d0d8
	*/
}

// Using hardhat fork holesky to test
//
//	 hardhat: {
//	  chainId: 17000,
//	  forking: {
//	    enabled: true,
//	    url: process.env.HOLESKY_ARCHIVE_URL!,
//	    blockNumber: 1502050,
//	  },
//	},
//
// npx hardhat node
func TestVerifyWithdrawalCredentialsGen2(t *testing.T) {
	//oracleStateFile := "../data/holesky_slot_1603985.json"
	//oracleHeaderFile := "../data/holesky_block_header_1603985.json"

	//oracleStateFile := "../data/holesky_slot_1705206.json"
	//oracleHeaderFile := "../data/holesky_block_header_1705206.json"

	oracleStateFile := "../data/holesky_state_1705206.json"
	oracleHeaderFile := "../data/holesky_header_1705206.json"

	chainClient, err := txsubmitter.NewChainClient(provider, "", PodOwner, 0, 0)
	if err != nil {
		t.Logf("%v", err)
		return
	}

	eigenPodProofs, err := eigenpodproofs.NewEigenPodProofs(uint64(ChainId), 60)
	if err != nil {
		t.Logf("%v", err)
		return
	}

	submitter := txsubmitter.NewEigenPodProofTxSubmitter(
		*chainClient,
		*eigenPodProofs,
	)
	t.Logf("do verify")
	tx, err := VerifyWithdrawalCredentialsGen2(submitter, oracleStateFile, oracleHeaderFile, common.HexToAddress(PodAddress), []uint64{1702059, 1702060})
	if err != nil {
		t.Logf("%v", err)
		return
	}
	t.Logf("fake:%v", tx.Hash())
	/*
		//send to chain
		opts := getTransOpts("HOLESKY_ACCOUNT_0")
		//tipCap, _ := big.NewInt(0).SetString("1000000000", 0)
		//opts.GasTipCap = tipCap
		gasPrice, _ := big.NewInt(0).SetString("20000000000", 0)
		opts.GasPrice = gasPrice
		opts.GasLimit = 1000000
		caller := bind.NewBoundContract(*tx.To(), abi.ABI{}, provider, provider, provider)
		realTx, err := caller.RawTransact(opts, tx.Data())
		if err != nil {
			t.Logf("RawTransact err:%v", err)
			return
		}
		t.Logf("raw tx:%v", realTx.Hash())
		txReceipt, err := bind.WaitMined(context.Background(), provider, realTx)
		if err != nil {
			t.Logf("txReceipt err:%v", err)
			return
		}
		gasFee := big.NewInt(0).Mul(big.NewInt(int64(txReceipt.GasUsed)), txReceipt.EffectiveGasPrice)
		t.Logf("gasFee:%v", gasFee)
		egAbi, err := EigenPod.EigenPodMetaData.GetAbi()
		if err != nil {
			t.Error("GetAbi err:", err)
			return
		}
		for _, log := range txReceipt.Logs {
			if log.Address == common.HexToAddress(PodAddress) {
				e, _ := egAbi.EventByID(log.Topics[0])
				if e.Name == "ValidatorRestaked" {
					r, err := contract.ParseValidatorRestaked(*log)
					if err != nil {
						t.Logf("ParseValidatorRestaked err:%v", err)
						return
					}
					t.Logf("r:%v", r.ValidatorIndex)
				}
			}
		}
	*/
}

func getTransOpts(account string) *bind.TransactOpts {
	privateKey := os.Getenv(account)
	if privateKey == "" {
		panic(fmt.Sprintf("no account: %s", account))
	}
	if strings.HasPrefix(privateKey, "0x") {
		privateKey = strings.TrimPrefix(privateKey, "0x")
	}
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err.Error())
	}
	transOpts, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, big.NewInt(ChainId))
	if err != nil {
		panic(err.Error())
	}
	return transOpts
}
