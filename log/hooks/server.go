package hooks

import (
	"github.com/sirupsen/logrus"
)

func NewServerHook(name string) *ServerHook {
	return &ServerHook{
		Name: name,
	}
}

type ServerHook struct {
	Name string
}

func (hook *ServerHook) Fire(entry *logrus.Entry) error {
	entry.Data["server"] = hook.Name
	return nil
}

func (hook *ServerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
