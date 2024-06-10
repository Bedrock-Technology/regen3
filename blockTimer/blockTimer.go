package blockTimer

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"sort"
)

//Simple SingleThread Block Timer, not suitable for

type BlockTimer struct {
	jobs  []*Job
	jobId uint64
}

type EntryId uint64

type IJob interface {
	JobRun() uint64
}

type Job struct {
	triggerBlock  uint64
	intervalBlock uint64
	jobId         uint64
	delete        uint8
	IJob
}

func NewBlockTimer() *BlockTimer {
	jobs := make([]*Job, 0)
	return &BlockTimer{
		jobs:  jobs,
		jobId: 0,
	}
}

func (bt *BlockTimer) NewJob(blockNow, intervalBlock, firstRun uint64, job IJob) (jobId EntryId) {
	bt.jobId++
	j := Job{
		triggerBlock:  blockNow + firstRun,
		intervalBlock: intervalBlock,
		jobId:         bt.jobId,
		IJob:          job,
		delete:        0,
	}
	bt.jobs = append(bt.jobs, &j)
	sort.Sort(bt)
	return EntryId(bt.jobId)
}

func (bt *BlockTimer) Len() int {
	return len(bt.jobs)
}

func (bt *BlockTimer) Less(i, j int) bool {
	return bt.jobs[i].triggerBlock < bt.jobs[j].triggerBlock
}

func (bt *BlockTimer) Swap(i, j int) {
	bt.jobs[i], bt.jobs[j] = bt.jobs[j], bt.jobs[i]
}

func (bt *BlockTimer) InvokeTimer(blockNow uint64) {
	for i := 0; i < len(bt.jobs); i++ {
		if bt.jobs[i].triggerBlock <= blockNow {
			if bt.jobs[i].delete == 1 {
				bt.jobs = append(bt.jobs[:i], bt.jobs[i+1:]...)
				continue
			}
			logrus.Info("Call JobRun block:", blockNow)
			delta := bt.jobs[i].JobRun()
			logrus.Info("Call JobRun delta:", delta)
			if bt.jobs[i].intervalBlock != 0 {
				if delta != 0 {
					bt.jobs[i].triggerBlock = bt.jobs[i].intervalBlock + delta
				} else {
					bt.jobs[i].triggerBlock = bt.jobs[i].intervalBlock + blockNow
				}
				sort.Sort(bt)
			} else {
				bt.jobs = append(bt.jobs[:i], bt.jobs[i+1:]...)
			}
		} else {
			return
		}
	}
}

func (bt *BlockTimer) DelTimer(entryId EntryId) {
	for _, v := range bt.jobs {
		if EntryId(v.jobId) == entryId {
			v.delete = 1
		}
	}
}

func (bt *BlockTimer) printTimer() {
	for _, v := range bt.jobs {
		fmt.Printf("trigger block %d, interval %d, jobId %d\n", v.triggerBlock, v.intervalBlock, v.jobId)
	}
}
