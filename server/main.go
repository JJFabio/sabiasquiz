package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	// 1. Servir archivos estáticos (Frontend)
	// Si existe la carpeta 'dist' (build de React), la servimos. Si no, servimos un index simple.
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// 2. API Endpoint
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := Response{
			Message: "¡Hola desde el API de SabiasQuiz! 🧠✨",
			Status:  "success",
		}
		json.NewEncoder(w).Encode(response)
	})

	// 3. Health Check (Importante para Traefik)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 SabiasQuiz Server API running on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
