// Package server provides a simple web server
// for serving HTML pages and static assets.
package server

import (
	"embed"
	"log"
	"net/http"
	"text/template"

	"dummy_ai/internal/env"
)

type (
	// pageData holds data for rendering HTML pages.
	pageData struct {
		Title    string
		WASMPath string
	}
)

var (
	// pageRoutes maps URL paths to [pageData].
	pageRoutes = map[string]*pageData{
		"/": {
			Title:    "DummyAI",
			WASMPath: "/wasm/home.wasm",
		},
		"/about": {
			Title:    "DummyAI",
			WASMPath: "/wasm/about.wasm",
		},
	}

	// pageTemplateFS embeds the template.html file.
	//go:embed template.html
	pageTemplateFS embed.FS

	// pageTemplate is the parsed template.html file as a [*template.Template].
	pageTemplate = template.Must(template.ParseFS(pageTemplateFS, "template.html"))

	// staticFileServer serves static files from the ./static directory.
	staticFileServer = http.FileServer(http.Dir("./static"))
)

// Run starts the web server.
func Run() {
	router := http.NewServeMux()
	router.HandleFunc("/", rootHandler)
	listenAndServe(env.ServerAddress(), router)
}

// rootHandler handles and routes incoming page requests
// and serves static files.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("[error] Method not allowed: %s", r.Method)
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	p, ok := pageRoutes[r.URL.Path]
	if ok {
		renderPage(w, p)
		return
	}
	serveStaticFile(w, r)
}

// renderPage executes the template and writes to the response.
// Returns error 500 on template execution error.
func renderPage(w http.ResponseWriter, p *pageData) {
	err := pageTemplate.Execute(w, p)
	if err != nil {
		log.Printf("[error] Failed to render page %s: %v", p.WASMPath, err)
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
	}
}

// serveStaticFile writes static file contents to the response.
// Returns error 404 on file not found.
func serveStaticFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=86400")
	staticFileServer.ServeHTTP(w, r)
}

// listenAndServe starts the web server.
// Exits the application on server startup error.
func listenAndServe(addr string, handler http.Handler) {
	log.Printf("[info] Server is running on %s", addr)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatalf("[fatal] Failed to run server: %v", err)
	}
}
