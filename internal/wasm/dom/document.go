package dom

type (
	// Document represents the JavaScript Document object.
	Document struct {
		*DOM
	}
)

// CreateElement creates a new HTML element with the specified tag name and returns it as a new *DOM object.
func (d *Document) CreateElement(tagName string) *DOM {
	return d.Call("createElement", tagName)
}

// GetDocument returns the web page's current *Document object.
func GetDocument() *Document {
	return GetWindow().Document()
}
