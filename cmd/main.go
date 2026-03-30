package main

import (
	"log"
	"net/http"

	"portfolio/internal/handlers"
	"portfolio/internal/rag"

	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// load knowledge base
	if err := rag.LoadKnowledge("knowledge"); err != nil {
		log.Fatalf("Failed to load knowledge base: %v", err)
	}
	log.Println("Knowledge base loaded successfully")

	// create router
	router := http.NewServeMux()

	// create server
	server := http.FileServer(http.Dir("./static"))
	router.Handle("/", server)

	// set up endpoints
	router.HandleFunc("/api/health", handlers.Health)
	router.HandleFunc("/api/chat", handlers.Chat)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(router)))
}
