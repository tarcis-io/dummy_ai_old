package server

import (
	"embed"
	"log"
	"net/http"
	"text/template"

	"dummy_ai/internal/env"
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
	router := http.NewServeMux()
	listenAndServe(env.ServerAddress(), router)
}

func listenAndServe(addr string, handler http.Handler) {
	log.Printf("INFO: Server running on %s", addr)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatalf("FATAL: Failed to run server: %v", err)
	}
}
