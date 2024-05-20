package DelegationManager

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/joho/godotenv"

	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var provider *ethclient.Client
var contract *DelegationManager

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
	ContractAddress = os.Getenv("HOLESKY_DELEGATION_MANAGER_CONTRACT")
	PodOwner = os.Getenv("HOLESKY_POD_WONER")

	client, err := ethclient.Dial(RpcHost)
	if err != nil {
		panic(err)
	}
	d, err := NewDelegationManager(common.HexToAddress(ContractAddress), client)
	if err != nil {
		panic(err)
	}
	provider = client
	contract = d
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

func TestNewDelegationManagerEventsName(t *testing.T) {
	dmAbi, _ := DelegationManagerMetaData.GetAbi()
	for _, v := range dmAbi.Events {
		t.Log("name:", v.Name)
	}
}

func TestDelegationManagerCaller_StakerNonce(t *testing.T) {
	nonce, err := contract.StakerNonce(&bind.CallOpts{
		Pending:     false,
		From:        common.HexToAddress(PodOwner),
		BlockNumber: nil,
		BlockHash:   common.Hash{},
		Context:     nil,
	}, common.HexToAddress(PodOwner))
	if err != nil {
		t.Error("call error:", err)
		return
	}
	t.Log("nonce:", nonce)
}

func TestDelegationManagerTransactor_CompleteQueuedWithdrawals(t *testing.T) {
	//blockNum := uint64(1442069)
	blockNum := uint64(1495178)
	//=== RUN   TestDelegationManagerTransactor_CompleteQueuedWithdrawals
	//    DelegationManager_test.go:130: CompleteQueuedWithdrawals err: execution reverted: DelegationManager._completeQueuedWithdrawal: minWithdrawalDelayBlocks period has not yet passed
	it, err := contract.FilterWithdrawalQueued(&bind.FilterOpts{
		Start:   blockNum,
		End:     &blockNum,
		Context: nil,
	})
	if err != nil {
		t.Error("err:", err)
		return
	}
	var idmw []IDelegationManagerWithdrawal
	for it.Next() {
		//jsonString, _ := json.Marshal(it.Event)
		//t.Log("event", string(jsonString))
		idmw = append(idmw, it.Event.Withdrawal)
	}
	if it.Error() != nil {
		t.Error("it err:", it.Error())
		return
	}
	var tokens [][]common.Address
	tokens = append(tokens, []common.Address{common.HexToAddress("0x0000000000000000000000000000000000000000")})
	middlewareTimesIndexes := []*big.Int{big.NewInt(0)}
	transOpts := getTransOpts("HOLESKY_ACCOUNT_0")
	transOpts.NoSend = true
	dmAbi, _ := DelegationManagerMetaData.GetAbi()
	data, _ := dmAbi.Pack("completeQueuedWithdrawals", idmw, tokens, middlewareTimesIndexes, []bool{false})
	estimateGasLimitPack, err := provider.EstimateGas(context.Background(), ethereum.CallMsg{
		From: common.HexToAddress(PodOwner),
		To:   func(address common.Address) *common.Address { return &address }(common.HexToAddress(ContractAddress)),
		Data: data,
	})
	if err != nil {
		t.Logf("estimate gas limit err: %v", err)
		return
	}
	t.Logf("estimateGasLimitPack:%v,%v", estimateGasLimitPack, hex.EncodeToString(data))
	//estimateGasLimitPack:90513,
	tx, err := contract.CompleteQueuedWithdrawals(transOpts, idmw, tokens, middlewareTimesIndexes, []bool{false})
	if err != nil {
		t.Errorf("CompleteQueuedWithdrawals err:%v", err)
		return
	}
	estimateGasLimit, err := provider.EstimateGas(context.Background(), ethereum.CallMsg{
		From: transOpts.From,
		To:   tx.To(),
		Data: tx.Data(),
	})
	t.Logf("estimateGasLimit: %v, data: %v", estimateGasLimit, hex.EncodeToString(tx.Data()))
	//estimateGasLimit: 90513
	t.Log("hash:", tx.Hash().String())
	transOpts.NoSend = false
	transOpts.GasLimit = estimateGasLimit * 4
	txReal, err := contract.CompleteQueuedWithdrawals(transOpts, idmw, tokens, middlewareTimesIndexes, []bool{false})
	if err != nil {
		t.Error("CompleteQueuedWithdrawals err:", err)
		return
	}
	t.Logf("txReal: %v, data: %v", txReal.Hash().String(), hex.EncodeToString(txReal.Data()))
}

