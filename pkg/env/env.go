package env

import (
	"os"
)

var (
	serverAddress = getEnv("SERVER_ADDRESS", ":8080")
)

func ServerAddress() string {

	return serverAddress
}

func getEnv(key string, defaultValue string) string {

	if value, isPresent := os.LookupEnv(key); isPresent {

		return value
	}

	return defaultValue
}
