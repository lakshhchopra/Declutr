package repository

import (
	"fmt"
	"sync"
	"time"

	"github.com/diablovocado/declutr/modules/workflow/domain"
)

// WorkflowRepository defines persistence contract for workflows, runs, and execution logs
type WorkflowRepository interface {
	CreateWorkflow(wf *domain.Workflow) error
	GetWorkflow(wfID string) (*domain.Workflow, error)
	ListWorkflows(vaultID string) ([]*domain.Workflow, error)
	UpdateWorkflow(wf *domain.Workflow) error
	DeleteWorkflow(wfID string) error
	ToggleWorkflow(wfID string, enabled bool) error

	AddRun(run *domain.WorkflowRun) error
	GetRuns(wfID string) ([]*domain.WorkflowRun, error)
	ListAllRuns(vaultID string) ([]*domain.WorkflowRun, error)

	AddLog(log *domain.WorkflowLog) error
	GetLogs(runID string) ([]*domain.WorkflowLog, error)

	GetStats(vaultID string) (*domain.WorkflowStats, error)
	ClearAllData(vaultID string) error
}

// InMemoryWorkflowRepository is a thread-safe in-memory store
type InMemoryWorkflowRepository struct {
	mu        sync.RWMutex
	workflows map[string]*domain.Workflow   // wfID -> wf
	runs      map[string][]*domain.WorkflowRun // wfID -> runs
	logs      map[string][]*domain.WorkflowLog // runID -> logs
}

// NewInMemoryWorkflowRepository creates a new in-memory workflow repository
func NewInMemoryWorkflowRepository() *InMemoryWorkflowRepository {
	return &InMemoryWorkflowRepository{
		workflows: make(map[string]*domain.Workflow),
		runs:      make(map[string][]*domain.WorkflowRun),
		logs:      make(map[string][]*domain.WorkflowLog),
	}
}

func (r *InMemoryWorkflowRepository) CreateWorkflow(wf *domain.Workflow) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	wf.UpdatedAt = time.Now()
	r.workflows[wf.WorkflowID] = wf
	return nil
}

func (r *InMemoryWorkflowRepository) GetWorkflow(wfID string) (*domain.Workflow, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	wf, ok := r.workflows[wfID]
	if !ok {
		return nil, fmt.Errorf("workflow %s not found", wfID)
	}
	return wf, nil
}

func (r *InMemoryWorkflowRepository) ListWorkflows(vaultID string) ([]*domain.Workflow, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var list []*domain.Workflow
	for _, wf := range r.workflows {
		if wf.VaultID == vaultID {
			list = append(list, wf)
		}
	}
	if len(list) == 0 {
		list = defaultSampleWorkflows(vaultID)
		for _, wf := range list {
			r.workflows[wf.WorkflowID] = wf
		}
	}
	return list, nil
}

func (r *InMemoryWorkflowRepository) UpdateWorkflow(wf *domain.Workflow) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	wf.UpdatedAt = time.Now()
	r.workflows[wf.WorkflowID] = wf
	return nil
}

func (r *InMemoryWorkflowRepository) DeleteWorkflow(wfID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.workflows, wfID)
	delete(r.runs, wfID)
	return nil
}

func (r *InMemoryWorkflowRepository) ToggleWorkflow(wfID string, enabled bool) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	wf, ok := r.workflows[wfID]
	if !ok {
		return fmt.Errorf("workflow %s not found", wfID)
	}
	wf.Enabled = enabled
	wf.UpdatedAt = time.Now()
	return nil
}

func (r *InMemoryWorkflowRepository) AddRun(run *domain.WorkflowRun) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.runs[run.WorkflowID] = append(r.runs[run.WorkflowID], run)

	if wf, ok := r.workflows[run.WorkflowID]; ok {
		wf.RunCount++
		if run.Status == domain.RunFailed {
			wf.FailureCount++
		}
		now := time.Now()
		wf.LastRunAt = &now
		wf.UpdatedAt = now
	}
	return nil
}

func (r *InMemoryWorkflowRepository) GetRuns(wfID string) ([]*domain.WorkflowRun, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.runs[wfID], nil
}

func (r *InMemoryWorkflowRepository) ListAllRuns(vaultID string) ([]*domain.WorkflowRun, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var all []*domain.WorkflowRun
	for wfID, runs := range r.runs {
		if wf, ok := r.workflows[wfID]; ok && wf.VaultID == vaultID {
			all = append(all, runs...)
		}
	}
	if len(all) == 0 {
		return defaultSampleRuns(vaultID), nil
	}
	return all, nil
}

func (r *InMemoryWorkflowRepository) AddLog(log *domain.WorkflowLog) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.logs[log.RunID] = append(r.logs[log.RunID], log)
	return nil
}

func (r *InMemoryWorkflowRepository) GetLogs(runID string) ([]*domain.WorkflowLog, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.logs[runID], nil
}

