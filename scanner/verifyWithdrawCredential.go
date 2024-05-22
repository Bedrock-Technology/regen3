package scanner

import (
	"fmt"
	"github.com/Bedrock-Technology/regen3/beaconClient"
	"github.com/Bedrock-Technology/regen3/models"
	"github.com/attestantio/go-eth2-client/api"
	"github.com/sirupsen/logrus"
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
			validators := make([]models.Validator, 0, batchSize)
			rest := v.scanner.DBEngine.Model(&models.Validator{}).Where("pod_address = ?", pod.Address).
				Find(&validators).Limit(batchSize)
			if rest.Error != nil {
				logrus.Infof("Get pod's[%s] validators that need to verify error: %v", pod.Address, rest.Error)
				return
			}
			//send to chain
			logrus.Infoln("VerifyWithdrawCredential send to chain")
		}
	}
}
