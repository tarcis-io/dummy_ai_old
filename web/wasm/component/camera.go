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
	loading.Set("type", "small")
	loading.Set("assistiveText", util.CameraLoading())

	tile := js.Global().Get("document").Call("createElement", "cds-tile")
	tile.Set("id", "camera-loading")
	tile.Call("appendChild", loading)

	return tile
}

func createCameraError(title string, text string) js.Value {

	div := js.Global().Get("document").Call("createElement", "div")
	div.Set("innerHTML", title)

	return div
}
