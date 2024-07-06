package config

import "github.com/spf13/viper"

func PostgresHost() string {
	return viper.GetString("POSTGRES_HOST")
}

func PostgresPort() string {
	return viper.GetString("POSTGRES_PORT")
}

func PostgresUser() string {
	return viper.GetString("POSTGRES_USER")
}

func PostgresPassword() string {
	return viper.GetString("POSTGRES_PASSWORD")
}

func PostgresDb() string {
	return viper.GetString("POSTGRES_DB")
}

func PostgresTimezone() string {
	return viper.GetString("POSTGRES_TIMEZONE")
}

func PostgresSslMode() string {
	return viper.GetString("POSTGRES_SSL_MODE")
}