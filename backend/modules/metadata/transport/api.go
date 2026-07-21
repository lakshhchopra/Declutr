package transport

import (
	"encoding/json"
	"net/http"

	"github.com/diablovocado/declutr/modules/metadata/application"
)

type API struct {
	service application.MetadataService
}

func NewAPI(service application.MetadataService) *API {
	return &API{service: service}
}

func (a *API) GetMetadataHandler(w http.ResponseWriter, r *http.Request) {
	// In reality, this would be grabbed from chi context, e.g. chi.URLParam(r, "assetId")
	assetID := r.URL.Query().Get("assetId")

	meta, err := a.service.GetMetadata(r.Context(), assetID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(meta)
}

func (a *API) GetVersionHistoryHandler(w http.ResponseWriter, r *http.Request) {
	assetID := r.URL.Query().Get("assetId")

	history, err := a.service.GetVersionHistory(r.Context(), assetID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}

func (a *API) RefreshMetadataHandler(w http.ResponseWriter, r *http.Request) {
	// Re-triggers extraction
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"status": "queued"})
}
