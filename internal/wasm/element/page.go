package element

import (
	"dummy_ai/internal/wasm/util"
)

func CreatePage(content *Element) *Element {
	div := Create("div")
	div.Set("className", "cds-theme-zone-g100")
	div.SetStyle("fontFamily", "'IBM Plex Sans', 'Helvetica Neue', Arial, sans-serif")
	div.SetStyle("position", "fixed")
	div.SetStyle("top", 0)
	div.SetStyle("right", 0)
	div.SetStyle("bottom", 0)
	div.SetStyle("left", 0)
	div.SetStyle("overflowY", "auto")
	div.SetStyleProperty("--cds-background", "black")
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
	div.SetStyle("margin", "var(--cds-spacing-03)")
	div.SetStyle("paddingTop", "var(--cds-spacing-09)")
	div.AppendChild(content)
	return div
}
