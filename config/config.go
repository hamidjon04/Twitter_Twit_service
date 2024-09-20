package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	DB_HOST              string
	DB_PORT              string
	DB_USER              string
	DB_PASSWORD          string
	DB_NAME              string
	TWIT_SERVICE	     string
}

func LoadConfig() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("error loading .env file or not found", err)
	}

	config := Config{}

	

	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToString(coalesce("DB_PORT", "5432"))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "your-password"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "your-db-name"))
	config.TWIT_SERVICE = cast.ToString(coalesce("TWIT-SERVICE", ":8082"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
