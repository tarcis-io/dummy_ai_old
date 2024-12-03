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
	inlineLoading.Call("setAttribute", "assistive-text", util.CameraLoading())

	tile := js.Global().Get("document").Call("createElement", "cds-tile")
	tile.Call("appendChild", inlineLoading)

	return tile
}
