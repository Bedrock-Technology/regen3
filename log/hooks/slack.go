package hooks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

func NewSlackHook(url string) *SlackHook {
	return &SlackHook{
		Url: url,
	}
}

type SlackHook struct {
	Url string
}

func (hook *SlackHook) Fire(entry *logrus.Entry) error {
	data := Payload{}
	data.Text = fmt.Sprintf("[%s] %s", entry.Data["server"], entry.Message)
	dataJson, err := json.Marshal(&data)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return nil
	}

	if entry.Level == logrus.ErrorLevel || entry.Level == logrus.FatalLevel || (entry.Level == logrus.InfoLevel && entry.Data["Report"] == "true") {
		_, _ = http.Post(hook.Url, "application/json",
			strings.NewReader(string(dataJson)))
	}

	return nil
}

func (hook *SlackHook) Levels() []logrus.Level {
	var levels []logrus.Level
	levels = append(levels, logrus.InfoLevel, logrus.ErrorLevel)
	return levels
}

type Payload struct {
	Text string `json:"text"`
}
