package component

import (
	"syscall/js"
)

import (
	"dummy_ai/web/wasm/util"
)

func NewCamera() js.Value {

	return newCameraLoading()
}

func newCameraLoading() js.Value {

	inlineLoading := js.Global().Get("document").Call("createElement", "cds-inline-loading")
	inlineLoading.Set("status", "active")
	inlineLoading.Set("innerHTML", util.Message(util.CameraLoading))

	return inlineLoading
}
