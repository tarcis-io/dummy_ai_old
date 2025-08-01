package util

type (
	locale struct {
		App            string `json:"app"`
		AppDescription string `json:"appDescription"`
		AppDevelopedBy string `json:"appDevelopedBy"`
		AppVersion     string `json:"appVersion"`
	}
)

var (
	currentLocale locale
)
