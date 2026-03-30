package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	anthropicClient "portfolio/internal/anthropic"
	"portfolio/internal/rag"
)

// CORS middleware — allows frontend to call the API
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Health — Kubernetes liveness probe
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Message string `json:"message"`
}

func Chat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Message) == "" {
		http.Error(w, "message required", http.StatusBadRequest)
		return
	}

	// retrieve relevant chunks
	chunks := rag.Retrieve(req.Message, 4)

	// build context from chunks
	var contextBuilder strings.Builder
	for _, chunk := range chunks {
		contextBuilder.WriteString(chunk.Text)
		contextBuilder.WriteString("\n\n")
	}
	context := contextBuilder.String()

	// build system prompt
	system := `You are an AI assistant on Anchita Hari Narayanan's personal portfolio website. 
	Your job is to answer questions about Anchita based only on the information provided to you.
	Be warm, conversational, and concise. Speak about Anchita in the third person.
	Do not use markdown formatting, bullet points, or asterisks in your responses. Plain text only.
	If you don't know something or it's not in the context, say so honestly rather than making things up.
	Keep responses to 2-4 sentences unless more detail is genuinely needed.

	Here is what you know about Anchita:` + context

	// call anthropic
	response, err := anthropicClient.Complete(system, req.Message)
	if err != nil {
		log.Printf("Anthropic error: %v", err)
		http.Error(w, "failed to generate response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ChatResponse{Message: response})
}
