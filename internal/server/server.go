package server

import (
	"embed"
	"log"
	"net/http"
	"text/template"

	"dummy_ai/internal/env"
)

var (
	//go:embed template.html
	httpTemplateFS embed.FS
	httpTemplate   = template.Must(template.ParseFS(httpTemplateFS, "template.html"))
)

func Run() {
	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	serverAddress := env.ServerAddress()
	log.Printf("INFO: Server starting on %s", serverAddress)
	err := http.ListenAndServe(serverAddress, router)
	if err != nil {
		log.Fatalf("FATAL: Failed to start server: %v", err)
	}
}
