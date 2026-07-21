package transport

import (
	"encoding/json"
	"net/http"

	"github.com/diablovocado/declutr/modules/ai/application"
)

type API struct {
	service application.AIAnalysisService
}

func NewAPI(service application.AIAnalysisService) *API {
	return &API{service: service}
}

func (a *API) GetAnalysisHandler(w http.ResponseWriter, r *http.Request) {
	assetID := r.URL.Query().Get("assetId")

	analysis, err := a.service.GetAnalysis(r.Context(), assetID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if analysis == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(analysis)
}

func (a *API) GetVersionHistoryHandler(w http.ResponseWriter, r *http.Request) {
	analysisID := r.URL.Query().Get("analysisId")

	history, err := a.service.GetVersionHistory(r.Context(), analysisID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}

func (a *API) RefreshAnalysisHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"status": "queued"})
}
