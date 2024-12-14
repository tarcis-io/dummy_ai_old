package component

import (
	"syscall/js"
)

import (
	"dummy_ai/web/wasm/util"
)

func CreateCamera() js.Value {

	return createCameraLoading()
}

func createCameraLoading() js.Value {

	loading := js.Global().Get("document").Call("createElement", "cds-loading")
	loading.Set("assistiveText", util.CameraLoading())

	tile := js.Global().Get("document").Call("createElement", "cds-tile")
	tile.Set("id", "camera-loading")
	tile.Call("appendChild", loading)

	return tile
}
