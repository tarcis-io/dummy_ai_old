package util

import (
	"strings"

	"dummy_ai/internal/env"
	"dummy_ai/internal/wasm/dom"
)

type (
	Language struct {
		code string
		name string
	}
)

func (l *Language) Code() string {
	return l.code
}

func (l *Language) Name() string {
	return l.name
}

var (
	english = &Language{
		code: "en",
		name: "English",
	}
	spanish = &Language{
		code: "es",
		name: "Español",
	}
	portuguese = &Language{
		code: "pt",
		name: "Português",
	}

	supportedLanguages = []*Language{
		english,
		spanish,
		portuguese,
	}

	currentLanguage = LookupLanguage()

	supportedLanguagesMap = map[string]*Language{
		english.code:    english,
		spanish.code:    spanish,
		portuguese.code: portuguese,
	}

	fallbackLanguage = english
)

func English() *Language {
	return english
}

func Spanish() *Language {
	return spanish
}

func Portuguese() *Language {
	return portuguese
}

func SupportedLanguages() []*Language {
	return supportedLanguages
}

func CurrentLanguage() *Language {
	return currentLanguage
}

func SetLanguage(language *Language) {
	currentLanguage = language
	dom.GetLocalStorage().SetItem("language", language.code)
}

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

func lookupLanguage(languageCode string) (*Language, bool) {
	v, ok := supportedLanguagesMap[languageCode]
	return v, ok
}

func lookupLocalStorageLanguage() (*Language, bool) {
	v, ok := dom.GetLocalStorage().GetItem("language")
	if ok {
		return lookupLanguage(v)
	}
	return nil, false
}

func lookupNavigatorLanguage() (*Language, bool) {
	v, ok := dom.GetNavigator().Language()
	if ok {
		return lookupLanguage(strings.SplitN(v, "-", 2)[0])
	}
	return nil, false
}

func lookupEnvLanguage() (*Language, bool) {
	return lookupLanguage(env.LanguageCode())
}
