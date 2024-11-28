package component

import (
	"syscall/js"
)

func NewCamera() js.Value {

	return newCameraLoading()
}

func newCameraLoading() js.Value {

	cameraLoading := js.Global().Get("document").Call("createElement", "div")
	cameraLoading.Set("innerHTML", "CameraLoading")

	return cameraLoading
}
