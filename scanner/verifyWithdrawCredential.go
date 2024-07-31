package scanner

import (
	"context"
	"errors"
	"fmt"
	"github.com/Bedrock-Technology/regen3/beaconClient"
	"github.com/Bedrock-Technology/regen3/contracts/EigenPod"
	"github.com/Bedrock-Technology/regen3/contracts/Restaking"
	"github.com/Bedrock-Technology/regen3/models"
	eigenpodproofs "github.com/Layr-Labs/eigenpod-proofs-generation"
	"github.com/Layr-Labs/eigenpod-proofs-generation/cli/core"
	"github.com/attestantio/go-eth2-client/api"
	apiv1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math/big"
	"strconv"
)

func (s *Scanner) InitVerifyWithdrawCredential() {
	s.BlockTimer.NewJob(s.Config.CheckVerifyWithdrawCredential.IntervalBlock,
		s.Config.CheckVerifyWithdrawCredential.FirstRun,
		&VerifyWithdrawCredentialRun{scanner: s})
	logrus.Infof("Add VerifyWithdrawCredential timer blockNow:%d, firstRun:%d, interval:%d",
		s.BlockTimer.BlockNow(), s.Config.CheckVerifyWithdrawCredential.FirstRun, s.Config.CheckVerifyWithdrawCredential.IntervalBlock)
}

type VerifyWithdrawCredentialRun struct {
	scanner *Scanner
}

func (v *VerifyWithdrawCredentialRun) JobRun() {
	logrus.Info("VerifyWithdrawCredentialRun")
	var beaconBlockHeader *api.Response[*apiv1.BeaconBlockHeader]
	var beaconBlockState *api.Response[*spec.VersionedBeaconState]
	var blockTime uint64
	for _, pod := range v.scanner.Pods {
		if pod.IsCredential == 1 {
			validators := make([]uint64, 0, v.scanner.Config.CheckVerifyWithdrawCredential.BatchSize)
			rest := v.scanner.DBEngine.Model(&models.Validator{}).Select("validator_index").
				Where("pod_address = ?", pod.Address).
				Where("credential_verified = ?", 0).Where("withdrawn_on_chain = ?", 0).
				Where("withdrawn_on_pod = ?", 0).Where("voluntary_exit = ?", 0).
				Limit(v.scanner.Config.CheckVerifyWithdrawCredential.BatchSize).Find(&validators)
			if rest.Error != nil {
				logrus.Errorln("get validators error:", rest.Error)
				return
			}
			if len(validators) != 0 {
				proofs, err := getValidatorProof(pod.Address, v.scanner, validators, &beaconBlockHeader, &beaconBlockState, &blockTime)
				if err != nil {
					logrus.Errorln("getValidatorProof error:", err)
					return
				}
				txData, err := packVerifyWithdrawCredentialTx(proofs, blockTime)
				if err != nil {
					logrus.Errorln("packVerifyWithdrawCredentialTx error:", err)
					continue
				}
				input, err := v.scanner.packVerifyWithdrawalCredentialsInput(txData, big.NewInt(int64(pod.PodIndex)))
				if err != nil {
					logrus.Errorln("packVerifyWithdrawalCredentialsInput error:", err)
					continue
				}
				realTx, err := v.scanner.sendRawTransaction(input, v.scanner.Config.RestakingContract)
				if err != nil {
					if errors.Is(err, errBaseFeeTooHigh) {
						logrus.Warnf("sendRawTransaction pod %v error:%v", pod.Address, errBaseFeeTooHigh)
						return
					}
					logrus.Errorf("sendRawTransaction pod %v error:%v", pod.Address, err)
					panic("sendRawTransaction error")
				}
				logrus.Infoln("waiting sendVerifyWithdrawCredential tx:", realTx.Hash())
				txReceipt, err := bind.WaitMined(context.Background(), v.scanner.EthClient, realTx)
				if err != nil {
					logrus.Errorf("waiting sendVerifyWithdrawCredential pod %v error:%v", pod.Address, err)
					panic("waiting error")
				}
				logrus.WithField("Report", "true").Infof("sendVerifyWithdrawCredential tx:%s", txReceipt.TxHash)
				if err := writeTransaction(v.scanner.DBEngine, txReceipt, TxVerifyWithdrawalCredentials); err != nil {
					logrus.Errorln("writeTransaction err:", err)
					panic("writeTransaction error")
				}
				if err := checkIfEventContained(txReceipt.Logs, validators, pod.Address, v.scanner); err != nil {
					logrus.Errorln("checkIfEventContained err:", err)
					panic("checkIfEventContained error")
				}
				if err := updateValidatorCredential(v.scanner.DBEngine, validators, pod.Address, txReceipt.BlockNumber.Uint64()); err != nil {
					logrus.Errorln("updateValidatorCredential err:", err)
					panic("updateValidatorCredential error")
				}
			}
		}
	}
}

