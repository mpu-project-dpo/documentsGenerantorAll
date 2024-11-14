package config

import (
	"github.com/spf13/viper"
)

func Parse[C any]() *C {
	cfg := new(C)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(cfg)

	return cfg
}