func (r *InMemoryWorkflowRepository) GetStats(vaultID string) (*domain.WorkflowStats, error) {
	wfs, _ := r.ListWorkflows(vaultID)
	runs, _ := r.ListAllRuns(vaultID)

	active := 0
	for _, wf := range wfs {
		if wf.Enabled {
			active++
		}
	}

	succ := 0
	fail := 0
	var totalDur int64
	for _, run := range runs {
		if run.Status == domain.RunSuccess {
			succ++
		} else if run.Status == domain.RunFailed {
			fail++
		}
		totalDur += run.DurationMs
	}

	rate := 1.0
	if len(runs) > 0 {
		rate = float64(succ) / float64(len(runs))
	}

	avgDur := 0.0
	if len(runs) > 0 {
		avgDur = float64(totalDur) / float64(len(runs))
	}

	return &domain.WorkflowStats{
		VaultID:         vaultID,
		TotalWorkflows:  len(wfs),
		ActiveWorkflows: active,
		TotalRuns:       len(runs),
		SuccessfulRuns: succ,
		FailedRuns:      fail,
		SuccessRate:     rate,
		AvgDurationMs:   avgDur,
	}, nil
}

func (r *InMemoryWorkflowRepository) ClearAllData(vaultID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for id, wf := range r.workflows {
		if wf.VaultID == vaultID {
			delete(r.workflows, id)
			delete(r.runs, id)
		}
	}
	return nil
}

// Sample Data Generators
func defaultSampleWorkflows(vaultID string) []*domain.Workflow {
	now := time.Now()
	t1 := now.Add(-2 * time.Hour)
	return []*domain.Workflow{
		{
			WorkflowID:  "wf-travel-001",
			VaultID:     vaultID,
			Name:        "Auto-tag Travel & Passport Documents",
			Description: "Automatically tags PDF documents matching Japan or Passport entities and moves them into the Japan Vacation collection.",
			Enabled:     true,
			Status:      "IDLE",
			Triggers: []domain.WorkflowTrigger{
				{TriggerID: "trig-1", WorkflowID: "wf-travel-001", TriggerType: domain.TriggerAssetUploaded, CreatedAt: now},
			},
			Conditions: []domain.WorkflowCondition{
				{ConditionID: "cond-1", WorkflowID: "wf-travel-001", Field: "fileType", Operator: domain.OpEquals, Value: "PDF", Combinator: domain.CombinatorAND},
				{ConditionID: "cond-2", WorkflowID: "wf-travel-001", Field: "entity", Operator: domain.OpContains, Value: "Japan", Combinator: domain.CombinatorAND},
			},
			Actions: []domain.WorkflowAction{
				{ActionID: "act-1", WorkflowID: "wf-travel-001", ActionType: domain.ActionApplyTags, Config: map[string]interface{}{"tags": []string{"Travel", "Passport"}}, ExecutionOrder: 1, CreatedAt: now},
				{ActionID: "act-2", WorkflowID: "wf-travel-001", ActionType: domain.ActionNotifyUser, Config: map[string]interface{}{"message": "Passport document auto-tagged and organized."}, ExecutionOrder: 2, CreatedAt: now},
			},
			LastRunAt:     &t1,
			RunCount:      12,
			FailureCount:  0,
			AvgDurationMs: 45,
			CreatedBy:     "USER",
			CreatedAt:     now.Add(-7 * 24 * time.Hour),
			UpdatedAt:     now,
		},
		{
			WorkflowID:  "wf-expiry-002",
			VaultID:     vaultID,
			Name:        "Document Expiration Alert & Timeline Sync",
			Description: "Fires weekly to scan for documents expiring within 60 days and creates milestone alerts.",
			Enabled:     true,
			Status:      "IDLE",
			Triggers: []domain.WorkflowTrigger{
				{TriggerID: "trig-2", WorkflowID: "wf-expiry-002", TriggerType: domain.TriggerDocExpiring, CreatedAt: now},
			},
			Conditions: []domain.WorkflowCondition{
				{ConditionID: "cond-3", WorkflowID: "wf-expiry-002", Field: "confidence", Operator: domain.OpGreaterThan, Value: "0.8", Combinator: domain.CombinatorAND},
			},
			Actions: []domain.WorkflowAction{
				{ActionID: "act-3", WorkflowID: "wf-expiry-002", ActionType: domain.ActionCreateReminder, Config: map[string]interface{}{"title": "Passport Renewal Reminder"}, ExecutionOrder: 1, CreatedAt: now},
			},
			LastRunAt:     &t1,
			RunCount:      5,
			FailureCount:  0,
			AvgDurationMs: 30,
			CreatedBy:     "SYSTEM",
			CreatedAt:     now.Add(-14 * 24 * time.Hour),
			UpdatedAt:     now,
		},
	}
}

func defaultSampleRuns(vaultID string) []*domain.WorkflowRun {
	now := time.Now()
	tCompleted := now.Add(-1 * time.Hour)
	return []*domain.WorkflowRun{
		{
			RunID:        "run-001",
			WorkflowID:   "wf-travel-001",
			VaultID:      vaultID,
			TriggerEvent: "ASSET_UPLOADED",
			Status:       domain.RunSuccess,
			DurationMs:   42,
			StartedAt:    now.Add(-1 * time.Hour),
			CompletedAt:  &tCompleted,
		},
	}
}
