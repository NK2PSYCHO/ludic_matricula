package main

import (
	"log"
	"net/http"

	apphttp "ludic-matricula/internal/infrastructure/http"
)

func main() {
	router := apphttp.NewRouter()

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
