package main

import (
	"dummy_ai/web/wasm/component"
	"syscall/js"
)

func main() {

	h2 := js.Global().Get("document").Call("createElement", "h2")
	h2.Set("innerHTML", "Error 404!")

	app := component.CreateApp(h2)

	body := js.Global().Get("document").Get("body")
	body.Call("appendChild", app)
}
