package util

import (
	"dummy_ai/internal/env"
	"dummy_ai/internal/wasm/dom"
)

type (
	Theme struct {
		code string
		name string
	}
)

func (t *Theme) Code() string {
	return t.code
}

func (t *Theme) Name() string {
	return t.name
}

var (
	black = &Theme{
		code: "black",
		name: "Black",
	}

	white = &Theme{
		code: "white",
		name: "White",
	}

	fallbackTheme = black

	supportedThemes = []*Theme{
		black,
		white,
	}

	supportedThemesMap = map[string]*Theme{
		black.code: black,
		white.code: white,
	}

	currentTheme *Theme
)

func BlackTheme() *Theme {
	return black
}

func WhiteTheme() *Theme {
	return white
}

func FallbackTheme() *Theme {
	return fallbackTheme
}

func SupportedThemes() []*Theme {
	return supportedThemes
}

func SupportedThemesMap() map[string]*Theme {
	return supportedThemesMap
}

func CurrentTheme() *Theme {
	return currentTheme
}

func SetTheme(theme *Theme) {
	dom.GetLocalStorage().SetItem("theme", theme.code)
}

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

func lookupLocalStorageTheme() (*Theme, bool) {
	v, ok := dom.GetLocalStorage().GetItem("theme")
	if ok {
		return lookupTheme(v)
	}
	return nil, false
}

func lookupPreferredTheme() (*Theme, bool) {
	if dom.GetWindow().MatchMedia("(prefers-color-scheme: dark)") {
		return black, true
	}
	if dom.GetWindow().MatchMedia("(prefers-color-scheme: light)") {
		return white, true
	}
	return nil, false
}

func lookupEnvTheme() (*Theme, bool) {
	return lookupTheme(env.Theme())
}

func lookupTheme(theme string) (*Theme, bool) {
	v, ok := supportedThemesMap[theme]
	return v, ok
}
