package util

import (
	"fmt"

	"dummy_ai/internal/wasm/dom"
)

type (
	// Locale holds the localized strings for the application.
	Locale struct {
		App            string `json:"app"`
		AppDescription string `json:"appDescription"`
		AppDevelopedBy string `json:"appDevelopedBy"`
		AppVersion     string `json:"appVersion"`
	}
)

var (
	// englishLocale represents the localized strings
	// for the English language.
	englishLocale = &Locale{
		App:            "DummyAI",
		AppDescription: "Artificial intelligence for dummies",
		AppDevelopedBy: "Developed by t@rcis.io",
		AppVersion:     "Version v0.0.1",
	}

	// currentLocale holds the currently loaded locale.
	currentLocale *Locale
)

// EnglishLocale returns the localized strings for the English language.
func EnglishLocale() *Locale {
	return englishLocale
}

// CurrentLocale returns the currently loaded locale.
func CurrentLocale() *Locale {
	return currentLocale
}

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

// FetchCurrentLocale fetches the current locale for the currently selected language.
func FetchCurrentLocale() *Locale {
	if currentLanguage == english {
		return englishLocale
	}
	v, err := dom.Fetch[Locale](fmt.Sprintf("/locale/%s.json", currentLanguage.code))
	if err != nil {
		return englishLocale
	}
	return v
}
