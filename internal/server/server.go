package server

import (
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

	httpTemplate = template.Must(template.ParseFiles("template.html"))
)

func Run() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	err := http.ListenAndServe(env.ServerAddress(), nil)
	if err != nil {
		log.Fatalf("FATAL: Failed to start server: %v", err)
	}
}
