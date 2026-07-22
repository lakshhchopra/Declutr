package transport

import (
	"encoding/json"
	"net/http"

	"github.com/diablovocado/declutr/modules/predictive/application"
	"github.com/diablovocado/declutr/modules/predictive/domain"
)

type PredictiveAPI struct {
	service *application.PredictiveService
}

func NewPredictiveAPI(service *application.PredictiveService) *PredictiveAPI {
	return &PredictiveAPI{service: service}
}

func (a *PredictiveAPI) GetPredictions(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		userID = "usr-default"
	}

	preds, err := a.service.GenerateAndGetPredictions(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(preds)
}

func (a *PredictiveAPI) AcceptPrediction(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID       string `json:"user_id"`
		PredictionID string `json:"prediction_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.UserID == "" {
		req.UserID = "usr-default"
	}

	if err := a.service.AcceptPrediction(r.Context(), req.UserID, req.PredictionID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *PredictiveAPI) DismissPrediction(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID       string `json:"user_id"`
		PredictionID string `json:"prediction_id"`
		Reason       string `json:"reason"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.UserID == "" {
		req.UserID = "usr-default"
	}

	if err := a.service.DismissPrediction(r.Context(), req.UserID, req.PredictionID, req.Reason); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *PredictiveAPI) ManageSettings(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		userID = "usr-default"
	}

	if r.Method == http.MethodPut {
		var settings domain.PredictionSettings
		if err := json.NewDecoder(r.Body).Decode(&settings); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		settings.UserID = userID
		if err := a.service.UpdateSettings(r.Context(), &settings); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	settings, err := a.service.GetSettings(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(settings)
}

func (a *PredictiveAPI) GetStats(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		userID = "usr-default"
	}

	stats, err := a.service.GetStats(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(stats)
}

func (a *PredictiveAPI) GetHistory(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		userID = "usr-default"
	}

	history, err := a.service.GetHistory(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(history)
}
