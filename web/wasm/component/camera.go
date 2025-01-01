package component

import (
	"syscall/js"
)

import (
	"dummy_ai/web/wasm/util"
)

func CreateCamera() js.Value {

	return createCameraError("Error!!!")
}

func createCameraLoading() js.Value {

	loading := js.Global().Get("document").Call("createElement", "cds-loading")
	loading.Set("type", "small")
	loading.Set("assistiveText", util.CameraLoading())

	return createCameraTile("camera-loading", loading)
}

func createCameraError(title string) js.Value {

	div := js.Global().Get("document").Call("createElement", "div")
	div.Set("innerHTML", title)

	return createCameraTile("camera-error", div)
}

func createCameraTile(id string, child js.Value) js.Value {

	tile := js.Global().Get("document").Call("createElement", "cds-tile")
	tile.Set("id", id)
	tile.Call("appendChild", child)

	return tile
}
