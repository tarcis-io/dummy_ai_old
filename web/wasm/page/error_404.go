package main

import (
	"syscall/js"
)

func main() {

	h2 := js.Global().Get("document").Call("createElement", "h2")
	h2.Set("innerHTML", "Error 404!")

	app := js.Global().Get("document").Call("getElementById", "app")
	app.Call("appendChild", h2)
}
