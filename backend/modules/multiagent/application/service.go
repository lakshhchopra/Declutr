package application

import (
	"context"
	"time"

	"github.com/diablovocado/declutr/modules/multiagent/domain"
	"github.com/diablovocado/declutr/modules/multiagent/repository"
	"github.com/diablovocado/declutr/shared/observability"
)

type MultiAgentService struct {
	repo        repository.MultiAgentRepository
	bus         *MessageBus
	planner     *MultiAgentTaskPlanner
	coordinator *CoordinatorAgent
}

func NewMultiAgentService(repo repository.MultiAgentRepository) *MultiAgentService {
	bus := NewMessageBus()
	planner := NewMultiAgentTaskPlanner()
	coordinator := NewCoordinatorAgent(bus, planner)

	return &MultiAgentService{
		repo:        repo,
		bus:         bus,
		planner:     planner,
		coordinator: coordinator,
	}
}

func (s *MultiAgentService) ProcessUserGoal(ctx context.Context, goalTitle string) (*domain.TaskGraph, *domain.ConsensusResult, error) {
	goalID := "gol-multi-" + observability.GenerateID(8)

	graph, consensus, err := s.coordinator.OrchestrateGoal(ctx, goalID, goalTitle)
	if err != nil {
		return nil, nil, err
	}

	_ = s.repo.SaveTaskGraph(ctx, graph)

	// Save Output in Shared Memory
	_ = s.repo.AddSharedMemory(ctx, &domain.SharedMemoryItem{
		ID:        "shm-" + observability.GenerateID(8),
		GoalID:    goalID,
		AgentID:   "agt-coordinator",
		Key:       "final_consensus_result",
		Value:     consensus.Explanation,
		Category:  "TASK_RESULT",
		CreatedAt: time.Now().UTC(),
	})

	return graph, consensus, nil
}

func (s *MultiAgentService) ListAgents(ctx context.Context) ([]domain.AgentRegistration, error) {
	return s.repo.ListAgents(ctx)
}

func (s *MultiAgentService) RegisterAgent(ctx context.Context, reg *domain.AgentRegistration) error {
	if reg.ID == "" {
		reg.ID = "agt-" + string(reg.Role)
	}
	reg.CreatedAt = time.Now().UTC()
	reg.UpdatedAt = time.Now().UTC()
	return s.repo.RegisterAgent(ctx, reg)
}

func (s *MultiAgentService) GetMessageAuditLogs(ctx context.Context) ([]*domain.AgentMessage, error) {
	return s.bus.GetAuditHistory(), nil
}

func (s *MultiAgentService) GetTaskGraph(ctx context.Context, goalID string) (*domain.TaskGraph, error) {
	return s.repo.GetTaskGraph(ctx, goalID)
}

func (s *MultiAgentService) ListHealthMetrics(ctx context.Context) ([]domain.AgentHealthMetric, error) {
	return s.repo.ListHealthMetrics(ctx)
}
