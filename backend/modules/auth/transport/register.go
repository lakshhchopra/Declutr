package transport

import (
	"encoding/json"
	"net/http"

	"github.com/diablovocado/declutr/modules/auth/application"
	"github.com/diablovocado/declutr/modules/auth/transport/models"
)

func RegisterHandler(service *application.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.RegisterRequest

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

		resp := models.RegisterResponse{
			UserID: userID,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
