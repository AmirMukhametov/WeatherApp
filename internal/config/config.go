package config

import "os"

func GetApiKey() {
	return os.Getenv("OPENWEATHER_API_KEY")
}