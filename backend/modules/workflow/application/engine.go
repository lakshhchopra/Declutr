package application

import (
	"context"
	"fmt"
	"log"

	"github.com/diablovocado/declutr/modules/workflow/domain"
)

// WorkflowExecutionEngine orchestrates asynchronous trigger handling, rule evaluation, and action execution
type WorkflowExecutionEngine struct {
	service *WorkflowService
}

// NewWorkflowExecutionEngine creates a new WorkflowExecutionEngine
func NewWorkflowExecutionEngine(service *WorkflowService) *WorkflowExecutionEngine {
	return &WorkflowExecutionEngine{service: service}
}

// DispatchEvent receives internal vault events (AssetUploaded, DocumentExpiring, MemoryCreated) and runs matching workflows
func (e *WorkflowExecutionEngine) DispatchEvent(ctx context.Context, vaultID string, eventType domain.TriggerType, eventCtx map[string]interface{}) error {
	log.Printf("[WorkflowExecutionEngine] Dispatching event: %s for vault: %s", eventType, vaultID)

	workflows, err := e.service.ListWorkflows(vaultID)
	if err != nil {
		return fmt.Errorf("workflow execution engine: list workflows failed: %w", err)
	}

	for _, wf := range workflows {
		if !wf.Enabled {
			continue
		}

		for _, trig := range wf.Triggers {
			if trig.TriggerType == eventType || trig.TriggerType == domain.TriggerManual {
				log.Printf("[WorkflowExecutionEngine] Match found! Executing workflow: %s (%s)", wf.Name, wf.WorkflowID)
				run, err := e.service.RunWorkflow(ctx, &domain.RunWorkflowRequest{
					WorkflowID: wf.WorkflowID,
					VaultID:    vaultID,
					Context:    eventCtx,
				})
				if err != nil {
					log.Printf("[WorkflowExecutionEngine] Run error for %s: %v", wf.WorkflowID, err)
				} else {
					log.Printf("[WorkflowExecutionEngine] Run completed: Status=%s, Duration=%dms", run.Status, run.DurationMs)
				}
				break
			}
		}
	}

	return nil
}
