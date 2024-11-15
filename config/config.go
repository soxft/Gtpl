package config

import (
	"github.com/spf13/viper"
	"gtpl/library/zag"
)

// Init initializes the configuration
func Init() {
	viper.SetConfigType("env")

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		zag.L.Fatalf("Reading .env file: %v", err)
	}

	viper.AutomaticEnv()

	zag.L.Infof("Config loaded %s", viper.AllSettings())

}
