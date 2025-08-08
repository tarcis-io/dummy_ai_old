package dom

type (
	// Document represents the JavaScript Document object.
	Document struct {
		*DOM
	}
)

// DocumentElement returns the <html> element of the document.
func (d *Document) DocumentElement() *DOM {
	return d.Get("documentElement")
}

// Head returns the <head> element of the document.
func (d *Document) Head() *DOM {
	return d.Get("head")
}

// Body returns the <body> element of the document.
func (d *Document) Body() *DOM {
	return d.Get("body")
}

// GetElementById returns the element with the specified id.
func (d *Document) GetElementById(id string) *DOM {
	return d.Call("getElementById", id)
}

// GetElementsByClassName returns all elements with the specified class name.
func (d *Document) GetElementsByClassName(className string) *DOM {
	return d.Call("getElementsByClassName", className)
}

// GetElementsByTagName returns all elements with the specified tag name.
func (d *Document) GetElementsByTagName(tagName string) *DOM {
	return d.Call("getElementsByTagName", tagName)
}

// CreateElement creates a new HTML element with the specified tag name.
func (d *Document) CreateElement(tagName string) *DOM {
	return d.Call("createElement", tagName)
}
