// Package env provides a centralized way to manage configuration settings
// for the application.
package env

import (
	"os"
)

type (
	// config holds application configuration settings.
	config struct {
		languageCode  string
		serverAddress string
	}
)

var (
	// configMap stores configuration settings loaded from env or defaults.
	configMap = &config{
		languageCode:  lookupEnv("LANGUAGE_CODE", "en"),
		serverAddress: lookupEnv("SERVER_ADDRESS", ":8080"),
	}
)

// LanguageCode returns the configured LANGUAGE_CODE for the application.
func LanguageCode() string {
	return configMap.languageCode
}

// ServerAddress returns the configured SERVER_ADDRESS for the application.
func ServerAddress() string {
	return configMap.serverAddress
}

// lookupEnv retrieves an environment variable or returns a default.
func lookupEnv(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return defaultValue
}
