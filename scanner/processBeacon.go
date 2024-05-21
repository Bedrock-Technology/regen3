package scanner

import (
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
	err := s.processDeposit(beaconBlock, orm)
	if err != nil {
		return err
	}
	//process VoluntaryExit
	//process WithdrawnOnChain
	return nil
}

func (s *Scanner) processDeposit(beaconBlock *api.Response[*spec.VersionedSignedBeaconBlock], orm *gorm.DB) error {
	deposits, err := beaconBlock.Data.Deposits()
	if err != nil {
		return err
	}
	modelValidators := make([]models.Validator, 0)

	for _, deposit := range deposits {
		withdrawal := common.BytesToAddress(deposit.Data.WithdrawalCredentials[12:]).String()
		pubKey := fmt.Sprintf("%#x", deposit.Data.PublicKey)
		if _, exist := s.Pods[withdrawal]; exist {
			modelValidator := models.Validator{
				PubKey:             pubKey,
				ValidatorIndex:     0,
				PodAddress:         withdrawal,
				CredentialVerified: 0,
				WithdrawnOnChain:   0,
				WithdrawnOnPod:     0,
				VoluntaryExit:      0,
			}
			modelValidators = append(modelValidators, modelValidator)
		}
	}

	blsPubKey := make([]phase0.BLSPubKey, 0)
	for _, modelValidator := range modelValidators {
		pubKeyS := fmt.Sprintf(`"%s"`, modelValidator.PubKey)
		pubKeyBls := phase0.BLSPubKey{}
		_ = pubKeyBls.UnmarshalJSON([]byte(pubKeyS))
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
				if v.Validator.PublicKey.String() == k.PubKey {
					modelValidators[i].ValidatorIndex = uint64(v.Index)
				}
			}
		}
		logrus.Infof("find Deposit len(%d)", len(blsPubKey))
		return orm.Create(&modelValidators).Error
	}
	return nil
}
