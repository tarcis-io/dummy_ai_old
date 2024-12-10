package component

import (
	"syscall/js"
)

func CreateCamera() js.Value {

	return createCameraLoading()
}

func createCameraLoading() js.Value {

	loading := js.Global().Get("document").Call("createElement", "cds-loading")

	tile := js.Global().Get("document").Call("createElement", "cds-tile")
	tile.Set("id", "camera-loading")
	tile.Call("appendChild", loading)

	return tile
}
