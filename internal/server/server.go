package server

import (
	"embed"
	"log"
	"net/http"
	"text/template"

	"dummy_ai/internal/env"
)

type (
	templateData struct {
		Title    string
		WASMPath string
	}
)

var (
	pageRoutes = map[string]*templateData{
		"/": {
			Title:    "DummyAI",
			WASMPath: "/wasm/home.wasm",
		},
		"/about": {
			Title:    "DummyAI",
			WASMPath: "/wasm/about.wasm",
		},
	}

	//go:embed template.html
	htmlTemplateFS embed.FS
	htmlTemplate   = template.Must(template.ParseFS(htmlTemplateFS, "template.html"))

	staticFileServer = http.FileServer(http.Dir("./static"))
)

func Run() {
	router := http.NewServeMux()
	router.HandleFunc("/", handleRequest)
	listenAndServe(env.ServerAddress(), router)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
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

func renderPage(w http.ResponseWriter, t *templateData) {
	err := htmlTemplate.Execute(w, t)
	if err != nil {
		log.Printf("[error] Failed to render page %s: %v", t.WASMPath, err)
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
	}
}

func serveStaticFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=86400")
	staticFileServer.ServeHTTP(w, r)
}

func listenAndServe(addr string, handler http.Handler) {
	log.Printf("[info] Server is running on %s", addr)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatalf("[fatal] Failed to run server: %v", err)
	}
}
