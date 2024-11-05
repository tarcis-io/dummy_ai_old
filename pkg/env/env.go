package env

var (
	serverAddress = getEnv("SERVER_ADDRESS", ":3000")
)

func ServerAddress() string {

	return serverAddress
}

func getEnv(key string, defaultValue string) string {

	return defaultValue
}
