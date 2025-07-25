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

func (s *Server) ListenAndServe() {
	err := http.ListenAndServe(env.ServerAddress(), s.router)
	if err != nil {
		log.Fatalf("FATAL: Failed to start server: %v", err)
	}
}

func New() *Server {
	router := http.NewServeMux()
	return &Server{
		router:       router,
		httpTemplate: template.Must(template.ParseFiles("template.html")),
	}
}
