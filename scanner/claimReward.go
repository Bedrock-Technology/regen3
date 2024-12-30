package scanner

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"

	"github.com/Bedrock-Technology/regen3/contracts/Restaking"
)

const RewardProofDir = "./rewardProofs"

func (s *Scanner) claimRewardByPod(podIndex int64, proof *Restaking.IRewardsCoordinatorRewardsMerkleClaim) error {
	restakingAbi, _ := Restaking.RestakingMetaData.GetAbi()
	input, err := restakingAbi.Pack("processClaim", big.NewInt(podIndex), proof)
	if err != nil {
		logrus.Errorln("Pack err:", err)
		return err
	}
	realTx, err := s.sendRawTransaction(input, s.Config.RestakingContract, uint64(podIndex), TxDelegateTo)
	if err != nil {
		return err
	}
	logrus.Infoln("waiting claimReward tx:", realTx.Hash())
	txReceipt, err := bind.WaitMined(context.Background(), s.EthClient, realTx)
	if err != nil {
		logrus.Errorf("waiting claimReward index %v error:%v", podIndex, err)
		panic("waiting error")
	}
	logrus.WithField("Report", "true").Infof("pod[%d] %s tx:%s", podIndex, TxClaimReward, txReceipt.TxHash)
	if err := writeTransaction(s.DBEngine, txReceipt, TxClaimReward); err != nil {
		logrus.Errorf("writeTransaction err:%v", err)
		return err
	}

	// if err := checkIfClaimContained(txReceipt.Logs, s); err != nil {
	// 	logrus.Errorf("checkIfClaimContained error:%v", err)
	// 	panic("checkIfClaimContained")
	// }

	logrus.Infof("staker:%v claimReward", podIndex)
	return nil
}

func (s *Scanner) ClaimRewardWithProof(proofs []*Restaking.IRewardsCoordinatorRewardsMerkleClaim, podIndexes []int64) {
	for k, v := range proofs {
		err := s.claimRewardByPod(podIndexes[k], v)
		if err != nil {
			if errors.Is(err, errBaseFeeTooHigh) {
				logrus.Warnf("%s pod[%d] error:%v", TxClaimReward, podIndexes[k], errBaseFeeTooHigh)
				return
			}
			logrus.Errorf("send claimRewardByPod pod[%d] error:%v", podIndexes[k], err)
			panic("claimReward err")
		}
		// delete file
		fileName := fmt.Sprintf("pod_%d.json", podIndexes[k])
		filePath := fmt.Sprintf("%s/%s", RewardProofDir, fileName)
		err = os.Remove(filePath)
		if err != nil {
			logrus.Errorf("remove %v error:%v", filePath, err)
			panic("remove error")
		}
		logrus.Infof("delete file:%s", filePath)
	}
}

func (s *Scanner) ClaimReward() {
	proofFiles := GetAllProofFile(RewardProofDir)
	if len(proofFiles) == 0 {
		logrus.Infoln("empty claimProof")
		return
	}

	for k, v := range proofFiles {
		proofFiles[k] = fmt.Sprintf("%s/%s", RewardProofDir, v)
	}

	proofs, podIndexes := GetProofs(proofFiles)
	if len(proofs) == 0 || len(podIndexes) == 0 || len(proofs) != len(podIndexes) {
		logrus.Infoln("empty proofs or podIndexes or not equal")
		return
	}

	proofs, podIndexes = FilterProofs(proofs, podIndexes, s)

	for k, v := range proofs {
		err := s.claimRewardByPod(podIndexes[k], v)
		if err != nil {
			if errors.Is(err, errBaseFeeTooHigh) {
				logrus.Warnf("%s pod[%d] error:%v", TxClaimReward, podIndexes[k], errBaseFeeTooHigh)
				return
			}
			logrus.Errorf("send claimRewardByPod pod[%d] error:%v", podIndexes[k], err)
			panic("claimReward err")
		}
		// delete file
		fileName := fmt.Sprintf("pod_%d.json", podIndexes[k])
		filePath := fmt.Sprintf("%s/%s", RewardProofDir, fileName)
		err = os.Remove(filePath)
		if err != nil {
			logrus.Errorf("remove %v error:%v", filePath, err)
			panic("remove error")
		}
	}
}

