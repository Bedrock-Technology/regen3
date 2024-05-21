package log

import (
	"github.com/sirupsen/logrus"
	"sync"
	"testing"
	"time"
)

func init() {
	_ = Init("info", "./test.log", "", "testModule")
}

func TestGoutine(t *testing.T) {
	sy := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		sy.Add(1)
		go func(a int) {
			requestLogger := logrus.WithFields(logrus.Fields{"request_id": a})
			requestLogger.Info("this is info with routine")
			sy.Done()
		}(i)
	}
	sy.Wait()
	logrus.Errorln("end")
}

func TestRotate(t *testing.T) {
	for {
		logrus.Infoln("test rotate")
		time.Sleep(10 * time.Second)
	}
}
