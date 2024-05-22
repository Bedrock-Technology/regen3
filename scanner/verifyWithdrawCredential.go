package scanner

import (
	"fmt"
	"github.com/Bedrock-Technology/regen3/beaconClient"
	"github.com/Bedrock-Technology/regen3/models"
	"github.com/Bedrock-Technology/regen3/proofgen"
	"github.com/attestantio/go-eth2-client/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

func (s *Scanner) InitVerifyWithdrawCredential() error {
	blockNow, err := s.findLatestBlockNumber()
	if err != nil {
		return err
	}
	s.BlockTimer.NewJob(blockNow, s.Config.CheckVerifyWithdrawCredential.IntervalBlock,
		s.Config.CheckVerifyWithdrawCredential.FirstRun,
		&VerifyWithdrawCredentialRun{
			scanner: s,
		})
	logrus.Infoln("Add VerifyWithdrawCredential timer blockNow:", blockNow)
	return nil
}

type VerifyWithdrawCredentialRun struct {
	scanner *Scanner
}

func (s *Scanner) findLatestBlockNumber() (uint64, error) {
	cursor, err := models.GetCursor(s.DBEngine, models.Scanner)
	if err != nil {
		logrus.Errorln("GetCursor:", err)
		return 0, err
	}
	for {
		slotBody, err := s.BeaconClient.SignedBeaconBlock(beaconClient.CTX, &api.SignedBeaconBlockOpts{
			Block: fmt.Sprintf("%d", cursor.Slot),
		})
		if err == nil {
			executionBlockNumber, errEb := slotBody.Data.ExecutionBlockNumber()
			if errEb == nil {
				return executionBlockNumber, nil
			}
		}
		cursor.Slot--
	}
}

const batchSize = int(20)

func (v *VerifyWithdrawCredentialRun) JobRun() {
	logrus.Info("VerifyWithdrawCredentialRun")
	for _, pod := range v.scanner.Pods {
		if pod.IsCredential == 1 {
			//get pod's validators that need to verify
			validators := make([]uint64, 0, batchSize)
			rest := v.scanner.DBEngine.Model(&models.Validator{}).Select("validator_index").Where("pod_address = ?", pod.Address).
				Where("credential_verified = ?", 0).
				Find(&validators).Limit(batchSize)
			if rest.Error != nil {
				logrus.Infof("Get pod's[%s] validators that need to verify error: %v", pod.Address, rest.Error)
				return
			}
			if len(validators) != 0 {
				logrus.Infof("need send[%v] to chain", validators)
				cursor, err := models.GetCursor(v.scanner.DBEngine, models.EigenOracle)
				if err != nil {
					logrus.Errorln("GetCursor:", err)
					return
				}

				statefileName := fmt.Sprintf(beaconStateFormat, v.scanner.Config.Network, cursor.Slot)
				statefilePath := fmt.Sprintf("%s/%s", v.scanner.Config.DataPath, statefileName)
				headerfileName := fmt.Sprintf(beaconHeaderFormat, v.scanner.Config.Network, cursor.Slot)
				headerfilePath := fmt.Sprintf("%s/%s", v.scanner.Config.DataPath, headerfileName)

				err = v.scanner.getBeaconStates(statefilePath, 0)
				if err != nil {
					return
				}
				err = v.scanner.getBeaconHeaders(headerfilePath, 0)
				if err != nil {
					return
				}
				tx, err := v.scanner.getVerifyWithdrawCredentialTx(statefilePath, headerfilePath, common.HexToAddress(pod.Address), validators)
				if err != nil {
					return
				}
				logrus.Infof("VerifyWithdrawCredential tx: %v", tx.Data())
			}
		}
	}
}

func (s *Scanner) getVerifyWithdrawCredentialTx(oracleStateFile, oracleHeaderFile string,
	podAddress common.Address,
	validators []uint64) (*types.Transaction, error) {

	tx, err := proofgen.VerifyWithdrawalCredentialsGen2(s.Submitter, oracleStateFile, oracleHeaderFile, podAddress, validators)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// holesky_beacon_state_1703748
const beaconStateFormat = "%s_state_%d.json"

// holesky_beacon_header_1703748
const beaconHeaderFormat = "%s_header_%d.json"

func (s *Scanner) getBeaconStates(filePath string, slot uint64) error {
	url := fmt.Sprintf("%s/eth/v2/debug/beacon/states/%v", s.Config.BeaconClient, slot)
	return downloadFile(filePath, url)
}

func (s *Scanner) getBeaconHeaders(filePath string, slot uint64) error {
	url := fmt.Sprintf("%s/eth/v1/beacon/headers/%v", s.Config.BeaconClient, slot)
	return downloadFile(filePath, url)
}

func downloadFile(filePath, url string) error {
	exist, _ := pathExists(filePath)
	if exist {
		return nil
	} else {
		logrus.Infoln("downloading ", filePath)
		// Create the file
		out, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer out.Close()

		stateUrl := url
		stateUrlResponse, err := http.Get(stateUrl)
		if err != nil {
			return err
		}
		if stateUrlResponse.StatusCode != http.StatusOK {
			return fmt.Errorf("download error:%v", stateUrlResponse.StatusCode)
		}
		defer stateUrlResponse.Body.Close()

		_, err = io.Copy(out, stateUrlResponse.Body)
		if err != nil {
			return err
		}
	}
	return nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
