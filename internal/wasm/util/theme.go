package util

import (
	"dummy_ai/internal/env"
	"dummy_ai/internal/wasm/dom"
)

type (
	Theme struct {
		theme string
	}
)

func (t *Theme) Theme() string {
	return t.theme
}

var (
	black = &Theme{
		theme: "black",
	}

	white = &Theme{
		theme: "white",
	}

	fallbackTheme = black

	supportedThemes = []*Theme{
		black,
		white,
	}

	supportedThemesMap = map[string]*Theme{
		black.theme: black,
		white.theme: white,
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

func LookupTheme() *Theme {
	v, ok := lookupLocalStorageTheme()
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

func lookupEnvTheme() (*Theme, bool) {
	return lookupTheme(env.Theme())
}

func lookupTheme(theme string) (*Theme, bool) {
	v, ok := supportedThemesMap[theme]
	return v, ok
}
