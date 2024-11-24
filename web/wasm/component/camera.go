package component

import (
	"syscall/js"
)

func NewCamera() js.Value {

	return NewCameraLoading()
}

func NewCameraLoading() js.Value {

	cameraLoading := js.Global().Get("document").Call("createElement", "div")
	cameraLoading.Set("innerHTML", "CameraLoading")

	return cameraLoading
}

func NewCameraError() js.Value {

	cameraError := js.Global().Get("document").Call("createElement", "div")
	cameraError.Set("innerHTML", "CameraError")

	return cameraError
}
