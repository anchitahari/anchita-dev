package main

import (
	"log"
	"net/http"

	"portfolio/internal/handlers"
)

func main() {
	// Create router
	router := http.NewServeMux()

	// Create server
	server := http.FileServer(http.Dir("./static"))
	router.Handle("/", server)

	// Set up endpoints
	router.HandleFunc("/api/health", handlers.Health)
	router.HandleFunc("/api/chat", handlers.Chat)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(router)))
}
