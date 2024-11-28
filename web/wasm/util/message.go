package util

const (
	App = iota
	CameraLoading
)

var (
	allMessages = map[string]map[int]string{
		English: {
			App:           "DummyAI",
			CameraLoading: "Loading...",
		},
		Spanish: {
			App:           "DummyAI",
			CameraLoading: "Cargando...",
		},
		Portuguese: {
			App:           "DummyAI",
			CameraLoading: "Carregando...",
		},
	}
	messages = allMessages[language]
)

func AllMessages() map[string]map[int]string {

	return allMessages
}

func Messages() map[int]string {

	return messages
}

func Message(key int) string {

	return messages[key]
}
