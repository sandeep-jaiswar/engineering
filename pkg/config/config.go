package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName   string
	IamPort   string
	Database  string
	LogLevel  string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	return &Config{
		AppName:   getEnv("APP_NAME", ""),
		IamPort:   getEnv("IAM_PORT", ""),
		Database:  getEnv("DATABASE_URL", ""),
		LogLevel:  getEnv("LOG_LEVEL", ""),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
