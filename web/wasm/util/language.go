package util

const (
	English    = "en"
	Spanish    = "es"
	Portuguese = "pt"
)

const (
	DefaultLanguage = English
)

var (
	language = preferredLanguage()
)

func Language() string {

	return language
}

func preferredLanguage() string {

	return English
}
