package env

import (
	"os"
)

type (
	config struct {
		languageCode  string
		serverAddress string
	}
)

var (
	configMap = &config{
		languageCode:  lookupEnv("LANGUAGE_CODE", "en"),
		serverAddress: lookupEnv("SERVER_ADDRESS", ":8080"),
	}
)

func LanguageCode() string {
	return configMap.languageCode
}

func ServerAddress() string {
	return configMap.serverAddress
}

func lookupEnv(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return defaultValue
}
