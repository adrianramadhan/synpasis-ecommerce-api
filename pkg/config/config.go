package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadEnv(path string) {
	if err := godotenv.Load(path); err != nil {
		log.Println("No .env file found, read from os environment variables.")
	}

	// Load os environment variables
	viper.AutomaticEnv()
}