package dom

type (
	Document struct {
		*DOM
	}
)

func (d *Document) CreateElement(tagName string) *DOM {
	return d.Call("createElement", tagName)
}

func GetDocument() *Document {
	return GetWindow().Document()
}
