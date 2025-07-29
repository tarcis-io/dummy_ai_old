package util

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
