package env

import (
	"os"
)

type (
	config struct {
		language      string
		serverAddress string
	}
)

var (
	envConfig *config
)

func init() {
	envConfig = &config{
		language:      env("LANGUAGE", "en"),
		serverAddress: env("SERVER_ADDRESS", ":8080"),
	}
}

func Language() string {
	return envConfig.language
}

func ServerAddress() string {
	return envConfig.serverAddress
}

func env(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}
