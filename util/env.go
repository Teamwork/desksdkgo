package util

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env file if it exists
func LoadEnv() {
	// Try to load .env file, but don't fail if it doesn't exist
	_ = godotenv.Load()
}

// GetEnv returns the value of the environment variable or the default value if not set
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
