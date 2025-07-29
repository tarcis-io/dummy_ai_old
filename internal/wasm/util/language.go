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

func LookupLanguage(code string) (*Language, bool) {
	v, ok := languageMap[code]
	return v, ok
}

func SetLanguage(language *Language) {
	dom.GetLocalStorage().SetItem("language", language.code)
}

func lookupLocalStorageLanguage() (*Language, bool) {
	v, ok := dom.GetLocalStorage().GetItem("language")
	if ok {
		return LookupLanguage(v)
	}
	return nil, false
}

func lookupNavigatorLanguage() (*Language, bool) {
	v, ok := dom.GetNavigator().Language()
	if ok {
		return LookupLanguage(strings.SplitN(v, "-", 2)[0])
	}
	return nil, false
}

func lookupEnvLanguage() (*Language, bool) {
	return LookupLanguage(env.LanguageCode())
}
