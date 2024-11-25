package scanner

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"

	"github.com/Bedrock-Technology/regen3/beaconClient"
	"github.com/Bedrock-Technology/regen3/contracts/Restaking"
	"github.com/Bedrock-Technology/regen3/contracts/Staking"
	"github.com/Bedrock-Technology/regen3/models"
	"github.com/attestantio/go-eth2-client/api"
	v1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

func (s *Scanner) GetAllPod() {
	contract, err := Restaking.NewRestaking(common.HexToAddress(s.Config.RestakingContract), s.EthClient)
	if err != nil {
		logrus.Errorln("NewRestaking err", err)
		return
	}

	podsNum, err := contract.GetTotalPods(&bind.CallOpts{})
	if err != nil {
		logrus.Errorln("GetTotalPods ", err)
	}
	logrus.Infof("podsNum [%v]", podsNum.Uint64())

	pods := make([]models.Pod, 0, podsNum.Uint64())
	for i := uint64(0); i < podsNum.Uint64(); i++ {
		podInfo, err1 := contract.GetPod(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err1 != nil {
			logrus.Errorln("podInfo err:", podInfo)
			return
		}
		podOwner, err2 := contract.PodOwners(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err2 != nil {
			logrus.Errorln("PodOwners err:", podInfo)
			return
		}
		pods = append(pods, models.Pod{
			PodIndex: i,
			Address:  podInfo.String(),
			Owner:    podOwner.String(),
		})
		logrus.Infof("pod %d, podAddress: %s, podOwner: %s", i, podInfo.String(), podOwner.String())
	}
}

func (s *Scanner) Setup(slot string) {
	contract, err := Restaking.NewRestaking(common.HexToAddress(s.Config.RestakingContract), s.EthClient)
	if err != nil {
		logrus.Errorln("NewRestaking err", err)
		return
	}

	podsNum, err := contract.GetTotalPods(&bind.CallOpts{})
	if err != nil {
		logrus.Errorln("GetTotalPods ", err)
	}
	logrus.Infof("podsNum [%v]", podsNum.Uint64())

	pods := make([]models.Pod, 0, podsNum.Uint64())
	for i := uint64(0); i < podsNum.Uint64(); i++ {
		podInfo, err1 := contract.GetPod(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err1 != nil {
			logrus.Errorln("podInfo err:", podInfo)
			return
		}
		podOwner, err2 := contract.PodOwners(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err2 != nil {
			logrus.Errorln("PodOwners err:", podInfo)
			return
		}
		pods = append(pods, models.Pod{
			PodIndex: i,
			Address:  podInfo.String(),
			Owner:    podOwner.String(),
		})
		logrus.Infof("pod %d, podAddress: %s, podOwner: %s", i, podInfo.String(), podOwner.String())
	}

	podsMap := make(map[uint64]models.Pod, len(pods))
	for _, v := range pods {
		podsMap[v.PodIndex] = v
	}

	stakeContract, err := Staking.NewStaking(common.HexToAddress(s.Config.StakingContract), s.EthClient)
	if err != nil {
		logrus.Errorln("NewStaking err:", err)
		return
	}

	nextId, err := stakeContract.GetNextValidatorId(&bind.CallOpts{})
	if err != nil {
		logrus.Errorln("GetNextValidatorId err:")
		return
	}
	logrus.Infof("nextId:%v", nextId.Uint64())

	validators := make([]models.Validator, 0, nextId.Uint64())
	for i := uint64(0); i < nextId.Uint64(); i++ {
		info, err := stakeContract.ValidatorRegistry(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			logrus.Errorln("ValidatorRegistry err:")
			return
		}
		logrus.Infof("pubkey:%v, podAddress: %v", hex.EncodeToString(info.Pubkey), podsMap[uint64(info.Eigenpod)].Address)

		if info.Restaking {
			validator := models.Validator{
				PubKey:     fmt.Sprintf("0x%s", hex.EncodeToString(info.Pubkey)),
				PodAddress: podsMap[uint64(info.Eigenpod)].Address,
			}

			pubKeyBls := phase0.BLSPubKey{}
			_ = pubKeyBls.UnmarshalJSON([]byte(fmt.Sprintf(`"%s"`, validator.PubKey)))
		Retry:
			validatorBeacon, err := s.BeaconClient.Validators(beaconClient.CTX, &api.ValidatorsOpts{
				State:   slot,
				PubKeys: []phase0.BLSPubKey{pubKeyBls},
			})
			if err != nil {
				logrus.Errorf("get %s error:%v", validator.PubKey, err)
				goto Retry
			}

			if len(validatorBeacon.Data) != 0 {
				for _, v := range validatorBeacon.Data {
					if v.Validator.PublicKey.String() == validator.PubKey && v.Status != v1.ValidatorStateWithdrawalDone {
						validator.ValidatorIndex = uint64(v.Index)
						validators = append(validators, validator)
					}
				}
			}
		}
	}

	trans := s.DBEngine.Begin()
	if err := trans.CreateInBatches(&pods, 50).Error; err != nil {
		trans.Rollback()
		logrus.Errorln("Insert error:", err)
		return
	}
	if err := trans.CreateInBatches(&validators, 50).Error; err != nil {
		logrus.Errorln("Insert error:", err)
		trans.Rollback()
		return
	}

	u, _ := strconv.ParseUint(slot, 0, 64)
	if err := trans.Create(&models.Cursor{Slot: u + 1, Meme: models.Scanner}).Error; err != nil {
		logrus.Errorln("Insert cursor error:", err)
		trans.Rollback()
		return
	}
	trans.Commit()
	logrus.Info("done, modify IsCredential manually")
}
