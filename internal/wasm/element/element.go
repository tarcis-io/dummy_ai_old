// Package element provides a Go-friendly wrapper
// for interacting with HTML elements.
package element

import (
	"dummy_ai/internal/wasm/dom"
)

type (
	// Element represents an HTML element.
	Element struct {
		dom *dom.DOM
	}
)

// Set sets a property of the element.
func (e *Element) Set(property string, value any) {
	e.dom.Set(property, value)
}

// SetId sets the id property of the element.
func (e *Element) SetId(id string) {
	e.Set("id", id)
}

// AppendChild appends a child element to the element.
func (e *Element) AppendChild(element *Element) {
	e.dom.Call("appendChild", element.dom)
}

// Create creates a new HTML element with the specified tag name.
func Create(tagName string) *Element {
	return &Element{
		dom: dom.GetDocument().CreateElement(tagName),
	}
}
