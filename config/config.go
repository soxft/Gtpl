package config

import (
	"github.com/spf13/viper"
	"log"
)

// Init initializes the configuration
func Init() {
	viper.SetConfigType("env")

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Reading .env file: %v", err)
	}

	viper.AutomaticEnv()
}
