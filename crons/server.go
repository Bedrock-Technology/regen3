package crons

import (
	"context"
	"time"

	"github.com/robfig/cron/v3"
)

type ScanCron struct {
	Cron *cron.Cron
}

func New() *ScanCron {
	scanCron := &ScanCron{}
	scanCron.Cron = cron.New(cron.WithLocation(time.UTC), cron.WithSeconds())
	return scanCron
}

func (sc *ScanCron) AddSpec(spec string, job cron.Job) (cron.EntryID, error) {
	wrappedJob := cron.NewChain(cron.SkipIfStillRunning(cron.DiscardLogger)).Then(job)
	return sc.Cron.AddJob(spec, wrappedJob)
}

func (sc *ScanCron) Start() {
	sc.Cron.Start()
}

func (sc *ScanCron) Stop() context.Context {
	return sc.Cron.Stop()
}

func (sc *ScanCron) Remove(id cron.EntryID) {
	sc.Cron.Remove(id)
}
