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

	staticFileServerDir = http.Dir("./static")
	staticFileServer    = http.FileServer(staticFileServerDir)
)

func Run() {
	router := http.NewServeMux()
	router.HandleFunc("/", handleRequest)
	listenAndServe(env.ServerAddress(), router)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	d, ok := pageRoutes[p]
	if ok {
		renderPage(w, d)
		return
	}
	_, err := staticFileServerDir.Open(p)
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	serveStaticFile(w, r)
}

func renderPage(w http.ResponseWriter, p *pageData) {
	err := htmlTemplate.Execute(w, p)
	if err != nil {
		log.Printf("ERROR: Failed to render page %s: %v", p.wasmPath, err)
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
	}
}

func serveStaticFile(w http.ResponseWriter, r *http.Request) {
	staticFileServer.ServeHTTP(w, r)
}

func listenAndServe(addr string, handler http.Handler) {
	log.Printf("INFO: Server is running on %s", addr)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatalf("FATAL: Failed to run server: %v", err)
	}
}
