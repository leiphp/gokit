package config

import (
	"github.com/spf13/viper"
)

func LoadConfig(path string, out interface{}) error {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(out)
}
