package main

import (
	"log"
	"net/http"

	"github.com/diablovocado/declutr/internal/auth"
	"github.com/diablovocado/declutr/internal/database"
	"github.com/diablovocado/declutr/internal/health"
	"github.com/diablovocado/declutr/internal/repository"
)

func main() {
	db := database.Connect()

	userRepo := &repository.PostgresUserRepository{
		DB: db,
	}

	authService := &auth.Service{
		UserRepo: userRepo,
	}

	http.HandleFunc("/health", health.Handler)
	http.HandleFunc("/api/v1/auth/register", auth.RegisterHandler(authService))

	log.Println("Declutr Backend Running on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
