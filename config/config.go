package config

import (
	"github.com/spf13/viper"
	"log"
)

func LoadEnv() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using system env")
	}
}

func Env(key string, fallback string) string {
	if value := viper.GetString(key); value != "" {
		return value
	}
	return fallback
}
