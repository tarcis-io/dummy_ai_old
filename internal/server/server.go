package server

import (
	"embed"
	"log"
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
	log.Printf("INFO: Server starting on %s", serverAddress)
	err := http.ListenAndServe(serverAddress, router)
	if err != nil {
		log.Fatalf("FATAL: Failed to start server: %v", err)
	}
}
