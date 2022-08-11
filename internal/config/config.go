package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig(configname string) (*Config, error) {

	var c Config

	viper.AutomaticEnv()
	viper.SetConfigName(configname)
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("rapido")
	viper.AddConfigPath("config")
	viper.AddConfigPath("../config/")
	viper.AddConfigPath("../../config/")
	viper.AddConfigPath("../../../config/")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error initializing config")
		panic(err)
	}

	err := viper.Unmarshal(&c)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal config file reason : [%s]", err)
	}
	return &c, nil
}
