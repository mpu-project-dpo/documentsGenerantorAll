package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

func Parse[C any]() *C {
	cfg := new(C)

	f, err := os.ReadFile("config/config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(f, &cfg)

	return cfg
}
