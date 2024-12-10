package main

import (
	"syscall/js"
)

import (
	"dummy_ai/web/wasm/component"
)

func main() {

	page := js.Global().Get("document").Call("createElement", "div")
	page.Call("appendChild", component.NewCamera())

	app := component.CreateApp(page)

	body := js.Global().Get("document").Get("body")
	body.Call("appendChild", app)
}
