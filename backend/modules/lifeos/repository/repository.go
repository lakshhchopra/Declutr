package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/diablovocado/declutr/modules/lifeos/domain"
)

type LifeOSRepository interface {
	CreateLifeArea(ctx context.Context, area *domain.LifeArea) error
	ListLifeAreas(ctx context.Context, userID string) ([]domain.LifeArea, error)

	CreateProject(ctx context.Context, prj *domain.Project) error
	ListProjects(ctx context.Context, userID string) ([]domain.Project, error)
	GetProjectByID(ctx context.Context, id string) (*domain.Project, error)

	CreateGoal(ctx context.Context, goal *domain.ProjectGoal) error
	ListGoals(ctx context.Context, userID string, projectID string) ([]domain.ProjectGoal, error)
	UpdateGoalProgress(ctx context.Context, id string, progressPct int, isCompleted bool) error

	ListTimelineEvents(ctx context.Context, userID string) ([]domain.LifeTimelineEvent, error)
	GetLifeMetrics(ctx context.Context, userID string) (*domain.LifeMetric, error)
}

type InMemoryLifeOSRepository struct {
	mu        sync.RWMutex
	areas     map[string]*domain.LifeArea
	projects  map[string]*domain.Project
	goals     map[string]*domain.ProjectGoal
	timeline  []domain.LifeTimelineEvent
	metrics   map[string]*domain.LifeMetric
}

func NewInMemoryLifeOSRepository() *InMemoryLifeOSRepository {
	repo := &InMemoryLifeOSRepository{
		areas:    make(map[string]*domain.LifeArea),
		projects: make(map[string]*domain.Project),
		goals:    make(map[string]*domain.ProjectGoal),
		timeline: []domain.LifeTimelineEvent{},
		metrics:  make(map[string]*domain.LifeMetric),
	}

	now := time.Now().UTC()

	// Seed 12 Default Life Areas
	defaultAreas := []string{
		domain.AreaPersonal, domain.AreaWork, domain.AreaBusiness, domain.AreaEducation,
		domain.AreaFinance, domain.AreaHealth, domain.AreaTravel, domain.AreaLegal,
		domain.AreaHome, domain.AreaFamily, domain.AreaResearch, domain.AreaHobbies,
	}

	for _, name := range defaultAreas {
		id := "area-" + name
		repo.areas[id] = &domain.LifeArea{
			ID:          id,
			UserID:      "usr-default",
			Name:        name,
			Description: fmt.Sprintf("%s Life Area", name),
			Icon:        name,
			Color:       "#6366F1",
			IsCustom:    false,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
	}

	// Seed Sample First-Class Projects
	p1 := &domain.Project{
		ID:             "prj-startup",
		UserID:         "usr-default",
		LifeAreaID:     "area-Business",
		Name:           "Launch Startup",
		Description:    "Building Declutr LifeOS Platform",
		Status:         domain.ProjectInProgress,
		Budget:         50000,
		AssociatedDocs: []string{"ast-pitch-deck-2026", "ast-cap-table"},
		People:         []string{"Alex Vance", "Sarah Chen"},
		TargetDate:     now.Add(90 * 24 * time.Hour),
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	p2 := &domain.Project{
		ID:             "prj-japan",
		UserID:         "usr-default",
		LifeAreaID:     "area-Travel",
		Name:           "Japan Vacation Trip",
		Description:    "Autumn vacation in Tokyo & Kyoto",
		Status:         domain.ProjectInProgress,
		Budget:         4500,
		AssociatedDocs: []string{"ast-flight-ticket-ba117"},
		People:         []string{"Family"},
		TargetDate:     now.Add(45 * 24 * time.Hour),
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	repo.projects[p1.ID] = p1
	repo.projects[p2.ID] = p2

	// Seed Sample Goals
	g1 := &domain.ProjectGoal{
		ID:            "gol-passport",
		ProjectID:     "prj-japan",
		UserID:        "usr-default",
		Title:         "Renew Passport & Check Visa Buffer",
		Description:   "Ensure passport has 6 months validity before departure",
		ProgressPct:   60,
		IsCompleted:   false,
		MissingAssets: []string{"ast-renewed-passport"},
		DueDate:       now.Add(14 * 24 * time.Hour),
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	repo.goals[g1.ID] = g1

	// Seed Timeline Events
	repo.timeline = append(repo.timeline, domain.LifeTimelineEvent{
		ID:          "tl-1",
		UserID:      "usr-default",
		ProjectID:   "prj-japan",
		Title:       "Flight Ticket Booked",
		EventType:   "TRIP",
		Description: "BA117 London to Tokyo confirmed",
		Timestamp:   now.Add(-2 * 24 * time.Hour),
	})

	return repo
}

func (r *InMemoryLifeOSRepository) CreateLifeArea(ctx context.Context, area *domain.LifeArea) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.areas[area.ID] = area
	return nil
}

func (r *InMemoryLifeOSRepository) ListLifeAreas(ctx context.Context, userID string) ([]domain.LifeArea, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var list []domain.LifeArea
	for _, a := range r.areas {
		if a.UserID == userID || userID == "" {
			list = append(list, *a)
		}
	}
	return list, nil
}

func (r *InMemoryLifeOSRepository) CreateProject(ctx context.Context, prj *domain.Project) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.projects[prj.ID] = prj
	return nil
}

func (r *InMemoryLifeOSRepository) ListProjects(ctx context.Context, userID string) ([]domain.Project, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var list []domain.Project
	for _, p := range r.projects {
		if p.UserID == userID || userID == "" {
			list = append(list, *p)
		}
	}
	return list, nil
}

func (r *InMemoryLifeOSRepository) GetProjectByID(ctx context.Context, id string) (*domain.Project, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.projects[id]
	if !ok {
		return nil, fmt.Errorf("project not found: %s", id)
	}
	return p, nil
}

func (r *InMemoryLifeOSRepository) CreateGoal(ctx context.Context, goal *domain.ProjectGoal) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.goals[goal.ID] = goal
	return nil
}

func (r *InMemoryLifeOSRepository) ListGoals(ctx context.Context, userID string, projectID string) ([]domain.ProjectGoal, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var list []domain.ProjectGoal
	for _, g := range r.goals {
		if (g.UserID == userID || userID == "") && (g.ProjectID == projectID || projectID == "") {
			list = append(list, *g)
		}
	}
	return list, nil
}

func (r *InMemoryLifeOSRepository) UpdateGoalProgress(ctx context.Context, id string, progressPct int, isCompleted bool) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	g, ok := r.goals[id]
	if !ok {
		return fmt.Errorf("goal not found: %s", id)
	}
	g.ProgressPct = progressPct
	g.IsCompleted = isCompleted
	g.UpdatedAt = time.Now().UTC()
	return nil
}

func (r *InMemoryLifeOSRepository) ListTimelineEvents(ctx context.Context, userID string) ([]domain.LifeTimelineEvent, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.timeline, nil
}

func (r *InMemoryLifeOSRepository) GetLifeMetrics(ctx context.Context, userID string) (*domain.LifeMetric, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return &domain.LifeMetric{
		UserID:         userID,
		ActiveProjects: len(r.projects),
		GoalCompletion: 0.85,
		LifeBalance:    92,
	}, nil
}
