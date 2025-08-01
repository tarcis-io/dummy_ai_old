package util

import (
	"fmt"

	"dummy_ai/internal/wasm/dom"
)

type (
	locale struct {
		App            string `json:"app"`
		AppDescription string `json:"appDescription"`
		AppDevelopedBy string `json:"appDevelopedBy"`
		AppVersion     string `json:"appVersion"`
	}
)

var (
	currentLocale = fetchCurrentLocale()
)

func App() string {
	return currentLocale.App
}

func AppDescription() string {
	return currentLocale.AppDescription
}

func AppDevelopedBy() string {
	return currentLocale.AppDevelopedBy
}

func AppVersion() string {
	return currentLocale.AppVersion
}

func fetchCurrentLocale() *locale {
	v, _ := dom.Fetch[locale](fmt.Sprintf("/locale/%s.json", currentLanguage.code))
	return v
}
