package blockTimer

import (
	"fmt"
	"testing"
	"time"
)

type Job1 struct {
	Name string
}

func (j *Job1) JobRun() {
	fmt.Println("job run:", j.Name)
	return
}

type Job2 struct {
	PrivateField string
	Called       uint64
	EntryId      EntryId
}

func (j *Job2) JobRun() {
	fmt.Println("job2 run")
	if j.Called == 3 {
		fmt.Println("del job2 run")
	}
	j.Called++
	return
}

type Job3 struct {
	Job string
}

func (j *Job3) JobRun() {
	fmt.Println("job3 run")
	return
}

func TestTimer(t *testing.T) {
	blockNow := uint64(100)
	bt := NewBlockTimer(100)
	bt.NewJob(100, 34, &Job1{
		Name: "Job1",
	})
	bt.NewJob(0, 44, &Job1{
		Name: "Job1+1",
	})
	bt.NewJob(200, 55, &Job2{
		PrivateField: "Job222",
	})
	bt.NewJob(300, 77, &Job3{
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
	blockNow := uint64(100)
	bt := NewBlockTimer(blockNow)
	bt.NewJob(100, 34, &Job1{
		Name: "Job1",
	})
	bt.NewJob(200, 55, &Job2{
		PrivateField: "Job222",
	})
	bt.NewJob(300, 77, &Job3{
		Job: "Job333",
	})
	bt.printTimer()
	for i := blockNow; i < blockNow+1000; i++ {
		fmt.Println("block:", i)
		bt.InvokeTimer(i)
	}
}

func TestTimerNoInterval(t *testing.T) {
	blockNow := uint64(100)
	bt := NewBlockTimer(blockNow)
	bt.NewJob(100, 34, &Job1{
		Name: "Job1",
	})
	bt.NewJob(0, 55, &Job2{
		PrivateField: "Job222",
	})
	bt.printTimer()
	for i := blockNow; i < blockNow+1000; i++ {
		fmt.Println("block:", i)
		bt.InvokeTimer(i)
	}
}
