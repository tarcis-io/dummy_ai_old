package component

import (
	"syscall/js"
)

import (
	"dummy_ai/web/wasm/util"
)

func CreateCamera() js.Value {

	return createCameraError("Error!")
}

func createCameraLoading() js.Value {

	loading := js.Global().Get("document").Call("createElement", "cds-loading")
	loading.Set("type", "small")
	loading.Set("assistiveText", util.CameraLoading())

	cameraTile := createCameraTile("camera-loading")
	cameraTile.Call("appendChild", loading)

	return cameraTile
}

func createCameraError(title string) js.Value {

	div := js.Global().Get("document").Call("createElement", "div")
	div.Set("innerHTML", title)

	cameraTile := createCameraTile("camera-loading")
	cameraTile.Call("appendChild", div)

	return cameraTile
}

func createCameraTile(id string) js.Value {

	tile := js.Global().Get("document").Call("createElement", "cds-tile")
	tile.Set("id", id)

	return tile
}
