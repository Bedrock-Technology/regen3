package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestSerializeConfig(t *testing.T) {
	config := Config{}
	out, err := yaml.Marshal(&config)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
