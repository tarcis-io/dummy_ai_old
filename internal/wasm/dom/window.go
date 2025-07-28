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

func GetWindow() *Window {
	return &Window{
		DOM: GetGlobal(),
	}
}
