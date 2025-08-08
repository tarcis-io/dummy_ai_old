package dom

type (
	// Window represents the JavaScript Window object.
	Window struct {
		*DOM
	}
)

// Console returns the Window's Console object.
func (w *Window) Console() *Console {
	return &Console{
		DOM: w.Get("console"),
	}
}

// Document returns the Window's Document object.
func (w *Window) Document() *Document {
	return &Document{
		DOM: w.Get("document"),
	}
}

// Navigator returns the Window's Navigator object.
func (w *Window) Navigator() *Navigator {
	return &Navigator{
		DOM: w.Get("navigator"),
	}
}

// LocalStorage returns the Window's local Storage object.
func (w *Window) LocalStorage() *Storage {
	return &Storage{
		DOM: w.Get("localStorage"),
	}
}

// SessionStorage returns the Window's session Storage object.
func (w *Window) SessionStorage() *Storage {
	return &Storage{
		DOM: w.Get("sessionStorage"),
	}
}

// MatchMedia returns a [*MediaQueryList] object
// representing the specified media query string.
func (w *Window) MatchMedia(query string) *MediaQueryList {
	return &MediaQueryList{
		DOM: w.Call("matchMedia", query),
	}
}

// GetWindow returns the current Window object.
func GetWindow() *Window {
	return &Window{
		DOM: Global(),
	}
}
