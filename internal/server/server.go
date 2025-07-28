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
	if r.Method != http.MethodGet {
		renderPageError(w, http.StatusMethodNotAllowed)
		return
	}
	p := r.URL.Path
	d, ok := pageRoutes[p]
	if ok {
		renderPage(w, d)
		return
	}
	_, err := staticFileServerDir.Open(p)
	if err != nil {
		renderPageError(w, http.StatusNotFound)
		return
	}
	serveStaticFile(w, r)
}

func renderPage(w http.ResponseWriter, p *pageData) {
	err := htmlTemplate.Execute(w, p)
	if err != nil {
		log.Printf("ERROR: Failed to render page %s: %v", p.wasmPath, err)
		renderPageError(w, http.StatusInternalServerError)
	}
}

func renderPageError(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func serveStaticFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=86400")
	staticFileServer.ServeHTTP(w, r)
}

func listenAndServe(addr string, handler http.Handler) {
	log.Printf("INFO: Server is running on %s", addr)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatalf("FATAL: Failed to run server: %v", err)
	}
}
