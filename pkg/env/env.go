package env

var (
	serverAddress = get("SERVER_ADDRESS", ":3000")
)

func ServerAddress() string {

	return serverAddress
}

func get(key string, defaultValue string) string {

	return defaultValue
}
