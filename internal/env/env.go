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
	appConfig = &config{
		languageCode:  lookupEnv("LANGUAGE_CODE", "en"),
		serverAddress: lookupEnv("SERVER_ADDRESS", ":8080"),
	}
)

func LanguageCode() string {
	return appConfig.languageCode
}

func ServerAddress() string {
	return appConfig.serverAddress
}

func lookupEnv(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return v
}
