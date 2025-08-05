package element

func CreatePage(content *Element) *Element {
	page := Create("div")
	page.AppendChild(content)
	return page
}
