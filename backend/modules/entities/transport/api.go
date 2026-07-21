package transport

import (
	"encoding/json"
	"net/http"

	"github.com/diablovocado/declutr/modules/entities/application"
)

type API struct {
	service application.EntityService
}

func NewAPI(service application.EntityService) *API {
	return &API{service: service}
}

func (a *API) GetEntitiesByVaultHandler(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")

	entities, err := a.service.GetEntitiesByVault(r.Context(), vaultID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entities)
}

func (a *API) GetAssetEntitiesHandler(w http.ResponseWriter, r *http.Request) {
	assetID := r.URL.Query().Get("assetId")

	occurrences, err := a.service.GetOccurrencesByAsset(r.Context(), assetID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(occurrences)
}

func (a *API) RefreshEntitiesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"status": "queued"})
}
