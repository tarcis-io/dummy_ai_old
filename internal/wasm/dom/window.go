package dom

type (
	Window struct {
		*DOM
	}
)

func (w *Window) Document() *Document {
	return &Document{
		DOM: w.Get("document"),
	}
}

func (w *Window) Navigator() *Navigator {
	return &Navigator{
		DOM: w.Get("navigator"),
	}
}

func (w *Window) LocalStorage() *Storage {
	return &Storage{
		DOM: w.Get("localStorage"),
	}
}

func (w *Window) SessionStorage() *Storage {
	return &Storage{
		DOM: w.Get("sessionStorage"),
	}
}

func GetWindow() *Window {
	return &Window{
		DOM: GetGlobal(),
	}
}
