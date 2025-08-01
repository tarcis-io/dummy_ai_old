package main

import (
	"dummy_ai/internal/wasm/dom"
	"dummy_ai/internal/wasm/util"
)

func main() {
	h2 := dom.GetDocument().CreateElement("h2")
	h2.Set("innerText", util.App())

	body := dom.GetDocument().Get("body")
	body.Call("appendChild", h2)
}
