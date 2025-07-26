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
	staticFileServer = http.FileServer(http.Dir("./static"))

	//go:embed template.html
	httpTemplateFS embed.FS
	httpTemplate   = template.Must(template.ParseFS(httpTemplateFS, "template.html"))
)

func Run() {
	router := http.NewServeMux()
	router.HandleFunc("/", requestHandler)
	serverAddress := env.ServerAddress()
	log.Printf("INFO: Server starting on %s", serverAddress)
	err := http.ListenAndServe(serverAddress, router)
	if err != nil {
		log.Fatalf("FATAL: Failed to start server: %v", err)
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path
	if requestPath == "/" {
		renderPage(w, &page{
			wasmPath: "/wasm/home.wasm",
		})
		return
	}
	staticFileServer.ServeHTTP(w, r)
}

func renderPage(w http.ResponseWriter, p *page) {
	err := httpTemplate.Execute(w, p)
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
	}
}
