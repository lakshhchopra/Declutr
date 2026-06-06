package auth

import (
	"encoding/json"
	"net/http"

	authmodels "github.com/diablovocado/declutr/internal/auth/models"
)

func RegisterHandler(service *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req authmodels.RegisterRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		if req.Email == "" ||
			req.SRPVerifier == "" ||
			req.SRPSalt == "" ||
			req.MVK.Ciphertext == "" {
			http.Error(w, "missing required fields", http.StatusBadRequest)
			return
		}

		userID, err := service.Register(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := authmodels.RegisterResponse{
			UserID: userID,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
