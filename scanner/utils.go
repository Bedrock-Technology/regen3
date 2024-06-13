package scanner

import (
	"errors"
	"fmt"
	"github.com/Bedrock-Technology/regen3/beaconClient"
	"github.com/Bedrock-Technology/regen3/models"
	"github.com/attestantio/go-eth2-client/api"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

const TxVerifyWithdrawalCredentials = "verifyWithdrawalCredentials"
const TxVerifyWithdrawalProof = "verifyWithdrawalProof"
const TxQueueWithdrawals = "queueWithdrawals"
const TxCompleteQueueWithdrawals = "completeQueueWithdrawals"

const beaconStateFormat = "%s_state_%d.json"
const beaconHeaderFormat = "%s_header_%d.json"

func addGasBuffer(gasLimit uint64) uint64 {
	return 6 * gasLimit / 5 // add 20% buffer to gas limit
}

func downloadFile(filePath, url string) error {
	exist, _ := pathExists(filePath)
	if exist {
		return nil
	} else {
		logrus.Infoln("downloading ", url)
		// Create the file
		out, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer out.Close()
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			os.Remove(filePath)
			logrus.Infof("rm: %s\n", filePath)
			return fmt.Errorf("download error:%v", resp.StatusCode)
		}
		defer resp.Body.Close()
		_, err = io.Copy(out, resp.Body)
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

func (s *Scanner) getBeaconStates(filePath string, slot uint64) error {
	url := fmt.Sprintf("%s/eth/v2/debug/beacon/states/%d", s.Config.BeaconClient, slot)
	return downloadFile(filePath, url)
}

func (s *Scanner) getBeaconHeaders(filePath string, slot uint64) error {
	url := fmt.Sprintf("%s/eth/v1/beacon/headers/%d", s.Config.BeaconClient, slot)
	return downloadFile(filePath, url)
}

func (s *Scanner) getBeaconBlocks(filePath string, slot uint64) error {
	url := fmt.Sprintf("%s/eth/v2/beacon/blocks/%d", s.Config.BeaconClient, slot)
	return downloadFile(filePath, url)
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
		if err != nil {
			var apiErr *api.Error
			if errors.As(err, &apiErr) {
				switch apiErr.StatusCode {
				case 404:
					logrus.Warnf("empty slot[%d]", cursor.Slot)
					cursor.Slot--
					continue
				default:
					logrus.Errorf("error[%d]:%v", apiErr.StatusCode, err)
					return 0, err
				}
			}
		}
		executionBlockNumber, err := slotBody.Data.ExecutionBlockNumber()
		if err != nil {
			return 0, err
		}
		return executionBlockNumber, nil
	}
}

func (s *Scanner) findLatestBlockNumberBySlot(slot uint64) (uint64, error) {
	for {
		slotBody, err := s.BeaconClient.SignedBeaconBlock(beaconClient.CTX, &api.SignedBeaconBlockOpts{
			Block: fmt.Sprintf("%d", slot),
		})
		if err != nil {
			var apiErr *api.Error
			if errors.As(err, &apiErr) {
				switch apiErr.StatusCode {
				case 404:
					logrus.Warnf("empty slot[%d]", slot)
					slot--
					continue
				default:
					logrus.Errorf("error[%d]:%v", apiErr.StatusCode, err)
					return 0, err
				}
			}
		}
		executionBlockNumber, err := slotBody.Data.ExecutionBlockNumber()
		if err != nil {
			return 0, err
		}
		return executionBlockNumber, nil
	}
}

var errBaseFeeTooHigh = errors.New("base fee too high")
