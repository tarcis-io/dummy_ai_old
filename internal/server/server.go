package server

import (
	"net/http"
	"text/template"
)

type (
	Server struct {
		router       *http.ServeMux
		httpTemplate *template.Template
	}
)

func New() *Server {
	router := http.NewServeMux()
	return &Server{
		router:       router,
		httpTemplate: template.Must(template.ParseFiles("template.html")),
	}
}
