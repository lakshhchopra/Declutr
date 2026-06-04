package main

import (
	"log"
	"net/http"

	"github.com/diablovocado/declutr/internal/auth"
	"github.com/diablovocado/declutr/internal/health"
)

func main() {
	http.HandleFunc("/health", health.Handler)
	http.HandleFunc("/api/v1/auth/register", auth.RegisterHandler)

	log.Println("Declutr Backend Running on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
