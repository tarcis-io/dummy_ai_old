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

// SetId sets the id attribute of the element.
func (e *Element) SetId(id string) {
	e.dom.Set("id", id)
}

// SetClassName sets the className attribute of the element.
func (e *Element) SetClassName(className string) {
	e.dom.Set("className", className)
}

// Create creates a new HTML element with the specified tag name.
func Create(tagName string) *Element {
	return &Element{
		dom: dom.GetDocument().CreateElement(tagName),
	}
}
