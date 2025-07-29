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
	language *Language

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

	languages = []*Language{
		english,
		spanish,
		portuguese,
	}

	languageMap = map[string]*Language{
		english.code:    english,
		spanish.code:    spanish,
		portuguese.code: portuguese,
	}

	fallbackLanguage = english
)

func init() {
	language = LookupLanguage()
}

func English() *Language {
	return english
}

func Spanish() *Language {
	return spanish
}

func Portuguese() *Language {
	return portuguese
}

func Languages() []*Language {
	return languages
}

func GetLanguage() *Language {
	return language
}

func SetLanguage(language *Language) {
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

func lookupLanguage(code string) (*Language, bool) {
	v, ok := languageMap[code]
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
