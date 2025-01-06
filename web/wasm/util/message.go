package util

const (
	app = iota
	cameraErrorNotAllowedText
	cameraErrorNotAllowedTitle
	cameraErrorNotFoundText
	cameraErrorNotFoundTitle
	cameraErrorNotSupportedText
	cameraErrorNotSupportedTitle
	cameraLoading
	cameraReloadPage
)

var (
	allMessages = map[string]map[int]string{
		English: {
			app:                          "DummyAI",
			cameraErrorNotAllowedText:    "Please enable camera access in your device settings",
			cameraErrorNotAllowedTitle:   "Camera not allowed",
			cameraErrorNotFoundText:      "Please check if your camera is connected and enabled",
			cameraErrorNotFoundTitle:     "Camera not found",
			cameraErrorNotSupportedText:  "Please check your device settings",
			cameraErrorNotSupportedTitle: "Camera not supported",
			cameraLoading:                "Loading...",
			cameraReloadPage:             "Reload page",
		},
		Spanish: {
			app:                          "DummyAI",
			cameraErrorNotAllowedText:    "",
			cameraErrorNotAllowedTitle:   "",
			cameraErrorNotFoundText:      "",
			cameraErrorNotFoundTitle:     "",
			cameraErrorNotSupportedText:  "Verifique la configuración de su dispositivo",
			cameraErrorNotSupportedTitle: "Cámara no soportada",
			cameraLoading:                "Cargando...",
			cameraReloadPage:             "Recargar página",
		},
		Portuguese: {
			app:                          "DummyAI",
			cameraErrorNotAllowedText:    "",
			cameraErrorNotAllowedTitle:   "",
			cameraErrorNotFoundText:      "",
			cameraErrorNotFoundTitle:     "",
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

func CameraErrorNotAllowedText() string {

	return messages[cameraErrorNotAllowedText]
}

func CameraErrorNotAllowedTitle() string {

	return messages[cameraErrorNotAllowedTitle]
}

func CameraErrorNotFoundText() string {

	return messages[cameraErrorNotFoundText]
}

func CameraErrorNotFoundTitle() string {

	return messages[cameraErrorNotFoundTitle]
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
