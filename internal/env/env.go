// Package env provides a centralized way to manage configuration settings for the application.
package env

import (
	"os"
)

type (
	// config represents the configuration settings for the application.
	config struct {
		languageCode  string
		serverAddress string
	}
)

var (
	// configMap stores the loaded configuration settings, initialized from environment variables or defaults.
	configMap = &config{
		languageCode:  lookupEnv("LANGUAGE_CODE", "en"),
		serverAddress: lookupEnv("SERVER_ADDRESS", ":8080"),
	}
)

// LanguageCode returns the configured language code for the application.
func LanguageCode() string {
	return configMap.languageCode
}

// ServerAddress returns the configured server address for the application.
func ServerAddress() string {
	return configMap.serverAddress
}

// lookupEnv retrieves the value of an environment variable or returns a default value if not found.
func lookupEnv(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return defaultValue
}
