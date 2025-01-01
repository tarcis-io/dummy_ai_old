package util

const (
	app = iota
	cameraLoading
	cameraReloadPage
)

var (
	allMessages = map[string]map[int]string{
		English: {
			app:              "DummyAI",
			cameraLoading:    "Loading...",
			cameraReloadPage: "Reload page",
		},
		Spanish: {
			app:              "DummyAI",
			cameraLoading:    "Cargando...",
			cameraReloadPage: "Recargar página",
		},
		Portuguese: {
			app:              "DummyAI",
			cameraLoading:    "Carregando...",
			cameraReloadPage: "Recarregar página",
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

func CameraReloadPage() string {

	return messages[cameraReloadPage]
}
