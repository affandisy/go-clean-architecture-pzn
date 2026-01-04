package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// NewViper is a function to load config from config.json
func NewViper() *viper.Viper {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("json")
	// config.AddConfigPath(".")
	config.AddConfigPath("./../")
	config.AddConfigPath("./")
	err := config.ReadInConfig()

	if err != nil {
		// return nil, err
		panic(fmt.Errorf("Fatal error config file: %w", err))
	}

	// return config, nil
	return config
}
