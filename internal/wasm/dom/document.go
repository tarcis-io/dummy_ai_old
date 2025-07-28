package dom

type (
	Document struct {
		*DOM
	}
)

func (d *Document) CreateElement(element string) *DOM {
	return d.Call("createElement", element)
}

func GetDocument() *Document {
	return GetWindow().Document()
}
