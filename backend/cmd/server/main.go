package main

import (
	"log"
	"net/http"

	"github.com/diablovocado/declutr/internal/auth"
	"github.com/diablovocado/declutr/internal/auth/srp"
	"github.com/diablovocado/declutr/internal/database"
	"github.com/diablovocado/declutr/internal/health"
	"github.com/diablovocado/declutr/internal/middleware"
	"github.com/diablovocado/declutr/internal/repository"
)

func main() {
	db := database.Connect()

	userRepo := &repository.PostgresUserRepository{
		DB: db,
	}

	authService := &auth.Service{
		UserRepo:   userRepo,
		Challenges: srp.NewChallengeStore(),
		SRP:        srp.NewEngine(),
	}

	http.HandleFunc("/health", health.Handler)

	http.HandleFunc(
		"/api/v1/auth/register",
		auth.RegisterHandler(authService),
	)

	http.HandleFunc(
		"/api/v1/auth/login/start",
		auth.LoginStartHandler(authService),
	)
	http.Handle(
		"/api/v1/me",
		middleware.Auth(userRepo)(auth.MeHandler()),
	)

	http.HandleFunc(
		"/api/v1/auth/login/finish",
		auth.LoginFinishHandler(authService),
	)

	log.Println("Declutr Backend Running on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
