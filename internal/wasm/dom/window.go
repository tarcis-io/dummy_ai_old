package dom

type (
	// Window represents the JavaScript Window object.
	Window struct {
		*DOM
	}
)

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

// GetWindow returns the current Window object.
func GetWindow() *Window {
	return &Window{
		DOM: GetGlobal(),
	}
}
