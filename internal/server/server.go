package server

import (
	"log"
	"net/http"
	"text/template"

	"dummy_ai/internal/env"
)

var (
	httpTemplate = template.Must(template.ParseFiles("template.html"))
)

func Run() {
	err := http.ListenAndServe(env.ServerAddress(), nil)
	if err != nil {
		log.Fatalf("FATAL: Failed to start server: %v", err)
	}
}
