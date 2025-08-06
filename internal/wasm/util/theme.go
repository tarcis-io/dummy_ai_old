package util

import (
	"dummy_ai/internal/env"
	"dummy_ai/internal/wasm/dom"
)

type (
	// Theme represents a supported theme.
	Theme struct {
		code string
		name string
	}
)

// Code returns the theme code.
func (t *Theme) Code() string {
	return t.code
}

// Name returns the theme name.
func (t *Theme) Name() string {
	return t.name
}

var (
	// darkTheme represents the dark theme.
	darkTheme = &Theme{
		code: "dark",
		name: "Dark",
	}

	// lightTheme represents the light theme.
	lightTheme = &Theme{
		code: "light",
		name: "Light",
	}

	// fallbackTheme is the default theme used
	// if no other theme can be determined.
	fallbackTheme = darkTheme

	// supportedThemes is a slice of all supported themes.
	supportedThemes = []*Theme{
		darkTheme,
		lightTheme,
	}

	// supportedThemesMap provides a quick lookup for supported themes.
	supportedThemesMap = map[string]*Theme{
		darkTheme.code:  darkTheme,
		lightTheme.code: lightTheme,
	}

	// currentTheme holds the currently selected theme.
	currentTheme *Theme
)

// DarkTheme returns the dark theme.
func DarkTheme() *Theme {
	return darkTheme
}

// LightTheme returns the light theme.
func LightTheme() *Theme {
	return lightTheme
}

// FallbackTheme returns the fallback theme.
func FallbackTheme() *Theme {
	return fallbackTheme
}

// SupportedThemes returns a slice of all supported themes.
func SupportedThemes() []*Theme {
	return supportedThemes
}

// SupportedThemesMap returns a map of all supported themes.
func SupportedThemesMap() map[string]*Theme {
	return supportedThemesMap
}

// CurrentTheme returns the currently selected theme.
func CurrentTheme() *Theme {
	return currentTheme
}

// SetTheme persists the chosen theme to the browser's local storage.
func SetTheme(theme *Theme) {
	dom.GetLocalStorage().SetItem("theme", theme.code)
}

// LookupTheme determines the user's preferred theme by checking:
// 1. The browser's local storage.
// 2. The browser's preferred color scheme.
// 3. The environment variables.
// It returns the fallback theme if no supported theme is found.
func LookupTheme() *Theme {
	v, ok := lookupLocalStorageTheme()
	if ok {
		return v
	}
	v, ok = lookupPreferredTheme()
	if ok {
		return v
	}
	v, ok = lookupEnvTheme()
	if ok {
		return v
	}
	return fallbackTheme
}

// lookupLocalStorageTheme looks up the theme from the local storage.
func lookupLocalStorageTheme() (*Theme, bool) {
	v, ok := dom.GetLocalStorage().GetItem("theme")
	if ok {
		return lookupTheme(v)
	}
	return nil, false
}

// lookupPreferredTheme looks up the theme from the browser's preferred color scheme.
func lookupPreferredTheme() (*Theme, bool) {
	if dom.GetWindow().MatchMedia("(prefers-color-scheme: dark)") {
		return darkTheme, true
	}
	if dom.GetWindow().MatchMedia("(prefers-color-scheme: light)") {
		return lightTheme, true
	}
	return nil, false
}

// lookupEnvTheme looks up the theme from the environment variables.
func lookupEnvTheme() (*Theme, bool) {
	return lookupTheme(env.ThemeCode())
}

// lookupTheme looks up a theme by its code.
func lookupTheme(themeCode string) (*Theme, bool) {
	v, ok := supportedThemesMap[themeCode]
	return v, ok
}
