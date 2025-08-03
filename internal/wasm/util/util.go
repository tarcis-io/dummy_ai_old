package util

func init() {
	currentLanguage = LookupLanguage()
	currentLocale = FetchCurrentLocale()
}
