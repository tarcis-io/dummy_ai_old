package server

import (
	"bytes"
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

func pageHandler(responseWriter http.ResponseWriter, wasmPath string, statusCode int) {
	var buffer bytes.Buffer
	err := httpTemplate.Execute(&buffer, wasmPath)
	if err != nil {
		log.Printf("ERROR: Failed to execute template %s: %v", wasmPath, err)
		if wasmPath == wasmPathError500 {
			http.Error(responseWriter, "An unexpected internal error occurred.", http.StatusInternalServerError)
			return
		}
		pageHandler(responseWriter, wasmPathError500, http.StatusInternalServerError)
		return
	}
	responseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")
	responseWriter.WriteHeader(statusCode)
	buffer.WriteTo(responseWriter)
}
