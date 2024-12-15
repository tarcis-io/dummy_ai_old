package component

import (
	"syscall/js"
)

import (
	"dummy_ai/web/wasm/util"
)

func CreateApp(page js.Value) js.Value {

	div := js.Global().Get("document").Call("createElement", "div")
	div.Set("id", "app")
	div.Set("className", "cds-theme-zone-g100")
	div.Call("appendChild", createAppHeader())
	div.Call("appendChild", page)

	return div
}

func createAppHeader() js.Value {

	headerName := js.Global().Get("document").Call("createElement", "cds-header-name")
	headerName.Set("innerHTML", util.App())

	header := js.Global().Get("document").Call("createElement", "cds-header")
	header.Call("appendChild", headerName)

	return header
}

func createAppPageContainer(page js.Value) js.Value {

	div := js.Global().Get("document").Call("createElement", "div")
	div.Set("id", "app-page-container")
	div.Call("appendChild", page)

	return div
}
