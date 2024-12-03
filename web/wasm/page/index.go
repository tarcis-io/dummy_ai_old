package main

import (
	"syscall/js"
)

import (
	"dummy_ai/web/wasm/component"
)

func main() {

	h2 := js.Global().Get("document").Call("createElement", "h2")
	h2.Set("innerHTML", "Hello World, DummyAI!")

	camera := component.NewCamera()

	page := js.Global().Get("document").Call("createElement", "div")
	page.Call("appendChild", h2)
	page.Call("appendChild", camera)

	app := component.CreateApp(page)

	body := js.Global().Get("document").Get("body")
	body.Call("appendChild", app)
}
