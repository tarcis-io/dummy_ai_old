package element

import (
	"dummy_ai/internal/wasm/util"
)

func CreatePage(content *Element) *Element {
	div := Create("div")
	div.AppendChild(createPageHeader())
	div.AppendChild(createPageContentContainer(content))
	return div
}

func createPageHeader() *Element {
	header := Create("cds-header")
	header.AppendChild(createPageHeaderName())
	return header
}

func createPageHeaderName() *Element {
	headerName := Create("cds-header-name")
	headerName.Set("textContent", util.App())
	return headerName
}

func createPageContentContainer(content *Element) *Element {
	div := Create("div")
	div.AppendChild(content)
	return div
}
