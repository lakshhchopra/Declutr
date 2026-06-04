package auth

import (
	"encoding/json"
	"net/http"

	authmodels "github.com/diablovocado/declutr/internal/auth/models"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req authmodels.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	resp := authmodels.RegisterResponse{
		UserID: "user_test_123",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