func TestDelegationManagerTransactor_DelegateTo(t *testing.T) {
	transOpts := getTransOpts("HOLESKY_ACCOUNT_0")
	operator := common.HexToAddress("0xA275A0a402874Addd807C479487182431D42E397")
	isswe := ISignatureUtilsSignatureWithExpiry{
		Signature: []byte{},
		Expiry:    big.NewInt(0),
	}
	approverSalt := [32]byte{0}
	tx, err := contract.DelegateTo(transOpts, operator,
		isswe, approverSalt)
	if err != nil {
		t.Error("CompleteQueuedWithdrawals err:", err)
		return
	}
	txReceipt, err := bind.WaitMined(context.Background(), provider, tx)
	if err != nil {
		t.Error("mine err:", err)
		return
	}
	t.Log("txReceipt:", txReceipt.TxHash.String())
}

func TestDelegationManagerCaller_DelegatedTo(t *testing.T) {
	op, err := contract.DelegatedTo(&bind.CallOpts{}, common.HexToAddress(PodOwner))
	if err != nil {
		t.Error("err:", err)
		return
	}
	t.Log("operator:", op.String())
}

func TestDelegationManagerCaller_GetDelegatableShares(t *testing.T) {
	add, amount, err := contract.GetDelegatableShares(&bind.CallOpts{}, common.HexToAddress(PodOwner))
	if err != nil {
		t.Error("err:", err)
		return
	}
	t.Logf("add:%v,amount:%v", add, amount)
}

func TestNewDelegationManager_FilterLogs(t *testing.T) {
	logs, err := provider.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(1448012),
		ToBlock:   big.NewInt(1448094),
		Addresses: []common.Address{common.HexToAddress(ContractAddress)},
		Topics:    nil,
	})
	if err != nil {
		t.Error("err:", err)
		return
	}
	dmAbi, err := DelegationManagerMetaData.GetAbi()
	if err != nil {
		t.Error("GetAbi err:", err)
		return
	}
	//logsJson, _ := json.Marshal(&logs)
	//t.Log("logs:", string(logsJson))
	for _, log := range logs {
		t.Log("topic:", log.Topics[0].Hex())
		t.Log("txHash:", log.TxHash.String())
		event, err := dmAbi.EventByID(log.Topics[0])
		if err != nil {
			t.Error("EventByID err:", err)
			return
		}
		t.Log("event:", event.Name)
		switch event.Name {
		case "WithdrawalCompleted":
			wc, err := contract.ParseWithdrawalCompleted(log)
			if err != nil {
				t.Error("ParseWithdrawalCompleted err:", err)
				return
			}
			t.Log("wc:", wc)
		case "OperatorMetadataURIUpdated":
			opUpdate, err := contract.ParseOperatorMetadataURIUpdated(log)
			if err != nil {
				t.Error("ParseOperatorMetadataURIUpdated err:", err)
				return
			}
			t.Log("opUpdate:", opUpdate)
		case "OperatorRegistered":
			opReg, err := contract.ParseOperatorRegistered(log)
			if err != nil {
				t.Error("ParseOperatorRegistered err:", err)
				return
			}
			t.Log("opReg:", opReg)
		case "OperatorSharesDecreased":
			opShareDec, err := contract.ParseOperatorSharesDecreased(log)
			if err != nil {
				t.Error("ParseOperatorSharesDecreased err:", err)
				return
			}
			t.Log("opShareDec:", opShareDec)
		case "StrategyWithdrawalDelayBlocksSet":
			swdbs, err := contract.ParseStrategyWithdrawalDelayBlocksSet(log)
			if err != nil {
				t.Error("ParseStrategyWithdrawalDelayBlocksSet err:", err)
				return
			}
			t.Log("swdbs:", swdbs)
		case "StakerUndelegated":
			stakerUndelegated, err := contract.ParseStakerUndelegated(log)
			if err != nil {
				t.Error("ParseStakerUndelegated err:", err)
				return
			}
			t.Log("stakerUndelegated:", stakerUndelegated)
		case "WithdrawalMigrated":
			wm, err := contract.ParseWithdrawalMigrated(log)
			if err != nil {
				t.Error("ParseWithdrawalMigrated err:", err)
				return
			}
			t.Log("WithdrawalMigrated:", wm)
		case "OperatorDetailsModified":
			odm, err := contract.ParseOperatorDetailsModified(log)
			if err != nil {
				t.Error("ParseOperatorDetailsModified err:", err)
				return
			}
			t.Log("odm:", odm)
		case "OperatorSharesIncreased":
			osi, err := contract.ParseOperatorSharesIncreased(log)
			if err != nil {
				t.Error("ParseOperatorSharesIncreased err:", err)
				return
			}
			t.Log("osi:", osi)
		case "OwnershipTransferred":
			ot, err := contract.ParseOwnershipTransferred(log)
			if err != nil {
				t.Error("ParseOwnershipTransferred err:", err)
				return
			}
			t.Log("ot:", ot)
		case "StakerForceUndelegated":
			sfu, err := contract.ParseStakerForceUndelegated(log)
			if err != nil {
				t.Error("ParseStakerForceUndelegated err:", err)
				return
			}
			t.Log("sfu:", sfu)
		case "MinWithdrawalDelayBlocksSet":
			mwdbs, err := contract.ParseMinWithdrawalDelayBlocksSet(log)
			if err != nil {
				t.Error("ParseMinWithdrawalDelayBlocksSet err:", log)
				return
			}
			t.Log("mwdbs:", mwdbs)
		case "PauserRegistrySet":
			prs, err := contract.ParsePauserRegistrySet(log)
			if err != nil {
				t.Error("ParsePauserRegistrySet err:", log)
				return
			}
			t.Log("prs:", prs)
		case "Unpaused":
			unpaused, err := contract.ParseUnpaused(log)
			if err != nil {
				t.Error("ParseUnpaused err:", log)
				return
			}
			t.Log("unpaused:", unpaused)
		case "WithdrawalQueued":
			wq, err := contract.ParseWithdrawalQueued(log)
			if err != nil {
				t.Error("ParseWithdrawalQueued err:", log)
				return
			}
			t.Log("wq:", wq)
		case "Initialized":
			init, err := contract.ParseInitialized(log)
			if err != nil {
				t.Error("ParseInitialized err:", log)
				return
			}
			t.Log("init:", init)
		case "Paused":
			paused, err := contract.ParsePaused(log)
			if err != nil {
				t.Error("ParsePaused err:", log)
				return
			}
			t.Log("paused:", paused)
		case "StakerDelegated":
			stakerD, err := contract.ParseStakerDelegated(log)
			if err != nil {
				t.Error("ParseStakerDelegated err:", log)
				return
			}
			stakerDJson, _ := json.Marshal(&stakerD)
			t.Log("stakerD:", string(stakerDJson))
		}
	}
}

