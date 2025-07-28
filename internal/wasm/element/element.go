package element

import (
	"dummy_ai/internal/wasm/dom"
)

type (
	Element struct {
		dom *dom.DOM
	}
)

func (e *Element) SetId(id string) {
	e.dom.Set("id", id)
}

func (e *Element) SetClassName(className string) {
	e.dom.Set("className", className)
}

func Create(tagName string) *Element {
	return &Element{
		dom: dom.GetDocument().CreateElement(tagName),
	}
}
