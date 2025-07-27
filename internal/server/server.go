package server

import (
	"embed"
	"net/http"
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

	staticFileServer = http.FileServer(http.Dir("./static"))
)

func Run() {
}
