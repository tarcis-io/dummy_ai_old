package component

import (
	"syscall/js"
)

func CreateApp(page js.Value) js.Value {

	div := js.Global().Get("document").Call("createElement", "div")
	div.Call("appendChild", page)

	return div
}
