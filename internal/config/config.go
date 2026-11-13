package config

import (
	"os"
    "log"
	"github.com/joho/godotenv"
)
func GetApiKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("OPENWEATHER_API_KEY")

    if apiKey == "" {
        log.Fatal("OPENWEATHER_API_KEY не найден")
    }

    return apiKey
}