package Restaking

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"os"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var (
	provider *ethclient.Client
	contract *Restaking
)

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

type IRewardsCoordinatorEarnerTreeMerkleLeafStrings struct {
	Earner          common.Address `json:"earner"`
	EarnerTokenRoot string         `json:"earnerTokenRoot"`
}

type IRewardsCoordinatorTokenTreeMerkleLeafStrings struct {
	Token              common.Address `json:"token"`
	CumulativeEarnings string         `json:"cumulativeEarnings"`
}

type IRewardsCoordinatorRewardsMerkleClaimStrings struct {
	Root               string                                          `json:"root"`
	RootIndex          uint32                                          `json:"rootIndex"`
	EarnerIndex        uint32                                          `json:"earnerIndex"`
	EarnerTreeProof    string                                          `json:"earnerTreeProof"`
	EarnerLeaf         IRewardsCoordinatorEarnerTreeMerkleLeafStrings  `json:"earnerLeaf"`
	TokenIndices       []uint32                                        `json:"tokenIndices"`
	TokenTreeProofs    []string                                        `json:"tokenTreeProofs"`
	TokenLeaves        []IRewardsCoordinatorTokenTreeMerkleLeafStrings `json:"tokenLeaves"`
	TokenTreeProofsNum uint32                                          `json:"tokenTreeProofsNum"`
	TokenLeavesNum     uint32                                          `json:"tokenLeavesNum"`
}

func convertToMerkleClaim(merkleClaimStrings *IRewardsCoordinatorRewardsMerkleClaimStrings) (*IRewardsCoordinatorRewardsMerkleClaim, error) {
	earnerTreeProof, err := hex.DecodeString(merkleClaimStrings.EarnerTreeProof[2:])
	if err != nil {
		return nil, err
	}

	tokenTreeProofs := [][]byte{}
	for _, v := range merkleClaimStrings.TokenTreeProofs {
		tokenTreeProof, err1 := hex.DecodeString(v[2:])
		if err1 != nil {
			return nil, err1
		}
		tokenTreeProofs = append(tokenTreeProofs, tokenTreeProof)
	}

	tokenLeaves := []IRewardsCoordinatorTokenTreeMerkleLeaf{}
	for _, v := range merkleClaimStrings.TokenLeaves {
		earnings, bool := big.NewInt(0).SetString(v.CumulativeEarnings, 0)
		if !bool {
			return nil, errors.New("string to format error")
		}
		tokenLeave := IRewardsCoordinatorTokenTreeMerkleLeaf{
			Token:              v.Token,
			CumulativeEarnings: earnings,
		}
		tokenLeaves = append(tokenLeaves, tokenLeave)
	}

	earnerTokenRoot := [32]byte{}
	tokenRoot, err := hex.DecodeString(merkleClaimStrings.EarnerLeaf.EarnerTokenRoot[2:])
	if err != nil {
		return nil, err
	}
	copy(earnerTokenRoot[:], tokenRoot)

	merkleClaim := IRewardsCoordinatorRewardsMerkleClaim{
		RootIndex:       merkleClaimStrings.RootIndex,
		EarnerIndex:     merkleClaimStrings.EarnerIndex,
		EarnerTreeProof: earnerTreeProof,
		EarnerLeaf: IRewardsCoordinatorEarnerTreeMerkleLeaf{
			Earner:          merkleClaimStrings.EarnerLeaf.Earner,
			EarnerTokenRoot: earnerTokenRoot,
		},
		TokenIndices:    merkleClaimStrings.TokenIndices,
		TokenTreeProofs: tokenTreeProofs,
		TokenLeaves:     tokenLeaves,
	}
	return &merkleClaim, nil
}

func TestUnmarshl(t *testing.T) {
	jsonProof := `{
	  "root": "0x4344082ad741d467845ee47dbeecccf16236fdde8d601754d39a5d694dde0549",
	  "rootIndex": 21,
	  "earnerIndex": 38789,
	  "earnerTreeProof": "0xfd4e957ecb3b34eed7ff241c2102dc85d82f38e6931eab213fa3d27bfe95945966de87d07dacf9e99f75c3359e54fd5c5f230714e1e27fdbf83b04c6048e8506cbb85c46a6e4413708542e257415179ee4bb58ef7e6f3786d6d2a146bef87ecc03d8c4677c20f1bd51480f8ac21fe83e6c9114f921a29d7dee10997bb618007e0953673628dbd0318bcf64ff89e765e14d7ca6dc7849087f8c821db18994478b1a6bc78b6522288defb6c34348bce8029118b93748eb9c4628abf8c9f03befb4eb388044ccad56ee78318b44209f3f81ccd4cabd4d24ce97e3d00cf57a498ed1895a73d4936d60db67bf8f7456d1f9368cb222a67e06759b376612208114dc394ebf4656814d255061e7c66f362e9a539431637cabe27239906592f8bae15dae53e1cfcaf45ffacb5a02a5971a4dbe894945eceb66a546ba6cee7fa008ae71f631df1a993791e4146296d72d5614a27901ba0bd8690f651e1bfd9878dfa3a8964be61db8346573e6af2a6c86a9dc2b5fb1240b4f225cd4ffa02cfda17076d7b5886558f9b8caf5222aa1036c38fb6dbb89443197e0e9c3b9e078b6d2c41837ade6cab80796468c0ce3f569b37319995462f24cb6c116107b2d70b5ee984a16151d707a0ff8c2af96406b8aa61909ca398d92e3798fc3ab7cd5baf6baf55d450a7444106b46239cb965cfaa7236c0f2ec8e5ea5656cee666ef010f27952d79133d82d102279a9ed60140c85a995ef7e7de0b492cd66e4caf6179366c2720b9c68de1a541ede97802594cd7ae2a7a433ae94b901c34097fc41b364e58146c0faf1",
	  "earnerLeaf": {
	    "earner": "0x3f4eaceb930b0edfa78a1dfcbae5c5494e6e9850",
	    "earnerTokenRoot": "0x00eb6975209027b214b072ab0530b93e2e4862699f13e14f894866a823c7ca21"
	  },
	  "tokenIndices": [
	    0,
	    1
	  ],
	  "tokenTreeProofs": [
	    "0xb2410f3c21b9e234365c9664cc2463407134c3cb6a76ef75c99b8629a572f8dd",
	    "0x8129fc4f1b38cc97d14981cd4ebb3d261f4290b8c42af00ac2c7e69228a8db26"
	  ],
	  "tokenLeaves": [
	    {
	      "token": "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
	      "cumulativeEarnings": "65955484345064147"
	    },
	    {
	      "token": "0xec53bf9167f50cdeb3ae105f56099aaab9061f83",
	      "cumulativeEarnings": "50364936135183671571768"
	    }
	  ],
	  "tokenTreeProofsNum": 2,
	  "tokenLeavesNum": 2
	}`
	merkleClaimString := IRewardsCoordinatorRewardsMerkleClaimStrings{}
	json.Unmarshal([]byte(jsonProof), &merkleClaimString)
	// t.Log(merkleClaimString)
	claim, err := convertToMerkleClaim(&merkleClaimString)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(claim)
}
