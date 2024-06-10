package blockTimer

import (
	"fmt"
	"testing"
	"time"
)

type Job1 struct {
	Name string
}

func (j *Job1) JobRun() uint64 {
	fmt.Println("job1 run")
	return 444
}

type Job2 struct {
	PrivateField string
	Called       uint64
	EntryId      EntryId
}

func (j *Job2) JobRun() uint64 {
	fmt.Println("job2 run")
	if j.Called == 3 {
		fmt.Println("del job2 run")
		Bt.DelTimer(j.EntryId)
	}
	j.Called++
	return 0
}

var Bt = NewBlockTimer()

type Job3 struct {
	Job string
}

func (j *Job3) JobRun() uint64 {
	fmt.Println("job3 run")
	return 0
}

func TestTimer(t *testing.T) {
	bt := NewBlockTimer()
	blockNow := uint64(100)
	bt.NewJob(blockNow, 100, 34, &Job1{
		Name: "Job1",
	})
	bt.NewJob(blockNow, 0, 44, &Job1{
		Name: "Job1+1",
	})
	bt.NewJob(blockNow, 200, 55, &Job2{
		PrivateField: "Job222",
	})
	bt.NewJob(blockNow, 300, 77, &Job3{
		Job: "Job333",
	})
	bt.printTimer()
	for i := blockNow; i < blockNow+1000; i = i + 99 {
		fmt.Println("block:", i)
		bt.InvokeTimer(i)
		time.Sleep(time.Second * 2)
	}
}

func TestTimer1(t *testing.T) {
	bt := NewBlockTimer()
	blockNow := uint64(100)
	bt.NewJob(blockNow, 100, 34, &Job1{
		Name: "Job1",
	})
	bt.NewJob(blockNow, 200, 55, &Job2{
		PrivateField: "Job222",
	})
	bt.NewJob(blockNow, 300, 77, &Job3{
		Job: "Job333",
	})
	bt.printTimer()
	for i := blockNow; i < blockNow+1000; i++ {
		fmt.Println("block:", i)
		bt.InvokeTimer(i)
	}
}

func TestTimerNoInterval(t *testing.T) {
	bt := NewBlockTimer()
	blockNow := uint64(100)
	bt.NewJob(blockNow, 100, 34, &Job1{
		Name: "Job1",
	})
	bt.NewJob(blockNow, 0, 55, &Job2{
		PrivateField: "Job222",
	})
	bt.printTimer()
	for i := blockNow; i < blockNow+1000; i++ {
		fmt.Println("block:", i)
		bt.InvokeTimer(i)
	}
}

func TestTimerDel(t *testing.T) {
	blockNow := uint64(100)
	Bt.NewJob(blockNow, 100, 34, &Job1{
		Name: "Job1",
	})
	job2 := &Job2{
		PrivateField: "Job222",
	}
	entry := Bt.NewJob(blockNow, 100, 55, job2)
	job2.EntryId = entry
	Bt.printTimer()
	for i := blockNow; i < blockNow+1000; i++ {
		fmt.Println("block:", i)
		Bt.InvokeTimer(i)
	}
}
