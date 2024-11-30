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

	app := js.Global().Get("document").Call("getElementById", "app")
	app.Call("appendChild", h2)
	app.Call("appendChild", camera)
}
