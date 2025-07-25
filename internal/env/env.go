package env

import (
	"os"
)

type (
	env struct {
		language      string
		serverAddress string
	}
)

var (
	appEnv *env
)

func init() {
	appEnv = &env{
		language:      lookupEnv("LANGUAGE", "en"),
		serverAddress: lookupEnv("SERVER_ADDRESS", ":8080"),
	}
}

func Language() string {
	return appEnv.language
}

func ServerAddress() string {
	return appEnv.serverAddress
}

func lookupEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}