func getValidatorProof(podAddress string, scanner *Scanner, validatorIndices []uint64,
	beaconBlockHeader **api.Response[*apiv1.BeaconBlockHeader], beaconBlockState **api.Response[*spec.VersionedBeaconState],
	blockTime *uint64) (*eigenpodproofs.VerifyValidatorFieldsCallParams, error) {
	if *beaconBlockHeader == nil || *beaconBlockState == nil || *blockTime == 0 {
		latestBlock, err := scanner.EthClient.BlockByNumber(context.Background(), nil)
		if err != nil {
			return nil, fmt.Errorf("getValidatorProof err:%v", err)
		}
		*blockTime = latestBlock.Time()
		eigenPod, err := EigenPod.NewEigenPod(common.HexToAddress(podAddress), scanner.EthClient)
		if err != nil {
			return nil, fmt.Errorf("NewEigenPod err:%v", err)
		}
		expectedBlockRoot, err := eigenPod.GetParentBlockRoot(nil, latestBlock.Time())
		if err != nil {
			return nil, fmt.Errorf("GetParentBlockRoot err:%v", err)
		}
		*beaconBlockHeader, err = scanner.BeaconClient.BeaconBlockHeader(beaconClient.CTX, &api.BeaconBlockHeaderOpts{
			Block: "0x" + common.Bytes2Hex(expectedBlockRoot[:]),
		})
		if err != nil {
			return nil, fmt.Errorf("GetParentBlockRoot err:%v", err)
		}
		*beaconBlockState, err = scanner.BeaconClient.BeaconState(beaconClient.CTX,
			&api.BeaconStateOpts{State: strconv.FormatUint(uint64((*beaconBlockHeader).Data.Header.Message.Slot), 10)})
		if err != nil {
			return nil, fmt.Errorf("BeaconState err:%v", err)
		}
		logrus.Infof("downloading state ok, slot[%d], podAddress[%s]", uint64((*beaconBlockHeader).Data.Header.Message.Slot), podAddress)
	}
	proofExecutor, err := eigenpodproofs.NewEigenPodProofs(scanner.Config.ChainId, 0)
	if err != nil {
		return nil, fmt.Errorf("NewEigenPodProofs err:%v", err)
	}
	logrus.Infof("using state to generate proofs, slot[%d], podAddress[%s]", uint64((*beaconBlockHeader).Data.Header.Message.Slot), podAddress)
	validatorProofs, err := proofExecutor.ProveValidatorContainers((*beaconBlockHeader).Data.Header.Message, (*beaconBlockState).Data, validatorIndices)
	if err != nil {
		return nil, fmt.Errorf("ProveValidatorContainers err:%v", err)
	}
	return validatorProofs, nil
}

func checkIfEventContained(logs []*types.Log, needCheck []uint64, podAddress string, s *Scanner) error {
	egAbi, _ := EigenPod.EigenPodMetaData.GetAbi()
	contract, err := EigenPod.NewEigenPod(common.HexToAddress(podAddress), s.EthClient)
	if err != nil {
		return err
	}
	validatorContains := make([]uint64, 0)
	for _, log := range logs {
		if log.Address == common.HexToAddress(podAddress) {
			e, err := egAbi.EventByID(log.Topics[0])
			if err != nil {
				return err
			}
			if e.Name == "ValidatorRestaked" {
				ew, _ := contract.ParseValidatorRestaked(*log)
				validatorContains = append(validatorContains, ew.ValidatorIndex.Uint64())
			}
		}
	}
	if len(needCheck) != len(validatorContains) {
		return fmt.Errorf("not equal")
	}
	for _, n := range needCheck {
		found := false
		for _, c := range validatorContains {
			if n == c {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("not equal")
		}
	}
	return nil
}

func packVerifyWithdrawCredentialTx(proofs *eigenpodproofs.VerifyValidatorFieldsCallParams, beaconTimestamp uint64) ([]byte, error) {
	eigenPodAbi, _ := EigenPod.EigenPodMetaData.GetAbi()
	stateRootProof := EigenPod.BeaconChainProofsStateRootProof{
		Proof:           proofs.StateRootProof.Proof.ToByteSlice(),
		BeaconStateRoot: proofs.StateRootProof.BeaconStateRoot,
	}
	indices := make([]*big.Int, len(proofs.ValidatorIndices))
	for i, v := range proofs.ValidatorIndices {
		indices[i] = big.NewInt(int64(v))
	}
	validatorFieldsProofs := make([][]byte, len(proofs.ValidatorFieldsProofs))
	for i, v := range proofs.ValidatorFieldsProofs {
		validatorFieldsProofs[i] = v.ToByteSlice()
	}
	curValidatorFields := core.CastValidatorFields(proofs.ValidatorFields)
	return eigenPodAbi.Pack("verifyWithdrawalCredentials", beaconTimestamp, stateRootProof, indices, validatorFieldsProofs, curValidatorFields)
}

func (s *Scanner) packVerifyWithdrawalCredentialsInput(txData []byte, podId *big.Int) ([]byte, error) {
	eigenPodAbi, _ := EigenPod.EigenPodMetaData.GetAbi()
	m := map[string]interface{}{}
	if err := eigenPodAbi.Methods["verifyWithdrawalCredentials"].Inputs.UnpackIntoMap(m, txData[4:]); err != nil {
		return nil, err
	}
	restakingAbi, _ := Restaking.RestakingMetaData.GetAbi()
	return restakingAbi.Pack("verifyWithdrawalCredentials", podId, m["beaconTimestamp"], m["stateRootProof"], m["validatorIndices"], m["validatorFieldsProofs"], m["validatorFields"])
}

func updateValidatorCredential(db *gorm.DB, validators []uint64, podAddress string, blockNumber uint64) error {
	return db.Model(&models.Validator{}).
		Where("validator_index in ?", validators).Where("pod_address = ?", podAddress).
		Update("credential_verified", blockNumber).Error
}