func TestDelegationManagerTransactor_QueueWithdrawals(t *testing.T) {
	shares, _ := big.NewInt(0).SetString("25000000000000000000", 0)
	qwp := IDelegationManagerQueuedWithdrawalParams{
		Strategies: []common.Address{common.HexToAddress("0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0")},
		Shares:     []*big.Int{shares},
		Withdrawer: common.HexToAddress(PodOwner),
	}
	qwps := []IDelegationManagerQueuedWithdrawalParams{qwp}
	opts := getTransOpts("HOLESKY_ACCOUNT_0")
	tx, err := contract.QueueWithdrawals(opts, qwps)
	if err != nil {
		t.Error("QueueWithdrawals err:", err)
		return
	}
	txReceipt, err := bind.WaitMined(context.Background(), provider, tx)
	if err != nil {
		t.Error("WaitMined err:", err)
		return
	}
	//GetDelegatableShares Decreased
	t.Log("txReceipt:", txReceipt)
	//parseLog
	if txReceipt.Status != 1 {
		t.Log("status:", txReceipt.Status)
		return
	}
	dmAbi, err := DelegationManagerMetaData.GetAbi()
	if err != nil {
		t.Error("GetAbi err:", err)
		return
	}

	for _, log := range txReceipt.Logs {
		event, err := dmAbi.EventByID(log.Topics[0])
		if err != nil {
			t.Error("EventByID err:", err)
			return
		}
		if event.Name == "OperatorSharesDecreased" {
			osd, err := contract.ParseOperatorSharesDecreased(*log)
			if err != nil {
				t.Error("ParseOperatorSharesDecreased err:", err)
				return
			}
			t.Logf("osd Operator:%v, Staker:%v, Stratedy:%v, Shares:%v", osd.Operator, osd.Staker, osd.Strategy, osd.Shares)
		} else if event.Name == "WithdrawalQueued" {
			//where to query WithdrawalQueued?? PendingWithdrawals only return true of false
			wq, err := contract.ParseWithdrawalQueued(*log)
			if err != nil {
				t.Error("ParseWithdrawalQueued err:", err)
				return
			}
			withdrawlJson, _ := json.Marshal(&wq.Withdrawal)
			t.Logf("wq withdrawalRoot:%v, Withdrawal:%v", string(wq.WithdrawalRoot[:]), string(withdrawlJson))
			//Withdrawal:{"Staker":"0x7f1fba3d5572a8a267c236d28b492f388c4db002","DelegatedTo":"0xa275a0a402874addd807c479487182431d42e397","Withdrawer":"0x7f1fba3d5572a8a267c236d28b492f388c4db002","Nonce":11,"StartBlock":1486674,"Strategies":["0xbeac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0"],"Shares":[15000000000000000000]}
		} else {
			t.Log("Not find:", event)
		}
	}
}
