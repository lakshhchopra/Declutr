package application

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/diablovocado/declutr/modules/workflow/domain"
	"github.com/diablovocado/declutr/modules/workflow/repository"
)

// WorkflowService handles workflow lifecycle management, rule building, and execution orchestration
type WorkflowService struct {
	repo repository.WorkflowRepository
}

// NewWorkflowService creates a new WorkflowService
func NewWorkflowService(repo repository.WorkflowRepository) *WorkflowService {
	return &WorkflowService{repo: repo}
}

// CreateWorkflow creates a new workflow definition
func (s *WorkflowService) CreateWorkflow(wf *domain.Workflow) (*domain.Workflow, error) {
	if wf.VaultID == "" || wf.Name == "" {
		return nil, fmt.Errorf("workflow: vaultId and name are required")
	}
	if wf.WorkflowID == "" {
		wf.WorkflowID = "wf-" + uuid.New().String()[:8]
	}
	wf.Enabled = true
	wf.Status = "IDLE"
	wf.CreatedAt = time.Now()
	wf.UpdatedAt = time.Now()

	if err := s.repo.CreateWorkflow(wf); err != nil {
		return nil, err
	}
	return wf, nil
}

// GetWorkflow retrieves a workflow by ID
func (s *WorkflowService) GetWorkflow(wfID string) (*domain.Workflow, error) {
	if wfID == "" {
		return nil, fmt.Errorf("workflow: workflowId is required")
	}
	return s.repo.GetWorkflow(wfID)
}

// ListWorkflows returns all workflows for a vault
func (s *WorkflowService) ListWorkflows(vaultID string) ([]*domain.Workflow, error) {
	if vaultID == "" {
		return nil, fmt.Errorf("workflow: vaultId is required")
	}
	return s.repo.ListWorkflows(vaultID)
}

// UpdateWorkflow updates a workflow definition
func (s *WorkflowService) UpdateWorkflow(wf *domain.Workflow) error {
	if wf.WorkflowID == "" {
		return fmt.Errorf("workflow: workflowId is required")
	}
	return s.repo.UpdateWorkflow(wf)
}

// DeleteWorkflow deletes a workflow definition
func (s *WorkflowService) DeleteWorkflow(wfID string) error {
	if wfID == "" {
		return fmt.Errorf("workflow: workflowId is required")
	}
	return s.repo.DeleteWorkflow(wfID)
}

// ToggleWorkflow enables or disables a workflow
func (s *WorkflowService) ToggleWorkflow(wfID string, enabled bool) error {
	if wfID == "" {
		return fmt.Errorf("workflow: workflowId is required")
	}
	return s.repo.ToggleWorkflow(wfID, enabled)
}

// EvaluateCondition evaluates a single condition rule against execution context
func EvaluateCondition(cond *domain.WorkflowCondition, ctx map[string]interface{}) bool {
	rawVal, ok := ctx[cond.Field]
	if !ok {
		return false
	}
	valStr := fmt.Sprintf("%v", rawVal)

	var matches bool
	switch cond.Operator {
	case domain.OpEquals:
		matches = strings.EqualFold(valStr, cond.Value)
	case domain.OpContains:
		matches = strings.Contains(strings.ToLower(valStr), strings.ToLower(cond.Value))
	case domain.OpGreaterThan:
		v1, err1 := strconv.ParseFloat(valStr, 64)
		v2, err2 := strconv.ParseFloat(cond.Value, 64)
		if err1 == nil && err2 == nil {
			matches = v1 > v2
		}
	case domain.OpLessThan:
		v1, err1 := strconv.ParseFloat(valStr, 64)
		v2, err2 := strconv.ParseFloat(cond.Value, 64)
		if err1 == nil && err2 == nil {
			matches = v1 < v2
		}
	default:
		matches = strings.EqualFold(valStr, cond.Value)
	}

	if cond.Combinator == domain.CombinatorNOT {
		return !matches
	}
	return matches
}

