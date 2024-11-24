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

	body := js.Global().Get("document").Get("body")
	body.Call("appendChild", h2)
	body.Call("appendChild", camera)
}
