package dom

type (
	// Console represents the JavaScript Console object.
	Console struct {
		*DOM
	}
)

// Log outputs the specified message to the console.
func (c *Console) Log(message string) {
	c.Call("log", message)
}
