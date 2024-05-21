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
	//jsonByte, _ := json.Marshal(&deposits)
	//t.Logf("deposite:%v", string(jsonByte))
	for _, v := range deposits {
		t.Logf("%v", common.BytesToAddress(v.Data.WithdrawalCredentials[12:]))
		t.Logf("publicKey:%v", fmt.Sprintf("%#x", v.Data.PublicKey))
		//t.Logf("publicKey:%v", hex.EncodeToString(v.Data.PublicKey[:]))
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
	//jsonByte, _ := json.Marshal(&deposits)
	//t.Logf("deposite:%v", string(jsonByte))
	for _, v := range withdrawals {
		t.Logf("index:%v, address:%v, amount:%d", v.Index, common.BytesToAddress(v.Address[:]), v.Amount)
		//t.Logf("publicKey:%v", hex.EncodeToString(v.Data.PublicKey[:]))
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
		State:   "1592408",
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
				//panic("empty slot")
				t.Errorf("empty slot")
				return
			case 503:
				//panic("node is syncing")
			}
		}
		//t.Fatal(err)
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
		Block: "finalized",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("header:%v", header)
}
