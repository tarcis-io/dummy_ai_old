package util

const (
	app = iota
	cameraErrorNotSupportedText
	cameraErrorNotSupportedTitle
	cameraLoading
	cameraReloadPage
)

var (
	allMessages = map[string]map[int]string{
		English: {
			app:                          "DummyAI",
			cameraErrorNotSupportedText:  "Please check your device settings",
			cameraErrorNotSupportedTitle: "Camera not supported",
			cameraLoading:                "Loading...",
			cameraReloadPage:             "Reload page",
		},
		Spanish: {
			app:                          "DummyAI",
			cameraErrorNotSupportedText:  "Verifique la configuración de su dispositivo",
			cameraErrorNotSupportedTitle: "Cámara no soportada",
			cameraLoading:                "Cargando...",
			cameraReloadPage:             "Recargar página",
		},
		Portuguese: {
			app:                          "DummyAI",
			cameraErrorNotSupportedText:  "Verifique as configurações do seu dispositivo",
			cameraErrorNotSupportedTitle: "Câmera não suportada",
			cameraLoading:                "Carregando...",
			cameraReloadPage:             "Recarregar página",
		},
	}
	messages = allMessages[language]
)

func App() string {

	return messages[app]
}

func CameraErrorNotSupportedText() string {

	return messages[cameraErrorNotSupportedText]
}

func CameraErrorNotSupportedTitle() string {

	return messages[cameraErrorNotSupportedTitle]
}

func CameraLoading() string {

	return messages[cameraLoading]
}

func CameraReloadPage() string {

	return messages[cameraReloadPage]
}
