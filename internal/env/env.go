package env

import (
	"os"
)

var (
	language      = env("LANGUAGE", "en")
	serverAddress = env("SERVER_ADDRESS", ":8080")
)

func Language() string {
	return language
}

func ServerAddress() string {
	return serverAddress
}

func env(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}
