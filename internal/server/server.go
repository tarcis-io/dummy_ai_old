package server

import (
	"embed"
	"net/http"
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

func listenAndServe(serverAddress string, router http.Handler) {
	err := http.ListenAndServe(serverAddress, router)
	if err != nil {
	}
}
