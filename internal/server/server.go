package server

import (
	"net/http"
	"text/template"
)

type (
	server struct {
		router       *http.ServeMux
		httpTemplate *template.Template
	}
)