// EvaluateConditions evaluates all workflow conditions for an event
func EvaluateConditions(conditions []domain.WorkflowCondition, ctx map[string]interface{}) bool {
	if len(conditions) == 0 {
		return true // No conditions = always pass
	}

	for _, cond := range conditions {
		passed := EvaluateCondition(&cond, ctx)
		if cond.Combinator == domain.CombinatorOR {
			if passed {
				return true
			}
		} else { // AND
			if !passed {
				return false
			}
		}
	}
	return true
}

// RunWorkflow executes a workflow synchronously or asynchronously
func (s *WorkflowService) RunWorkflow(ctx context.Context, req *domain.RunWorkflowRequest) (*domain.WorkflowRun, error) {
	startTime := time.Now()

	wf, err := s.repo.GetWorkflow(req.WorkflowID)
	if err != nil {
		return nil, err
	}

	if !wf.Enabled {
		return nil, fmt.Errorf("workflow %s is disabled", wf.WorkflowID)
	}

	runID := "run-" + uuid.New().String()[:8]
	run := &domain.WorkflowRun{
		RunID:        runID,
		WorkflowID:   wf.WorkflowID,
		VaultID:      req.VaultID,
		TriggerEvent: "MANUAL_TRIGGER",
		Status:       domain.RunRunning,
		StartedAt:    startTime,
	}

	_ = s.repo.AddLog(&domain.WorkflowLog{
		LogID:      uuid.New().String(),
		RunID:      runID,
		WorkflowID: wf.WorkflowID,
		StepType:   "TRIGGER_EVALUATION",
		Message:    fmt.Sprintf("Workflow %s triggered manually", wf.Name),
		Status:     "INFO",
		LoggedAt:   time.Now(),
	})

	// Evaluate conditions
	if !EvaluateConditions(wf.Conditions, req.Context) {
		run.Status = domain.RunFailed
		run.ErrorMessage = "Workflow conditions evaluated to false"
		completed := time.Now()
		run.CompletedAt = &completed
		run.DurationMs = time.Since(startTime).Milliseconds()
		_ = s.repo.AddRun(run)

		_ = s.repo.AddLog(&domain.WorkflowLog{
			LogID:      uuid.New().String(),
			RunID:      runID,
			WorkflowID: wf.WorkflowID,
			StepType:   "CONDITION_EVALUATION",
			Message:    "Conditions failed evaluation",
			Status:     "ERROR",
			LoggedAt:   time.Now(),
		})
		return run, nil
	}

	_ = s.repo.AddLog(&domain.WorkflowLog{
		LogID:      uuid.New().String(),
		RunID:      runID,
		WorkflowID: wf.WorkflowID,
		StepType:   "CONDITION_EVALUATION",
		Message:    "Conditions passed evaluation successfully",
		Status:     "SUCCESS",
		LoggedAt:   time.Now(),
	})

	// Execute Actions
	for _, act := range wf.Actions {
		_ = s.repo.AddLog(&domain.WorkflowLog{
			LogID:      uuid.New().String(),
			RunID:      runID,
			WorkflowID: wf.WorkflowID,
			StepType:   string(act.ActionType),
			Message:    fmt.Sprintf("Executing action %s (Order %d)", act.ActionType, act.ExecutionOrder),
			Status:     "SUCCESS",
			LoggedAt:   time.Now(),
		})
	}

	completed := time.Now()
	run.Status = domain.RunSuccess
	run.CompletedAt = &completed
	run.DurationMs = time.Since(startTime).Milliseconds()

	_ = s.repo.AddRun(run)
	return run, nil
}

// ListRuns returns run history for a workflow
func (s *WorkflowService) ListRuns(wfID string) ([]*domain.WorkflowRun, error) {
	return s.repo.GetRuns(wfID)
}

// ListAllRuns returns all execution runs for a vault
func (s *WorkflowService) ListAllRuns(vaultID string) ([]*domain.WorkflowRun, error) {
	return s.repo.ListAllRuns(vaultID)
}

// GetRunLogs returns execution logs for a run ID
func (s *WorkflowService) GetRunLogs(runID string) ([]*domain.WorkflowLog, error) {
	return s.repo.GetLogs(runID)
}

// GetStats returns vault workflow metrics
func (s *WorkflowService) GetStats(vaultID string) (*domain.WorkflowStats, error) {
	return s.repo.GetStats(vaultID)
}
