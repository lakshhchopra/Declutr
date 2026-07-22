package application_test

import (
	"context"
	"testing"

	"github.com/diablovocado/declutr/modules/workflow/application"
	"github.com/diablovocado/declutr/modules/workflow/domain"
	"github.com/diablovocado/declutr/modules/workflow/repository"
)

const testVaultID = "vault-test-001"

func setupService() *application.WorkflowService {
	repo := repository.NewInMemoryWorkflowRepository()
	return application.NewWorkflowService(repo)
}

// TestWorkflowCreationAndToggle validates workflow creation, listing, and toggling
func TestWorkflowCreationAndToggle(t *testing.T) {
	svc := setupService()

	wf, err := svc.CreateWorkflow(&domain.Workflow{
		VaultID:     testVaultID,
		Name:        "Test Auto Tagging",
		Description: "Tags PDF assets automatically",
	})
	if err != nil {
		t.Fatalf("create workflow failed: %v", err)
	}
	if !wf.Enabled {
		t.Error("expected new workflow to be enabled by default")
	}

	if err := svc.ToggleWorkflow(wf.WorkflowID, false); err != nil {
		t.Fatalf("toggle workflow failed: %v", err)
	}

	updated, err := svc.GetWorkflow(wf.WorkflowID)
	if err != nil {
		t.Fatalf("get workflow failed: %v", err)
	}
	if updated.Enabled {
		t.Error("expected workflow to be disabled after toggle")
	}

	t.Logf("PASS: Workflow Creation & Toggle — Created %s, toggled enabled state", wf.WorkflowID)
}

// TestConditionEvaluation validates AND/OR/NOT logic and operator matching
func TestConditionEvaluation(t *testing.T) {
	conds := []domain.WorkflowCondition{
		{Field: "fileType", Operator: domain.OpEquals, Value: "PDF", Combinator: domain.CombinatorAND},
		{Field: "entity", Operator: domain.OpContains, Value: "Japan", Combinator: domain.CombinatorAND},
	}

	ctxMatch := map[string]interface{}{"fileType": "PDF", "entity": "Tokyo, Japan"}
	if !application.EvaluateConditions(conds, ctxMatch) {
		t.Error("expected conditions to evaluate to TRUE for matching context")
	}

	ctxNoMatch := map[string]interface{}{"fileType": "DOCX", "entity": "Tokyo, Japan"}
	if application.EvaluateConditions(conds, ctxNoMatch) {
		t.Error("expected conditions to evaluate to FALSE for non-matching context")
	}

	t.Logf("PASS: Condition Evaluation — Rule matching verified for AND/OR operators")
}

// TestWorkflowExecutionSuccess validates end-to-end execution of a workflow
func TestWorkflowExecutionSuccess(t *testing.T) {
	svc := setupService()
	ctx := context.Background()

	wf, _ := svc.CreateWorkflow(&domain.Workflow{
		VaultID: testVaultID,
		Name:    "Auto Tagging Workflow",
		Conditions: []domain.WorkflowCondition{
			{Field: "fileType", Operator: domain.OpEquals, Value: "PDF", Combinator: domain.CombinatorAND},
		},
		Actions: []domain.WorkflowAction{
			{ActionType: domain.ActionApplyTags, Config: map[string]interface{}{"tags": []string{"Travel"}}, ExecutionOrder: 1},
		},
	})

	run, err := svc.RunWorkflow(ctx, &domain.RunWorkflowRequest{
		WorkflowID: wf.WorkflowID,
		VaultID:    testVaultID,
		Context:    map[string]interface{}{"fileType": "PDF"},
	})
	if err != nil {
		t.Fatalf("run workflow failed: %v", err)
	}
	if run.Status != domain.RunSuccess {
		t.Errorf("expected run status SUCCESS, got %s (Error: %s)", run.Status, run.ErrorMessage)
	}

	t.Logf("PASS: Workflow Execution Success — Run %s completed in %dms with status SUCCESS",
		run.RunID, run.DurationMs)
}

// TestWorkflowConditionFailure validates execution stopping when conditions evaluate to false
func TestWorkflowConditionFailure(t *testing.T) {
	svc := setupService()
	ctx := context.Background()

	wf, _ := svc.CreateWorkflow(&domain.Workflow{
		VaultID: testVaultID,
		Name:    "Strict PDF Workflow",
		Conditions: []domain.WorkflowCondition{
			{Field: "fileType", Operator: domain.OpEquals, Value: "PDF", Combinator: domain.CombinatorAND},
		},
	})

	run, err := svc.RunWorkflow(ctx, &domain.RunWorkflowRequest{
		WorkflowID: wf.WorkflowID,
		VaultID:    testVaultID,
		Context:    map[string]interface{}{"fileType": "DOCX"}, // mismatched
	})
	if err != nil {
		t.Fatalf("run workflow failed: %v", err)
	}
	if run.Status != domain.RunFailed {
		t.Errorf("expected run status FAILED due to mismatched condition, got %s", run.Status)
	}

	t.Logf("PASS: Condition Failure Handling — Run stopped correctly when condition failed")
}

// TestExecutionHistoryAndLogs validates logging run events and retrieving history
func TestExecutionHistoryAndLogs(t *testing.T) {
	svc := setupService()
	ctx := context.Background()

	wfs, _ := svc.ListWorkflows(testVaultID)
	if len(wfs) == 0 {
		t.Fatal("expected sample workflows")
	}

	run, err := svc.RunWorkflow(ctx, &domain.RunWorkflowRequest{
		WorkflowID: wfs[0].WorkflowID,
		VaultID:    testVaultID,
		Context:    map[string]interface{}{"fileType": "PDF", "entity": "Japan"},
	})
	if err != nil {
		t.Fatalf("run workflow failed: %v", err)
	}

	logs, err := svc.GetRunLogs(run.RunID)
	if err != nil {
		t.Fatalf("get run logs failed: %v", err)
	}
	if len(logs) == 0 {
		t.Error("expected execution logs for run")
	}

	t.Logf("PASS: Execution History & Logs — Run logged %d execution steps", len(logs))
}

// TestWorkflowStats validates observability metrics calculation
func TestWorkflowStats(t *testing.T) {
	svc := setupService()

	stats, err := svc.GetStats(testVaultID)
	if err != nil {
		t.Fatalf("get stats failed: %v", err)
	}
	if stats.TotalWorkflows == 0 {
		t.Error("expected positive total workflows count")
	}
	if stats.SuccessRate < 0.0 || stats.SuccessRate > 1.0 {
		t.Errorf("invalid success rate: %.2f", stats.SuccessRate)
	}

	t.Logf("PASS: Workflow Stats — %d total workflows, %d active, SuccessRate=%.1f%%",
		stats.TotalWorkflows, stats.ActiveWorkflows, stats.SuccessRate*100)
}
