package component

import (
	"syscall/js"
)

func NewCamera() js.Value {

	return newCameraLoading()
}

func newCameraLoading() js.Value {

	cameraLoading := js.Global().Get("document").Call("createElement", "div")
	cameraLoading.Set("innerHTML", "Camera Loading")

	return cameraLoading
}

func newCameraError() js.Value {

	cameraError := js.Global().Get("document").Call("createElement", "div")
	cameraError.Set("innerHTML", "Camera Error")

	return cameraError
}

func newCameraStreaming() js.Value {

	cameraStreaming := js.Global().Get("document").Call("createElement", "div")
	cameraStreaming.Set("innerHTML", "Camera Streaming")

	return cameraStreaming
}
