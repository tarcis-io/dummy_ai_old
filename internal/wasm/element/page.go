package element

import (
	"dummy_ai/internal/wasm/util"
)

func CreatePage(content *Element) *Element {
	page := Create("div")
	page.AppendChild(CreatePageHeader())
	page.AppendChild(content)
	return page
}

func CreatePageHeader() *Element {
	headerName := Create("cds-header-name")
	headerName.SetTextContent(util.App())
	header := Create("cds-header")
	header.AppendChild(headerName)
	return header
}
