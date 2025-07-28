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

func (w *Window) LocalStorage() *LocalStorage {
	return &LocalStorage{
		DOM: w.Get("localStorage"),
	}
}

func (w *Window) Navigator() *Navigator {
	return &Navigator{
		DOM: w.Get("navigator"),
	}
}

func GetWindow() *Window {
	return &Window{
		DOM: GetGlobal(),
	}
}
