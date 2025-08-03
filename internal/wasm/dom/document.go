package dom

type (
	// Document represents the JavaScript Document object.
	Document struct {
		*DOM
	}
)

// CreateElement creates a new HTML element with the specified tag name.
func (d *Document) CreateElement(tagName string) *DOM {
	return d.Call("createElement", tagName)
}

// GetDocument returns the current JavaScript Document object.
func GetDocument() *Document {
	return GetWindow().Document()
}
