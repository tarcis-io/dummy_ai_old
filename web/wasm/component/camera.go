package component

import (
	"syscall/js"
)

func NewCamera() js.Value {

	camera := js.Global().Get("document").Call("createElement", "div")
	camera.Set("innerHTML", "Camera")

	return camera
}
