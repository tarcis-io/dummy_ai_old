package server

import (
	"embed"
	"text/template"
)

type (
	pageData struct {
		wasmPath string
	}
)

var (
	//go:embed template.html
	htmlTemplateFS embed.FS
	htmlTemplate   = template.Must(template.ParseFS(htmlTemplateFS, "template.html"))
)

func Run() {
}
