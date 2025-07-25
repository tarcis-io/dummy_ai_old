package server

import (
	"log"
	"net/http"
	"text/template"

	"dummy_ai/internal/env"
)

type (
	Server struct {
		router       *http.ServeMux
		httpTemplate *template.Template
	}
)

func (s *Server) RegisterFileServer(urlPath, dir string) {
	s.router.Handle(urlPath, http.StripPrefix(urlPath, http.FileServer(http.Dir(dir))))
}

func (s *Server) ListenAndServe() {
	err := http.ListenAndServe(env.ServerAddress(), s.router)
	if err != nil {
		log.Fatalf("FATAL: Failed to start server: %v", err)
	}
}

func New() *Server {
	return &Server{
		router:       http.NewServeMux(),
		httpTemplate: template.Must(template.ParseFiles("template.html")),
	}
}
