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
	// pageData represents the data needed to render an HTML page.
	pageData struct {
		Title    string
		WASMPath string
	}
)

var (
	// pageRoutes maps URL paths to their corresponding page data.
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

	// pageTemplateFS embeds the HTML template file for rendering pages.
	//go:embed template.html
	pageTemplateFS embed.FS

	// pageTemplate is the parsed HTML template used for rendering pages.
	pageTemplate = template.Must(template.ParseFS(pageTemplateFS, "template.html"))

	// staticFileServer serves static files from the "./static" directory.
	staticFileServer = http.FileServer(http.Dir("./static"))
)

// Run starts the web server and listens for incoming requests.
func Run() {
	router := http.NewServeMux()
	router.HandleFunc("/", rootHandler)
	listenAndServe(env.ServerAddress(), router)
}

// rootHandler handles all incoming requests
// and routes them to the appropriate handler.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is a GET request.
	// If not, reply with an error 405.
	if r.Method != http.MethodGet {
		log.Printf("[error] Method not allowed: %s", r.Method)
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Check if the URL path corresponds to a defined page.
	p, ok := pageRoutes[r.URL.Path]
	if ok {
		renderPage(w, p)
		return
	}
	// Check if the URL path corresponds to a static file.
	// If not, reply with an error 404.
	serveStaticFile(w, r)
}

// renderPage executes the HTML template with the provided page data
// and writes the result to the response writer.
// If there is an error executing the template, it replies with an error 500.
func renderPage(w http.ResponseWriter, p *pageData) {
	err := pageTemplate.Execute(w, p)
	if err != nil {
		log.Printf("[error] Failed to render page %s: %v", p.WASMPath, err)
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
	}
}

// serveStaticFile writes the contents of a static file to the response writer.
// If the file is not found, it replies with an error 404.
func serveStaticFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=86400")
	staticFileServer.ServeHTTP(w, r)
}

// listenAndServe starts the web server and listens for incoming requests.
// If there is an error starting the server, it exits the application.
func listenAndServe(addr string, handler http.Handler) {
	log.Printf("[info] Server is running on %s", addr)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatalf("[fatal] Failed to run server: %v", err)
	}
}
