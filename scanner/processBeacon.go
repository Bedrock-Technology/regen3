package scanner

import (
	"errors"
	"fmt"
	"github.com/Bedrock-Technology/regen3/beaconClient"
	"github.com/Bedrock-Technology/regen3/models"
	"github.com/attestantio/go-eth2-client/api"
	"github.com/attestantio/go-eth2-client/spec"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (s *Scanner) processBeacon(beaconBlock *api.Response[*spec.VersionedSignedBeaconBlock], orm *gorm.DB) error {
	if err := s.processDeposit(beaconBlock, orm); err != nil {
		return err
	}
	if err := s.processVoluntaryExit(beaconBlock, orm); err != nil {
		return err
	}
	return s.processWithdrawnOnChain(beaconBlock, orm)
}

func (s *Scanner) processDeposit(beaconBlock *api.Response[*spec.VersionedSignedBeaconBlock], orm *gorm.DB) error {
	deposits, err := beaconBlock.Data.Deposits()
	if err != nil {
		return err
	}

	var modelValidators []models.Validator
	for _, deposit := range deposits {
		withdrawal := common.BytesToAddress(deposit.Data.WithdrawalCredentials[12:]).String()
		pubKey := fmt.Sprintf("%#x", deposit.Data.PublicKey)
		if _, exist := s.Pods[withdrawal]; exist {
			modelValidators = append(modelValidators, models.Validator{
				PubKey:             pubKey,
				PodAddress:         withdrawal,
				CredentialVerified: 0,
				WithdrawnOnChain:   0,
				WithdrawnOnPod:     0,
				VoluntaryExit:      0,
			})
		}
	}

	var blsPubKey []phase0.BLSPubKey
	for _, modelValidator := range modelValidators {
		var pubKeyBls phase0.BLSPubKey
		_ = pubKeyBls.UnmarshalJSON([]byte(fmt.Sprintf(`"%s"`, modelValidator.PubKey)))
		blsPubKey = append(blsPubKey, pubKeyBls)
	}

	slot, err := beaconBlock.Data.Slot()
	if err != nil {
		return err
	}

	if len(blsPubKey) != 0 {
		validators, err := s.BeaconClient.Validators(beaconClient.CTX, &api.ValidatorsOpts{
			State:   fmt.Sprintf("%d", slot),
			PubKeys: blsPubKey,
		})
		if err != nil {
			return err
		}
		for _, v := range validators.Data {
			for i, k := range modelValidators {
				if v.Validator.PublicKey.String() == k.PubKey && v.Validator.EffectiveBalance >= phase0.Gwei(32000000000) {
					modelValidators[i].ValidatorIndex = uint64(v.Index)
				}
			}
		}
		logrus.Infof("find Deposit len(%d), slot[%d]", len(blsPubKey), slot)
		result := orm.Create(&modelValidators)
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			logrus.Warnf("duplicated key, slot[%d],error:%v", slot, result.Error)
			return nil
		} else {
			return result.Error
		}
	}
	return nil
}

func (s *Scanner) processVoluntaryExit(beaconBlock *api.Response[*spec.VersionedSignedBeaconBlock], orm *gorm.DB) error {
	voluntaryExits, err := beaconBlock.Data.VoluntaryExits()
	if err != nil {
		return err
	}

	var vExitValidators []uint64
	for _, voluntaryExit := range voluntaryExits {
		vExitValidators = append(vExitValidators, uint64(voluntaryExit.Message.ValidatorIndex))
	}

	var validators []uint64
	if err := orm.Model(&models.Validator{}).Select("validator_index").
		Where("validator_index in ?", vExitValidators).Find(&validators).Error; err != nil {
		return err
	}

	if len(validators) != 0 {
		slot, _ := beaconBlock.Data.Slot()
		if err := orm.Model(&models.Validator{}).Where("validator_index in ?", validators).
			Update("voluntary_exit", uint64(slot)).Error; err != nil {
			return err
		}
		logrus.Infof("find VoluntaryExit len(%d), slot[%d]", len(validators), slot)
	}
	return nil
}

func (s *Scanner) processWithdrawnOnChain(beaconBlock *api.Response[*spec.VersionedSignedBeaconBlock], orm *gorm.DB) error {
	withdrawals, err := beaconBlock.Data.Withdrawals()
	if err != nil {
		return err
	}

	var vWithdraws []uint64
	for _, withdraw := range withdrawals {
		if withdraw.Amount > 29000000000 {
			vWithdraws = append(vWithdraws, uint64(withdraw.ValidatorIndex))
		}
	}

	var validators []uint64
	if err := orm.Model(&models.Validator{}).Select("validator_index").
		Where("validator_index in ?", vWithdraws).Find(&validators).Error; err != nil {
		return err
	}

	if len(validators) != 0 {
		slot, _ := beaconBlock.Data.Slot()
		if err := orm.Model(&models.Validator{}).Where("validator_index in ?", validators).
			Update("withdrawn_on_chain", uint64(slot)).Error; err != nil {
			return err
		}
		logrus.Infof("find WithdrawnOnChain len(%d), slot[%d]", len(validators), slot)
	}
	return nil
}
