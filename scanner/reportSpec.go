package scanner

import (
	"context"
	"github.com/Bedrock-Technology/regen3/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"time"
)

type ReportSpec struct {
	Scanner *Scanner
}

type Transactions struct {
	Total uint64
	Incr  uint64
}

func (r *ReportSpec) Run() {
	logrus.Info("ReportSpec Run")
	balance, err := r.Scanner.EthClient.BalanceAt(context.Background(),
		common.HexToAddress(r.Scanner.Config.KeyAgent.Address), nil)
	if err != nil {
		logrus.Error("balanceAt error:", err)
		return
	}
	balanceDecimal := decimal.NewFromBigInt(balance, -18)
	ago := time.Now().UTC().Add(-12 * time.Hour)
	trans := Transactions{}
	rst := r.Scanner.
		DBEngine.Raw("SELECT COUNT(*) AS total, COUNT(CASE WHEN created_at > ? THEN 1 END) AS incr FROM transactions", ago).Scan(&trans)
	if rst.Error != nil {
		logrus.Errorln("get validatorNums error:", rst.Error)
		return
	}
	cursor, err := models.GetCursor(r.Scanner.DBEngine, models.Scanner)
	if err != nil {
		logrus.Errorln("GetCursor:", err)
		return
	}
	end, err := r.Scanner.BeaconClient.GetLatestSlotNumber()
	if err != nil {
		logrus.Errorln("Get Beacon Latest Slot Number err:", err)
		return
	}
	diff := end - cursor.Slot
	logrus.WithField("Report", "true").Infof("balance:%s, total:%d, incr:%d, diff:%d",
		balanceDecimal.Truncate(9), trans.Total, trans.Incr, diff)
}
