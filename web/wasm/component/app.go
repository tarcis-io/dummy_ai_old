package component

import (
	"syscall/js"
)

func CreateApp(page js.Value) js.Value {

	div := js.Global().Get("document").Call("createElement", "div")
	div.Get("classList").Call("add", "cds-theme-zone-g100")
	div.Get("style").Set("font-family", "'IBM Plex Sans', 'Helvetica Neue', Arial, sans-serif")
	div.Get("style").Set("position", "fixed")
	div.Get("style").Set("top", "0")
	div.Get("style").Set("right", "0")
	div.Get("style").Set("bottom", "0")
	div.Get("style").Set("left", "0")
	div.Get("style").Set("background-color", "var(--cds-background)")
	div.Get("style").Set("color", "var(--cds-text-primary)")
	div.Call("appendChild", page)

	return div
}
