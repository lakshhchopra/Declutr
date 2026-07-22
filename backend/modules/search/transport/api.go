package transport

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/diablovocado/declutr/modules/search/application"
	"github.com/diablovocado/declutr/modules/search/domain"
)

// SearchAPI handles HTTP requests for the Hybrid Knowledge Search Engine
type SearchAPI struct {
	service *application.SearchService
}

// NewSearchAPI creates a new SearchAPI instance
func NewSearchAPI(service *application.SearchService) *SearchAPI {
	return &SearchAPI{service: service}
}

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func errJSON(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, map[string]string{"error": msg})
}

// ExecuteSearch executes a hybrid search query
// POST /api/v1/search/query
func (a *SearchAPI) ExecuteSearch(w http.ResponseWriter, r *http.Request) {
	var req domain.SearchQueryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.VaultID == "" {
		errJSON(w, http.StatusBadRequest, "invalid request body or missing vaultId")
		return
	}

	resp, err := a.service.ExecuteSearch(r.Context(), &req)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

// SaveSearch bookmarks a search query with filters
// POST /api/v1/search/saved
func (a *SearchAPI) SaveSearch(w http.ResponseWriter, r *http.Request) {
	var ss domain.SavedSearch
	if err := json.NewDecoder(r.Body).Decode(&ss); err != nil || ss.VaultID == "" {
		errJSON(w, http.StatusBadRequest, "invalid request body or missing vaultId")
		return
	}
	if err := a.service.SaveSearch(&ss); err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "search saved", "savedId": ss.SavedID})
}

// GetSavedSearches lists saved searches for a vault
// GET /api/v1/search/saved?vaultId=<id>
func (a *SearchAPI) GetSavedSearches(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	if vaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	saved, err := a.service.GetSavedSearches(vaultID)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"savedSearches": saved,
		"total":         len(saved),
	})
}

// DeleteSavedSearch deletes a saved search
// DELETE /api/v1/search/saved?savedId=<id>
func (a *SearchAPI) DeleteSavedSearch(w http.ResponseWriter, r *http.Request) {
	savedID := r.URL.Query().Get("savedId")
	if savedID == "" {
		errJSON(w, http.StatusBadRequest, "savedId is required")
		return
	}
	if err := a.service.DeleteSavedSearch(savedID); err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "saved search deleted", "savedId": savedID})
}

// GetHistory returns recent search query history
// GET /api/v1/search/history?vaultId=<id>&limit=<n>
func (a *SearchAPI) GetHistory(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	if vaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	limitStr := r.URL.Query().Get("limit")
	limit := 10
	if limitStr != "" {
		if parsed, err := strconv.Atoi(limitStr); err == nil {
			limit = parsed
		}
	}
	history, err := a.service.GetHistory(vaultID, limit)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"history": history,
		"total":   len(history),
	})
}

// GetSuggestions provides search-as-you-type autocomplete suggestions
// GET /api/v1/search/suggestions?vaultId=<id>&q=<query>
func (a *SearchAPI) GetSuggestions(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	q := r.URL.Query().Get("q")
	if vaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}

	var suggestions []string
	if strings.TrimSpace(q) != "" {
		suggestions = append(suggestions, q+" in Japan Vacation")
		suggestions = append(suggestions, q+" PDF files")
		suggestions = append(suggestions, q+" Dr. Sharma medical records")
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"suggestions": suggestions,
		"query":       q,
	})
}

// GetStats returns vault search engine statistics
// GET /api/v1/search/stats?vaultId=<id>
func (a *SearchAPI) GetStats(w http.ResponseWriter, r *http.Request) {
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

// GetPreferences returns ranking weights and search preferences
// GET /api/v1/search/preferences?vaultId=<id>
func (a *SearchAPI) GetPreferences(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	if vaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	prefs, err := a.service.GetPreferences(vaultID)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, prefs)
}

// UpdatePreferences updates ranking weights and search defaults
// PUT /api/v1/search/preferences
func (a *SearchAPI) UpdatePreferences(w http.ResponseWriter, r *http.Request) {
	var prefs domain.SearchPreferences
	if err := json.NewDecoder(r.Body).Decode(&prefs); err != nil || prefs.VaultID == "" {
		errJSON(w, http.StatusBadRequest, "invalid request body or missing vaultId")
		return
	}
	if err := a.service.UpdatePreferences(&prefs); err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "search preferences updated"})
}
