package server

import (
	"embed"
	"log"
	"net/http"
	"text/template"

	"dummy_ai/internal/env"
)

var (
	wasmRoutes = map[string]string{
		"/":      "/wasm/home.wasm",
		"/about": "/wasm/about.wasm",
	}
	wasmPathError404 = "/wasm/error_404.wasm"
	wasmPathError500 = "/wasm/error_500.wasm"

	//go:embed template.html
	httpTemplateFS embed.FS
	httpTemplate   = template.Must(template.ParseFS(httpTemplateFS, "template.html"))
)

func Run() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	err := http.ListenAndServe(env.ServerAddress(), nil)
	if err != nil {
		log.Fatalf("FATAL: Failed to start server: %v", err)
	}
}
