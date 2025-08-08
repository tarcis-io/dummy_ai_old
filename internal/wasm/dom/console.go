package dom

type (
	Console struct {
		*DOM
	}
)

func (c *Console) Log(text string) {
	c.Call("log", text)
}

func GetConsole() *Console {
	return GetWindow().Console()
}
