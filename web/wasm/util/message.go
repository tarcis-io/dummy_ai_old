package util

const (
	app = iota
	cameraLoading
)

var (
	allMessages = map[string]map[int]string{
		English: {
			app:           "DummyAI",
			cameraLoading: "Loading...",
		},
		Spanish: {
			app:           "DummyAI",
			cameraLoading: "Cargando...",
		},
		Portuguese: {
			app:           "DummyAI",
			cameraLoading: "Carregando...",
		},
	}
	messages = allMessages[language]
)

func App() string {

	return messages[app]
}

func CameraLoading() string {

	return messages[cameraLoading]
}
