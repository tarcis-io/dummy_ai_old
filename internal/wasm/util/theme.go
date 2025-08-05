package util

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

func CurrentTheme() *Theme {
	return currentTheme
}
