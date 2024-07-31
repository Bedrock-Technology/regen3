package blockTimer

import (
	"fmt"
	"sort"
)

//Simple SingleThread Block Timer, not suitable for

type BlockTimer struct {
	jobs     []*Job
	jobId    uint64
	blockNow uint64
}

func (bt *BlockTimer) BlockNow() uint64 {
	return bt.blockNow
}

type EntryId uint64

type IJob interface {
	JobRun()
}

type Job struct {
	triggerBlock  uint64
	intervalBlock uint64
	jobId         uint64
	IJob
}

func NewBlockTimer(blockNow uint64) *BlockTimer {
	jobs := make([]*Job, 0)
	return &BlockTimer{
		jobs:     jobs,
		jobId:    0,
		blockNow: blockNow,
	}
}

func (bt *BlockTimer) NewJob(intervalBlock, firstRun uint64, job IJob) (jobId EntryId) {
	bt.jobId++
	j := Job{
		triggerBlock:  bt.blockNow + firstRun,
		intervalBlock: intervalBlock,
		jobId:         bt.jobId,
		IJob:          job,
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
			bt.jobs[i].JobRun()
			if bt.jobs[i].intervalBlock != 0 {
				bt.jobs[i].triggerBlock = bt.jobs[i].intervalBlock + blockNow
				sort.Sort(bt)
			} else {
				bt.jobs = append(bt.jobs[:i], bt.jobs[i+1:]...)
			}
		} else {
			return
		}
	}
	bt.blockNow = blockNow
}

func (bt *BlockTimer) printTimer() {
	for _, v := range bt.jobs {
		fmt.Printf("trigger block %d, interval %d, jobId %d\n", v.triggerBlock, v.intervalBlock, v.jobId)
	}
}
