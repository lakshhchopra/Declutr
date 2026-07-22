package transport

import (
	"encoding/json"
	"net/http"

	"github.com/diablovocado/declutr/modules/multiagent/application"
	"github.com/diablovocado/declutr/modules/multiagent/domain"
)

type MultiAgentAPI struct {
	service *application.MultiAgentService
}

func NewMultiAgentAPI(service *application.MultiAgentService) *MultiAgentAPI {
	return &MultiAgentAPI{service: service}
}

func (a *MultiAgentAPI) ProcessGoal(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	graph, consensus, err := a.service.ProcessUserGoal(r.Context(), req.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"task_graph": graph,
		"consensus":  consensus,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}

func (a *MultiAgentAPI) ListAgents(w http.ResponseWriter, r *http.Request) {
	agents, err := a.service.ListAgents(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(agents)
}

func (a *AgentAPI) RegisterAgent(w http.ResponseWriter, r *http.Request) {
	var req domain.AgentRegistration
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.service.RegisterAgent(r.Context(), &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *MultiAgentAPI) GetMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := a.service.GetMessageAuditLogs(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(messages)
}

func (a *MultiAgentAPI) GetTaskGraph(w http.ResponseWriter, r *http.Request) {
	goalID := r.URL.Query().Get("goal_id")
	graph, err := a.service.GetTaskGraph(r.Context(), goalID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(graph)
}

func (a *MultiAgentAPI) GetHealth(w http.ResponseWriter, r *http.Request) {
	health, err := a.service.ListHealthMetrics(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(health)
}
