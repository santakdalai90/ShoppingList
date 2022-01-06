package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// LoadConfig loads the configuration from config/config.toml
func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error while reading config: %s", err.Error()))
	}
}
