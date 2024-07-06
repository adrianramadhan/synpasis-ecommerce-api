package config

import "github.com/spf13/viper"

func AppPort() string {
	return viper.GetString("APP_PORT")
}

func AppEnv() string {
	return viper.GetString("APP_ENV")
}