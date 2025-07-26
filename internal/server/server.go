package server

import (
	"log"
	"net/http"

	"dummy_ai/internal/env"
)

func Run() {
	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	serverAddress := env.ServerAddress()
	log.Printf("INFO: Server starting on %s", serverAddress)
	err := http.ListenAndServe(serverAddress, router)
	if err != nil {
		log.Fatalf("FATAL: Failed to start server: %v", err)
	}
}
