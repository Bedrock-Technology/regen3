package scanner

import (
	"context"
	"time"

	"github.com/Bedrock-Technology/regen3/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type ReportSpec struct {
	Scanner        *Scanner
	LastReportTime time.Time
}

type Transactions struct {
	TotalFee string
	IncrFee  string
	Total    uint64
	Incr     uint64
}

func (r *ReportSpec) Run() {
	logrus.Info("ReportSpec Run")
	now := time.Now().UTC()
	balance, err := r.Scanner.EthClient.BalanceAt(context.Background(),
		common.HexToAddress(r.Scanner.Config.KeyAgent.Address), nil)
	if err != nil {
		logrus.Error("balanceAt error:", err)
		return
	}
	balanceDecimal := decimal.NewFromBigInt(balance, -18)
	trans := Transactions{}
	logrus.Infof("LastReportTime:%s", r.LastReportTime.String())
	rst := r.Scanner.
		DBEngine.Raw(`SELECT
    COUNT(*) AS total,
    COUNT(CASE WHEN created_at > ? THEN 1 END) AS incr,
    SUM(fee) AS total_fee,
    SUM(CASE WHEN created_at > ? THEN fee ELSE 0 END) AS incr_fee
	FROM
    transactions`, r.LastReportTime, r.LastReportTime).Scan(&trans)
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
	r.LastReportTime = now
	totalFee, _ := decimal.NewFromString(trans.TotalFee)
	incrFee, _ := decimal.NewFromString(trans.IncrFee)
	logrus.WithField("Report", "true").Infof("balance:%s, total:%d[%s], incr:%d[%s], diff:%d",
		balanceDecimal.Truncate(9), trans.Total, totalFee.Mul(decimal.New(1, -18)).Truncate(9),
		trans.Incr, incrFee.Mul(decimal.New(1, -18)).Truncate(9), diff)
}
