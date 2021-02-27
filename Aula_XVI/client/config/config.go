package config

import "github.com/spf13/viper"

// Load is a func that returns the configuration fetched from the configuration
// file.
func Load() (*viper.Viper, error) {
	conf := viper.GetViper()
	conf.AddConfigPath(".")
	conf.SetConfigFile("configuration")
	conf.SetConfigType("yaml")
	err := conf.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return conf, nil
}
