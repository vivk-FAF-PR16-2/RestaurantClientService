package config

import (
	"github.com/spf13/viper"
	"log"
)

func Load() {
	viper.AddConfigPath("./config")

	viper.SetConfigName("cfg")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error: can not read config file: %v", err)
	}
}
