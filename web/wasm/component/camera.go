package component

import (
	"syscall/js"
)

import (
	"dummy_ai/web/wasm/util"
)

func CreateCamera() js.Value {

	js.Global().Call("setTimeout", onCameraLoad(), 1500)
	return createCameraLoading()
}

func createCameraLoading() js.Value {

	loading := js.Global().Get("document").Call("createElement", "cds-loading")
	loading.Set("type", "small")
	loading.Set("assistiveText", util.CameraLoading())

	cameraTile := createCameraTile("camera-loading")
	cameraTile.Call("appendChild", loading)

	return cameraTile
}

func createCameraError(title string, text string) js.Value {

	h4 := js.Global().Get("document").Call("createElement", "h4")
	h4.Set("innerHTML", title)

	p := js.Global().Get("document").Call("createElement", "p")
	p.Set("innerHTML", text)

	onButtonClick := js.FuncOf(func(this js.Value, args []js.Value) any {

		js.Global().Get("location").Call("reload")
		return nil
	})

	button := js.Global().Get("document").Call("createElement", "cds-button")
	button.Set("kind", "danger")
	button.Set("innerHTML", util.CameraReloadPage())
	button.Call("addEventListener", "click", onButtonClick)

	cameraTile := createCameraTile("camera-error")
	cameraTile.Call("appendChild", h4)
	cameraTile.Call("appendChild", p)
	cameraTile.Call("appendChild", button)

	return cameraTile
}

func createCameraTile(id string) js.Value {

	tile := js.Global().Get("document").Call("createElement", "cds-tile")
	tile.Set("id", id)

	return tile
}

func onCameraLoad() js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) any {

		var (
			cameraLoading = js.Global().Get("document").Call("getElementById", "camera-loading")
			mediaDevices  = js.Global().Get("navigator").Get("mediaDevices")
		)

		if mediaDevices.IsUndefined() {

			cameraLoading.Call("replaceWith", createCameraError(util.CameraErrorNotSupportedTitle(), util.CameraErrorNotSupportedText()))
			return nil
		}

		getUserMedia := mediaDevices.Get("getUserMedia")

		if getUserMedia.IsUndefined() {

			cameraLoading.Call("replaceWith", createCameraError(util.CameraErrorNotSupportedTitle(), util.CameraErrorNotSupportedText()))
			return nil
		}

		return nil
	})
}
