// Package util provides shared utilities for the application.
package util

// init runs on startup.
// It configures the package's initial state.
func init() {
	currentLanguage = LookupLanguage()
	currentLocale = FetchLocale(currentLanguage)
	currentTheme = LookupTheme()
}
