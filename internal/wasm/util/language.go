package util

import (
	"strings"

	"dummy_ai/internal/env"
	"dummy_ai/internal/wasm/dom"
)

type (
	// Language represents a supported language.
	Language struct {
		code string
		name string
	}
)

// Code returns the language code.
func (l *Language) Code() string {
	return l.code
}

// Name returns the language name.
func (l *Language) Name() string {
	return l.name
}

var (
	// englishLanguage represents the English language.
	englishLanguage = &Language{
		code: "en",
		name: "English",
	}

	// spanishLanguage represents the Spanish language.
	spanishLanguage = &Language{
		code: "es",
		name: "Español",
	}

	// portugueseLanguage represents the Portuguese language.
	portugueseLanguage = &Language{
		code: "pt",
		name: "Português",
	}

	// fallbackLanguage is the default language used
	// if no other language can be determined.
	fallbackLanguage = englishLanguage

	// supportedLanguages is a slice of all supported languages.
	supportedLanguages = []*Language{
		englishLanguage,
		spanishLanguage,
		portugueseLanguage,
	}

	// supportedLanguagesMap provides a quick lookup for supported languages.
	supportedLanguagesMap = map[string]*Language{
		englishLanguage.code:    englishLanguage,
		spanishLanguage.code:    spanishLanguage,
		portugueseLanguage.code: portugueseLanguage,
	}

	// currentLanguage holds the currently selected language.
	currentLanguage *Language
)

// EnglishLanguage returns the English language.
func EnglishLanguage() *Language {
	return englishLanguage
}

// SpanishLanguage returns the Spanish language.
func SpanishLanguage() *Language {
	return spanishLanguage
}

// PortugueseLanguage returns the Portuguese language.
func PortugueseLanguage() *Language {
	return portugueseLanguage
}

// FallbackLanguage returns the fallback language.
func FallbackLanguage() *Language {
	return fallbackLanguage
}

// SupportedLanguages returns a slice of all supported languages.
func SupportedLanguages() []*Language {
	return supportedLanguages
}

// SupportedLanguagesMap returns a map of all supported languages.
func SupportedLanguagesMap() map[string]*Language {
	return supportedLanguagesMap
}

// CurrentLanguage returns the currently selected language.
func CurrentLanguage() *Language {
	return currentLanguage
}

// SetLanguage persists the chosen language to the browser's local storage.
func SetLanguage(language *Language) {
	dom.GetLocalStorage().SetItem("language", language.code)
}

// LookupLanguage determines the user's preferred language by checking:
// 1. The browser's local storage.
// 2. The browser's navigator settings.
// 3. The environment variables.
// It returns the fallback language if no supported language is found.
func LookupLanguage() *Language {
	v, ok := lookupLocalStorageLanguage()
	if ok {
		return v
	}
	v, ok = lookupNavigatorLanguage()
	if ok {
		return v
	}
	v, ok = lookupEnvLanguage()
	if ok {
		return v
	}
	return fallbackLanguage
}

// lookupLocalStorageLanguage looks up the language from the local storage.
func lookupLocalStorageLanguage() (*Language, bool) {
	v, ok := dom.GetLocalStorage().GetItem("language")
	if ok {
		return lookupLanguage(v)
	}
	return nil, false
}

// lookupNavigatorLanguage looks up the language from the browser's navigator.
func lookupNavigatorLanguage() (*Language, bool) {
	v, ok := dom.GetNavigator().Language()
	if ok {
		return lookupLanguage(strings.SplitN(v, "-", 2)[0])
	}
	return nil, false
}

// lookupEnvLanguage looks up the language from the environment variables.
func lookupEnvLanguage() (*Language, bool) {
	return lookupLanguage(env.LanguageCode())
}

// lookupLanguage looks up a language by its code.
func lookupLanguage(languageCode string) (*Language, bool) {
	v, ok := supportedLanguagesMap[languageCode]
	return v, ok
}
