package main

import (
	"goSim/internal/middleware"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	loggerMux := middleware.Middleware(mux)
	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", loggerMux); err != nil {
		log.Fatalf("Server falied %v", err)
	}
}
