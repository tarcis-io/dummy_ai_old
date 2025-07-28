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

func GetWindow() *Window {
	return &Window{
		DOM: GetGlobal(),
	}
}
