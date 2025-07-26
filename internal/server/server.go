package server

import (
	"embed"
	"log"
	"net/http"
	"text/template"

	"dummy_ai/internal/env"
)

type (
	page struct {
		wasmPath string
	}
)

var (
	//go:embed template.html
	httpTemplateFS embed.FS
	httpTemplate   = template.Must(template.ParseFS(httpTemplateFS, "template.html"))
)

func Run() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderPage(w, &page{
			wasmPath: "/wasm/home.wasm",
		})
	})
	router.Handle("/", http.FileServer(http.Dir("./static")))
	serverAddress := env.ServerAddress()
	log.Printf("INFO: Server starting on %s", serverAddress)
	err := http.ListenAndServe(serverAddress, router)
	if err != nil {
		log.Fatalf("FATAL: Failed to start server: %v", err)
	}
}

func renderPage(w http.ResponseWriter, p *page) {
	err := httpTemplate.Execute(w, p)
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
	}
}
