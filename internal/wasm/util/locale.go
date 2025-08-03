package util

import (
	"fmt"

	"dummy_ai/internal/wasm/dom"
)

type (
	// locale holds the localized strings for the application.
	locale struct {
		App            string `json:"app"`
		AppDescription string `json:"appDescription"`
		AppDevelopedBy string `json:"appDevelopedBy"`
		AppVersion     string `json:"appVersion"`
	}
)

var (
	// englishLocale represents the localized strings
	// for the English language.
	englishLocale = &locale{
		App:            "DummyAI",
		AppDescription: "Artificial intelligence for dummies",
		AppDevelopedBy: "Developed by t@rcis.io",
		AppVersion:     "Version v0.0.1",
	}

	// currentLocale holds the currently loaded locale.
	currentLocale = fetchCurrentLocale()
)

// App returns the localized name of the application.
func App() string {
	return currentLocale.App
}

// AppDescription returns the localized description of the application.
func AppDescription() string {
	return currentLocale.AppDescription
}

// AppDevelopedBy returns the localized "developed by" credit.
func AppDevelopedBy() string {
	return currentLocale.AppDevelopedBy
}

// AppVersion returns the localized version of the application.
func AppVersion() string {
	return currentLocale.AppVersion
}

// fetchCurrentLocale fetches the current locale for the currently selected language.
func fetchCurrentLocale() *locale {
	if currentLanguage == english {
		return englishLocale
	}
	v, err := dom.Fetch[locale](fmt.Sprintf("/locale/%s.json", currentLanguage.code))
	if err != nil {
		return englishLocale
	}
	return v
}
