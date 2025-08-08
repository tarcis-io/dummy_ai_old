// Package util provides shared utilities for the application.
package util

// init runs on startup.
// It configures the package's initial state.
// 1. The current language.
// 2. The current locale.
// 3. The current theme.
func init() {
	currentLanguage = LookupLanguage()
	currentLocale = FetchLocale(currentLanguage)
	currentTheme = LookupTheme()
}
