package server

import (
	"embed"
	"log"
	"net/http"
	"text/template"
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

	staticFileServer = http.FileServer(http.Dir("./static"))
)

func Run() {
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
