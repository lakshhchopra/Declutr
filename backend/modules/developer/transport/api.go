package transport

import (
	"encoding/json"
	"net/http"

	"github.com/diablovocado/declutr/modules/developer/application"
	"github.com/diablovocado/declutr/modules/developer/domain"
)

// DeveloperAPI handles HTTP REST endpoints for public API keys, webhooks, and OAuth.
type DeveloperAPI struct {
	service *application.DeveloperService
}

func NewDeveloperAPI(service *application.DeveloperService) *DeveloperAPI {
	return &DeveloperAPI{service: service}
}

func (a *DeveloperAPI) CreateAPIKey(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateAPIKeyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		userID = "usr-dev-default"
	}

	key, rawSecret, err := a.service.GenerateAPIKey(r.Context(), userID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"api_key":    key,
		"raw_secret": rawSecret, // Secret returned only ONCE
	})
}

func (a *DeveloperAPI) ListAPIKeys(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		userID = "usr-dev-default"
	}

	keys, err := a.service.ListAPIKeys(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(keys)
}

func (a *DeveloperAPI) RevokeAPIKey(w http.ResponseWriter, r *http.Request) {
	keyID := r.URL.Query().Get("id")
	if keyID == "" {
		http.Error(w, "missing key id", http.StatusBadRequest)
		return
	}

	if err := a.service.RevokeAPIKey(r.Context(), keyID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *DeveloperAPI) RegisterWebhook(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateWebhookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		userID = "usr-dev-default"
	}

	hook, err := a.service.RegisterWebhook(r.Context(), userID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(hook)
}

func (a *DeveloperAPI) ListWebhooks(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		userID = "usr-dev-default"
	}

	hooks, err := a.service.ListWebhooks(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(hooks)
}

func (a *DeveloperAPI) GetWebhookDeliveries(w http.ResponseWriter, r *http.Request) {
	webhookID := r.URL.Query().Get("webhook_id")

	deliveries, err := a.service.GetDeliveries(r.Context(), webhookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dlq, _ := a.service.GetDLQ(r.Context())

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"deliveries": deliveries,
		"dlq_items":  dlq,
	})
}

func (a *DeveloperAPI) CreateOAuthApp(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name         string   `json:"name"`
		RedirectURIs []string `json:"redirect_uris"`
		Scopes       []string `json:"scopes"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		userID = "usr-dev-default"
	}

	client, clientSecret, err := a.service.CreateOAuthApp(r.Context(), userID, req.Name, req.RedirectURIs, req.Scopes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"client":        client,
		"client_secret": clientSecret,
	})
}

func (a *DeveloperAPI) ExchangeOAuthToken(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		GrantType    string `json:"grant_type"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := a.service.ExchangeOAuthToken(r.Context(), req.ClientID, req.ClientSecret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(token)
}

func (a *DeveloperAPI) ServeOpenAPISpec(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	spec := map[string]interface{}{
		"openapi": "3.0.3",
		"info": map[string]interface{}{
			"title":       "Declutr Developer Platform Public API",
			"version":     "v1.0.0",
			"description": "Production REST API for Declutr Vaults, Assets, Search, AI Copilot, Workflows, Notifications, and Webhooks.",
		},
		"paths": map[string]interface{}{
			"/api/v1/vaults": map[string]interface{}{
				"get": map[string]interface{}{
					"summary": "List User & Organization Vaults",
					"tags":    []string{"Vaults"},
				},
			},
			"/api/v1/search/query": map[string]interface{}{
				"post": map[string]interface{}{
					"summary": "Execute Hybrid Vector & Keyword Search",
					"tags":    []string{"Search"},
				},
			},
			"/api/v1/copilot/messages": map[string]interface{}{
				"post": map[string]interface{}{
					"summary": "Send RAG Grounded AI Copilot Prompt",
					"tags":    []string{"AI Copilot"},
				},
			},
		},
	}
	_ = json.NewEncoder(w).Encode(spec)
}
