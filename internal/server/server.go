package server

import (
	"embed"
	"text/template"
)

type (
	pageData struct {
		title    string
		wasmPath string
	}
)

var (
	pageRoutes = map[string]*pageData{
		"/": {
			title:    "DummyAI",
			wasmPath: "/wasm/home.wasm",
		},
		"/about": {
			title:    "DummyAI",
			wasmPath: "/wasm/about.wasm",
		},
	}

	//go:embed template.html
	htmlTemplateFS embed.FS
	htmlTemplate   = template.Must(template.ParseFS(htmlTemplateFS, "template.html"))
)

func Run() {
}
