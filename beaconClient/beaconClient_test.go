package beaconClient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/attestantio/go-eth2-client/api"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

var client *Client

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	url := os.Getenv("HOLESKY_BEACON_URL")
	client, _ = NewClient(url)
}

func TestGetBlock(t *testing.T) {
	block, err := client.SignedBeaconBlock(context.Background(), &api.SignedBeaconBlockOpts{
		Block: "finalized",
	})
	if err != nil {
		t.Fatal(err)
	}
	jsonByte, _ := json.Marshal(&block)
	t.Log(string(jsonByte))
}

func TestGetBlockEmptySlot(t *testing.T) {
	block, err := client.SignedBeaconBlock(context.Background(), &api.SignedBeaconBlockOpts{
		Block: "1662790",
	})
	if err != nil {
		var apiErr *api.Error
		if errors.As(err, &apiErr) {
			switch apiErr.StatusCode {
			case 404:
				panic("empty slot")
			case 503:
				panic("node is syncing")
			}
		}
		t.Fatal(err)
	}
	jsonByte, _ := json.Marshal(&block)
	t.Log(string(jsonByte))
}

func TestGetBlockDepositsSlot(t *testing.T) {
	block, err := client.SignedBeaconBlock(context.Background(), &api.SignedBeaconBlockOpts{
		Block: "1592408",
	})
	if err != nil {
		var apiErr *api.Error
		if errors.As(err, &apiErr) {
			switch apiErr.StatusCode {
			case 404:
				panic("empty slot")
			case 503:
				panic("node is syncing")
			}
		}
		t.Fatal(err)
	}
	eblock, err := block.Data.ExecutionBlockNumber()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("eblock:%v", eblock)
	deposits, _ := block.Data.Deposits()
	// jsonByte, _ := json.Marshal(&deposits)
	// t.Logf("deposite:%v", string(jsonByte))
	for _, v := range deposits {
		t.Logf("%v", common.BytesToAddress(v.Data.WithdrawalCredentials[12:]))
		t.Logf("publicKey:%v", fmt.Sprintf("%#x", v.Data.PublicKey))
		// t.Logf("publicKey:%v", hex.EncodeToString(v.Data.PublicKey[:]))
	}
}

func TestGetBlockWithdrawalSlot(t *testing.T) {
	block, err := client.SignedBeaconBlock(context.Background(), &api.SignedBeaconBlockOpts{
		Block: "1592408",
	})
	if err != nil {
		var apiErr *api.Error
		if errors.As(err, &apiErr) {
			switch apiErr.StatusCode {
			case 404:
				panic("empty slot")
			case 503:
				panic("node is syncing")
			}
		}
		t.Fatal(err)
	}
	eblock, err := block.Data.ExecutionBlockNumber()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("eblock:%v", eblock)
	withdrawals, _ := block.Data.Withdrawals()
	// jsonByte, _ := json.Marshal(&deposits)
	// t.Logf("deposite:%v", string(jsonByte))
	for _, v := range withdrawals {
		t.Logf("index:%v, address:%v, amount:%d", v.Index, common.BytesToAddress(v.Address[:]), v.Amount)
		// t.Logf("publicKey:%v", hex.EncodeToString(v.Data.PublicKey[:]))
	}
}

func TestGetValidatorStateSlot(t *testing.T) {
	pubKey := fmt.Sprintf(`"%s"`, "0xaee4bc69d8ea4997cea56fb7aa8a56af42a4f8c41105651b6deb316ebc436a5b43f7c6657401fdee88b2e891d0cb42a3")
	pubKeyBls := phase0.BLSPubKey{}
	err := pubKeyBls.UnmarshalJSON([]byte(pubKey))
	if err != nil {
		t.Fatal("err:", err)
	}
	v, err := client.Validators(context.TODO(), &api.ValidatorsOpts{
		// State:   "1592408",
		State:   "head",
		PubKeys: []phase0.BLSPubKey{pubKeyBls},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", v)
	for _, vali := range v.Data {
		t.Logf("pubKey:%s, index:%d", vali.Validator.PublicKey, vali.Index)
	}
}

func TestGetVoluntaryExitInBlock(t *testing.T) {
	Block := 1510910
	block, err := client.SignedBeaconBlock(context.Background(), &api.SignedBeaconBlockOpts{
		Block: fmt.Sprintf("%d", Block),
	})
	if err != nil {
		var apiErr *api.Error
		if errors.As(err, &apiErr) {
			switch apiErr.StatusCode {
			case 404:
				// panic("empty slot")
				t.Errorf("empty slot")
				return
			case 503:
				// panic("node is syncing")
			}
		}
		// t.Fatal(err)
	}
	v, err := block.Data.VoluntaryExits()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("v:%s", v)
	for _, vv := range v {
		t.Logf("block:%v,vv:%v", Block, vv.Message.ValidatorIndex)
	}
}

func TestHeader(t *testing.T) {
	header, err := client.BeaconBlockHeader(context.Background(), &api.BeaconBlockHeaderOpts{
		Block: "head",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("header:%v", header)
}

func TestEmptyHeader(t *testing.T) {
	header, err := client.BeaconBlockHeader(context.Background(), &api.BeaconBlockHeaderOpts{
		Block: fmt.Sprintf("%d", 1662790),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("header:%v", header)
}

func TestNotDepositPubKey(t *testing.T) {
	pubKeyS := `"0x8cc013a19a332b806908437217bf8ae5eb6898dd72125c88d9c3faf3c3c2ffb84b6418395dc59c0ec2012bc454711dd1"`
	pubKeyBls := phase0.BLSPubKey{}
	_ = pubKeyBls.UnmarshalJSON([]byte(pubKeyS))
	validatorOnbeacon, err := client.Validators(context.Background(), &api.ValidatorsOpts{
		State:   "head",
		PubKeys: []phase0.BLSPubKey{pubKeyBls},
	})
	if err != nil {
		t.Logf("err:%v", err)
		return
	}
	t.Logf("validatorOnbeacon:%v", validatorOnbeacon)
	t.Logf("validatorOnbeacon:%v", len(validatorOnbeacon.Data))
	t.Logf("validatorOnbeacon:%v", validatorOnbeacon.Data[0].Index)
}

func TestNewClient(t *testing.T) {
	header, err := client.GetLatestSlotNumber()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("header:%v", header)
}

func TestValidatorStatus(t *testing.T) {
	validatorIndices := []uint64{
		1796388,
		1799558,
		1799559,
		1799560,
		1799561,
		1799562,
		1799563,
		1799564,
		1799565,
		1799566,
	}
	var vi []phase0.ValidatorIndex
	for _, v := range validatorIndices {
		vi = append(vi, phase0.ValidatorIndex(v))
	}

	validatorOnbeacon, err := client.Validators(context.Background(), &api.ValidatorsOpts{
		State:   "head",
		Indices: vi,
	})
	if err != nil {
		t.Logf("err:%v", err)
		return
	}
	fmt.Println(validatorOnbeacon.Data)
}
