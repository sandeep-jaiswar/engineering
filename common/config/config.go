package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config structure to hold all configuration fields
type Config struct {
	AppName   string
	IamPort   string
	Database  string
	LogLevel  string
}

// LoadConfig loads configuration from environment variables or .env files
func LoadConfig() (*Config, error) {
	// Load .env file (if present)
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	// Return the loaded configuration
	return &Config{
		AppName:   getEnv("APP_NAME", ""),
		IamPort:   getEnv("IAM_PORT", ""),
		Database:  getEnv("DATABASE_URL", ""),
		LogLevel:  getEnv("LOG_LEVEL", ""),
	}, nil
}

// Helper to get environment variables with a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