func FilterProofs(proofs []*Restaking.IRewardsCoordinatorRewardsMerkleClaim, podIndexes []int64, s *Scanner) ([]*Restaking.IRewardsCoordinatorRewardsMerkleClaim, []int64) {
	filterdProofs := []*Restaking.IRewardsCoordinatorRewardsMerkleClaim{}
	filterdPodIndexes := []int64{}
	for k, v := range proofs {
		del := bool(true)
		for _, tokenLeave := range v.TokenLeaves {
			if tokenLeave.Token.String() == s.Config.EigenToken {
				eigenThreshold := big.NewInt(int64(s.Config.EigenTokenThreshold))
				eigenThreshold.Mul(eigenThreshold, big.NewInt(1e9))
				logrus.Infof("eigenThreshold[%v], earnings[%v]", eigenThreshold, tokenLeave.CumulativeEarnings)
				if tokenLeave.CumulativeEarnings.Cmp(eigenThreshold) > 0 {
					filterdProofs = append(filterdProofs, proofs[k])
					filterdPodIndexes = append(filterdPodIndexes, podIndexes[k])
					del = false
				}
			}
		}
		if del {
			logrus.Infof("tokenLeaves skip and delete %d", podIndexes[k])
			// delete file
			fileName := fmt.Sprintf("pod_%d.json", podIndexes[k])
			filePath := fmt.Sprintf("%s/%s", RewardProofDir, fileName)
			err := os.Remove(filePath)
			if err != nil {
				logrus.Errorf("remove %v error:%v", filePath, err)
				panic("remove error")
			}
			logrus.Infof("delete file:%s", filePath)
		}
	}
	return filterdProofs, filterdPodIndexes
}

func GetProofs(proofFiles []string) (proofs []*Restaking.IRewardsCoordinatorRewardsMerkleClaim, podIndexes []int64) {
	for _, v := range proofFiles {
		fileBytes, err := os.ReadFile(v)
		if err != nil {
			logrus.Errorf("read file %v err:%v", v, err)
			continue
		}

		merkleClaimStrings := IRewardsCoordinatorRewardsMerkleClaimStrings{}
		err = json.Unmarshal(fileBytes, &merkleClaimStrings)
		if err != nil {
			logrus.Errorf("json Unmarshal error:%v", err)
			continue
		}
		merkleClaim, err := convertToMerkleClaim(&merkleClaimStrings)
		if err != nil {
			logrus.Errorf("convertToMerkleClaim error:%v", err)
			continue
		}

		splites := strings.Split(v, "_")
		if len(splites) != 2 {
			logrus.Errorf("splite %v error", v)
			continue
		}
		ssplites := strings.Split(splites[1], ".")
		if len(ssplites) != 2 {
			logrus.Errorf("ssplite %v error", ssplites)
			continue
		}
		podIndex, err := strconv.Atoi(ssplites[0])
		if err != nil {
			logrus.Errorf("convert %v error", ssplites[0])
			continue
		}
		proofs = append(proofs, merkleClaim)
		podIndexes = append(podIndexes, int64(podIndex))
	}
	return proofs, podIndexes
}

func GetAllProofFile(path string) (proofFiles []string) {
	files, err := os.ReadDir(path)
	if err != nil {
		logrus.Errorf("ReadDir err:%v", err)
		return nil
	}
	if len(files) == 0 {
		logrus.Infoln("ReadDir empty")
		return nil
	}

	for _, v := range files {
		if v.IsDir() {
			continue
		}
		splite := strings.Split(v.Name(), ".")
		if len(splite) != 2 || splite[1] != "json" {
			continue
		}
		proofFiles = append(proofFiles, v.Name())
	}

	return proofFiles
}

// func checkIfClaimContained(logs []*types.Log, s *Scanner) error {
// 	egAbi, _ := EigenPod.EigenPodMetaData.GetAbi()
// 	for _, log := range logs {
// 		if log.Address == common.HexToAddress(s.Config.RewardCoordinator) {
// 			e, err := egAbi.EventByID(log.Topics[0])
// 			if err != nil {
// 				return err
// 			}
// 			if e.Name == "RewardsClaimed" {
// 				return nil
// 			}
// 		}
// 	}
// 	return errors.New("not found")
// }

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

func convertToMerkleClaim(merkleClaimStrings *IRewardsCoordinatorRewardsMerkleClaimStrings) (*Restaking.IRewardsCoordinatorRewardsMerkleClaim, error) {
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

	tokenLeaves := []Restaking.IRewardsCoordinatorTokenTreeMerkleLeaf{}
	for _, v := range merkleClaimStrings.TokenLeaves {
		earnings, bool := big.NewInt(0).SetString(v.CumulativeEarnings, 0)
		if !bool {
			return nil, errors.New("string to format error")
		}
		tokenLeave := Restaking.IRewardsCoordinatorTokenTreeMerkleLeaf{
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

	merkleClaim := Restaking.IRewardsCoordinatorRewardsMerkleClaim{
		RootIndex:       merkleClaimStrings.RootIndex,
		EarnerIndex:     merkleClaimStrings.EarnerIndex,
		EarnerTreeProof: earnerTreeProof,
		EarnerLeaf: Restaking.IRewardsCoordinatorEarnerTreeMerkleLeaf{
			Earner:          merkleClaimStrings.EarnerLeaf.Earner,
			EarnerTokenRoot: earnerTokenRoot,
		},
		TokenIndices:    merkleClaimStrings.TokenIndices,
		TokenTreeProofs: tokenTreeProofs,
		TokenLeaves:     tokenLeaves,
	}
	return &merkleClaim, nil
}
