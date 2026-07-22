package transport

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/diablovocado/declutr/modules/lifeos/application"
)

type LifeOSAPI struct {
	service *application.LifeOSService
}

func NewLifeOSAPI(service *application.LifeOSService) *LifeOSAPI {
	return &LifeOSAPI{service: service}
}

func (a *LifeOSAPI) GetDashboard(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		userID = "usr-default"
	}

	dash, err := a.service.GetUnifiedDashboard(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(dash)
}

func (a *LifeOSAPI) ListLifeAreas(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		userID = "usr-default"
	}

	areas, err := a.service.ListLifeAreas(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(areas)
}

func (a *LifeOSAPI) ManageProjects(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		userID = "usr-default"
	}

	if r.Method == http.MethodPost {
		var req struct {
			LifeAreaID  string    `json:"life_area_id"`
			Name        string    `json:"name"`
			Description string    `json:"description"`
			TargetDate  time.Time `json:"target_date"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		prj, err := a.service.CreateProject(r.Context(), userID, req.LifeAreaID, req.Name, req.Description, req.TargetDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(prj)
		return
	}

	projects, err := a.service.ListProjects(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(projects)
}

func (a *LifeOSAPI) ManageGoals(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		userID = "usr-default"
	}
	projectID := r.URL.Query().Get("project_id")

	if r.Method == http.MethodPost {
		var req struct {
			ProjectID   string    `json:"project_id"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			DueDate     time.Time `json:"due_date"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		goal, err := a.service.CreateGoal(r.Context(), userID, req.ProjectID, req.Title, req.Description, req.DueDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(goal)
		return
	}

	goals, err := a.service.ListGoals(r.Context(), userID, projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(goals)
}

func (a *LifeOSAPI) UpdateGoalProgress(w http.ResponseWriter, r *http.Request) {
	var req struct {
		GoalID      string `json:"goal_id"`
		ProgressPct int    `json:"progress_pct"`
		IsCompleted bool   `json:"is_completed"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.service.UpdateGoalProgress(r.Context(), req.GoalID, req.ProgressPct, req.IsCompleted); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *LifeOSAPI) GetTimeline(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		userID = "usr-default"
	}

	events, err := a.service.GetTimeline(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(events)
}

func (a *LifeOSAPI) GetMetrics(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		userID = "usr-default"
	}

	metrics, err := a.service.GetMetrics(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(metrics)
}
