package application

import (
	"context"

	"github.com/diablovocado/declutr/modules/multiagent/domain"
	"github.com/diablovocado/declutr/shared/observability"
)

type MultiAgentTaskPlanner struct{}

func NewMultiAgentTaskPlanner() *MultiAgentTaskPlanner {
	return &MultiAgentTaskPlanner{}
}

func (tp *MultiAgentTaskPlanner) BuildTaskGraph(ctx context.Context, goalID string, goalTitle string) (*domain.TaskGraph, error) {
	graph := &domain.TaskGraph{
		GoalID:    goalID,
		GoalTitle: goalTitle,
		Status:    "PENDING",
		Tasks:     []domain.CoordinatorTask{},
	}

	task1ID := "tsk-" + observability.GenerateID(6)
	task2ID := "tsk-" + observability.GenerateID(6)
	task3ID := "tsk-" + observability.GenerateID(6)
	task4ID := "tsk-" + observability.GenerateID(6)

	// Step 1: Parallel Research & Search tasks
	task1 := domain.CoordinatorTask{
		ID:            task1ID,
		GoalID:        goalID,
		AssignedRole:  domain.RoleSearch,
		Action:        "execute_hybrid_search",
		Parameters:    map[string]interface{}{"query": goalTitle},
		Dependencies:  []string{},
		ExecutionMode: domain.ExecParallel,
		Status:        "PENDING",
		Confidence:    0.95,
	}

	task2 := domain.CoordinatorTask{
		ID:            task2ID,
		GoalID:        goalID,
		AssignedRole:  domain.RoleKnowledge,
		Action:        "query_knowledge_graph",
		Parameters:    map[string]interface{}{"entity": goalTitle},
		Dependencies:  []string{},
		ExecutionMode: domain.ExecParallel,
		Status:        "PENDING",
		Confidence:    0.92,
	}

	// Step 2: Sequential Organization task dependent on Search & Knowledge
	task3 := domain.CoordinatorTask{
		ID:            task3ID,
		GoalID:        goalID,
		AssignedRole:  domain.RoleOrganization,
		Action:        "structure_and_classify",
		Parameters:    map[string]interface{}{"target": goalTitle},
		Dependencies:  []string{task1ID, task2ID},
		ExecutionMode: domain.ExecSequential,
		Status:        "PENDING",
		Confidence:    0.96,
	}

	// Step 3: Security & Consensus Review task
	task4 := domain.CoordinatorTask{
		ID:            task4ID,
		GoalID:        goalID,
		AssignedRole:  domain.RoleSecurity,
		Action:        "verify_permissions_and_consensus",
		Parameters:    map[string]interface{}{"rule": "PRIVACY_FIRST"},
		Dependencies:  []string{task3ID},
		ExecutionMode: domain.ExecSequential,
		Status:        "PENDING",
		Confidence:    0.99,
	}

	graph.Tasks = []domain.CoordinatorTask{task1, task2, task3, task4}
	return graph, nil
}
