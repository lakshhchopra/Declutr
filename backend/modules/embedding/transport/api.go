package transport

import (
	"encoding/json"
	"net/http"

	"github.com/diablovocado/declutr/modules/embedding/application"
	"github.com/diablovocado/declutr/modules/embedding/domain"
)

// EmbeddingAPI handles HTTP requests for the Embedding Engine
type EmbeddingAPI struct {
	service *application.EmbeddingService
}

// NewEmbeddingAPI creates a new EmbeddingAPI instance
func NewEmbeddingAPI(service *application.EmbeddingService) *EmbeddingAPI {
	return &EmbeddingAPI{service: service}
}

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func errJSON(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, map[string]string{"error": msg})
}

// GenerateEmbedding creates a vector embedding for a structured representation
// POST /api/v1/embedding/generate
func (a *EmbeddingAPI) GenerateEmbedding(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Input domain.StructuredRepresentationInput `json:"input"`
		Opts  *domain.GenerationOptions            `json:"options,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Input.VaultID == "" {
		errJSON(w, http.StatusBadRequest, "invalid request body or missing vaultId")
		return
	}

	emb, err := a.service.GenerateEmbedding(r.Context(), &body.Input, body.Opts)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, emb)
}

// RefreshEmbeddings triggers an incremental vector refresh for a vault
// POST /api/v1/embedding/refresh
func (a *EmbeddingAPI) RefreshEmbeddings(w http.ResponseWriter, r *http.Request) {
	var body struct {
		VaultID string `json:"vaultId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.VaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	if err := a.service.RefreshEmbeddings(r.Context(), body.VaultID); err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "embedding refresh triggered", "vaultId": body.VaultID})
}

// GetStatus returns the operational status of the embedding pipeline
// GET /api/v1/embedding/status?vaultId=<id>
func (a *EmbeddingAPI) GetStatus(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	if vaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	status, err := a.service.GetStatus(vaultID)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, status)
}

// GetStats returns vault-level embedding statistics
// GET /api/v1/embedding/stats?vaultId=<id>
func (a *EmbeddingAPI) GetStats(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	if vaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	stats, err := a.service.GetStats(vaultID)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, stats)
}

// GetHistory returns generation job history and version upgrades
// GET /api/v1/embedding/history?vaultId=<id>
func (a *EmbeddingAPI) GetHistory(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	if vaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	history, err := a.service.GetHistory(vaultID)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, history)
}

// UpdateProvider updates the active provider configuration for a vault
// PUT /api/v1/embedding/provider
func (a *EmbeddingAPI) UpdateProvider(w http.ResponseWriter, r *http.Request) {
	var cfg domain.EmbeddingProviderConfig
	if err := json.NewDecoder(r.Body).Decode(&cfg); err != nil || cfg.VaultID == "" {
		errJSON(w, http.StatusBadRequest, "invalid request body or missing vaultId")
		return
	}
	if err := a.service.UpdateProviderConfig(&cfg); err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "provider configuration updated"})
}

// RebuildVersion upgrades embedding model/version and re-embeds items
// POST /api/v1/embedding/rebuild
func (a *EmbeddingAPI) RebuildVersion(w http.ResponseWriter, r *http.Request) {
	var body struct {
		VaultID      string              `json:"vaultId"`
		ProviderName domain.ProviderName `json:"providerName"`
		ModelName    string              `json:"modelName"`
		VersionTag   string              `json:"versionTag"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.VaultID == "" {
		errJSON(w, http.StatusBadRequest, "invalid request body or missing vaultId")
		return
	}
	ver, err := a.service.RebuildForVersion(r.Context(), body.VaultID, body.ProviderName, body.ModelName, body.VersionTag)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, ver)
}
