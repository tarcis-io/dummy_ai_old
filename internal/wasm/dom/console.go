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

// Info outputs the specified message to the console
// with the info log level.
func (c *Console) Info(message string) {
	c.Call("info", message)
}

// Warn outputs the specified message to the console
// with the warning log level.
func (c *Console) Warn(message string) {
	c.Call("warn", message)
}

// Error outputs the specified message to the console
// with the error log level.
func (c *Console) Error(message string) {
	c.Call("error", message)
}
