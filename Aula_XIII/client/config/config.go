package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Load the configuration from a config file.
func Load() {
	viper.SetConfigFile("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Error reading config file: %s", err))
	}
}
