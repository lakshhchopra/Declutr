package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/diablovocado/declutr/modules/multiagent/domain"
)

type MultiAgentRepository interface {
	RegisterAgent(ctx context.Context, agent *domain.AgentRegistration) error
	GetAgentByID(ctx context.Context, id string) (*domain.AgentRegistration, error)
	ListAgents(ctx context.Context) ([]domain.AgentRegistration, error)

	SaveTaskGraph(ctx context.Context, graph *domain.TaskGraph) error
	GetTaskGraph(ctx context.Context, goalID string) (*domain.TaskGraph, error)

	AddSharedMemory(ctx context.Context, item *domain.SharedMemoryItem) error
	ListSharedMemory(ctx context.Context, goalID string) ([]domain.SharedMemoryItem, error)

	RecordHealth(ctx context.Context, metric *domain.AgentHealthMetric) error
	ListHealthMetrics(ctx context.Context) ([]domain.AgentHealthMetric, error)
}

type InMemoryMultiAgentRepository struct {
	mu           sync.RWMutex
	agents       map[string]*domain.AgentRegistration
	graphs       map[string]*domain.TaskGraph
	sharedMemory map[string]*domain.SharedMemoryItem
	health       map[string]*domain.AgentHealthMetric
}

func NewInMemoryMultiAgentRepository() *InMemoryMultiAgentRepository {
	repo := &InMemoryMultiAgentRepository{
		agents:       make(map[string]*domain.AgentRegistration),
		graphs:       make(map[string]*domain.TaskGraph),
		sharedMemory: make(map[string]*domain.SharedMemoryItem),
		health:       make(map[string]*domain.AgentHealthMetric),
	}

	// Seed Specialist Agents
	seedRoles := []domain.AgentRole{
		domain.RoleCoordinator,
		domain.RoleKnowledge,
		domain.RoleMemory,
		domain.RoleResearch,
		domain.RoleOrganization,
		domain.RoleWorkflow,
		domain.RoleSearch,
		domain.RoleSecurity,
		domain.RoleIntegration,
		domain.RoleTimeline,
	}

	for idx, r := range seedRoles {
		id := fmt.Sprintf("agt-%s", r)
		reg := &domain.AgentRegistration{
			ID:             id,
			Role:           r,
			Name:           fmt.Sprintf("%s Specialist", r),
			Capabilities:   []string{string(r) + "_CAPABILITY"},
			SupportedTools: []string{"SEARCH", "WORKFLOW", "MEMORY"},
			Status:         "READY",
			Version:        "2.0.0",
			Health:         "HEALTHY",
			Priority:       idx + 1,
			CreatedAt:      time.Now().UTC(),
			UpdatedAt:      time.Now().UTC(),
		}
		repo.agents[id] = reg

		repo.health[id] = &domain.AgentHealthMetric{
			AgentID:     id,
			Role:        r,
			Status:      "HEALTHY",
			LatencyMs:   12,
			SuccessRate: 0.99,
			TotalTasks:  150,
			FailedTasks: 1,
		}
	}

	return repo
}

func (r *InMemoryMultiAgentRepository) RegisterAgent(ctx context.Context, agent *domain.AgentRegistration) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.agents[agent.ID] = agent
	return nil
}

func (r *InMemoryMultiAgentRepository) GetAgentByID(ctx context.Context, id string) (*domain.AgentRegistration, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	a, ok := r.agents[id]
	if !ok {
		return nil, fmt.Errorf("agent registration not found: %s", id)
	}
	return a, nil
}

func (r *InMemoryMultiAgentRepository) ListAgents(ctx context.Context) ([]domain.AgentRegistration, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []domain.AgentRegistration
	for _, a := range r.agents {
		result = append(result, *a)
	}
	return result, nil
}

func (r *InMemoryMultiAgentRepository) SaveTaskGraph(ctx context.Context, graph *domain.TaskGraph) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.graphs[graph.GoalID] = graph
	return nil
}

func (r *InMemoryMultiAgentRepository) GetTaskGraph(ctx context.Context, goalID string) (*domain.TaskGraph, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	g, ok := r.graphs[goalID]
	if !ok {
		return nil, fmt.Errorf("task graph not found for goal: %s", goalID)
	}
	return g, nil
}

func (r *InMemoryMultiAgentRepository) AddSharedMemory(ctx context.Context, item *domain.SharedMemoryItem) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.sharedMemory[item.ID] = item
	return nil
}

func (r *InMemoryMultiAgentRepository) ListSharedMemory(ctx context.Context, goalID string) ([]domain.SharedMemoryItem, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []domain.SharedMemoryItem
	for _, m := range r.sharedMemory {
		if m.GoalID == goalID || goalID == "" {
			result = append(result, *m)
		}
	}
	return result, nil
}

func (r *InMemoryMultiAgentRepository) RecordHealth(ctx context.Context, metric *domain.AgentHealthMetric) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.health[metric.AgentID] = metric
	return nil
}

func (r *InMemoryMultiAgentRepository) ListHealthMetrics(ctx context.Context) ([]domain.AgentHealthMetric, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []domain.AgentHealthMetric
	for _, h := range r.health {
		result = append(result, *h)
	}
	return result, nil
}
