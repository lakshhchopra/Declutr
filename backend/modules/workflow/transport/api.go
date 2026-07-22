package transport

import (
	"encoding/json"
	"net/http"

	"github.com/diablovocado/declutr/modules/workflow/application"
	"github.com/diablovocado/declutr/modules/workflow/domain"
)

// WorkflowAPI handles HTTP endpoints for the Workflow Automation Engine
type WorkflowAPI struct {
	service *application.WorkflowService
}

// NewWorkflowAPI creates a new WorkflowAPI instance
func NewWorkflowAPI(service *application.WorkflowService) *WorkflowAPI {
	return &WorkflowAPI{service: service}
}

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func errJSON(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, map[string]string{"error": msg})
}

// CreateWorkflow handles creating a new workflow definition
// POST /api/v1/workflows
func (a *WorkflowAPI) CreateWorkflow(w http.ResponseWriter, r *http.Request) {
	var wf domain.Workflow
	if err := json.NewDecoder(r.Body).Decode(&wf); err != nil || wf.VaultID == "" || wf.Name == "" {
		errJSON(w, http.StatusBadRequest, "invalid request body, missing vaultId or name")
		return
	}
	created, err := a.service.CreateWorkflow(&wf)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, created)
}

// ListWorkflows returns workflows for a vault
// GET /api/v1/workflows?vaultId=<id>
func (a *WorkflowAPI) ListWorkflows(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	if vaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	list, err := a.service.ListWorkflows(vaultID)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"workflows": list,
		"total":     len(list),
	})
}

// UpdateWorkflow updates a workflow definition
// PUT /api/v1/workflows
func (a *WorkflowAPI) UpdateWorkflow(w http.ResponseWriter, r *http.Request) {
	var wf domain.Workflow
	if err := json.NewDecoder(r.Body).Decode(&wf); err != nil || wf.WorkflowID == "" {
		errJSON(w, http.StatusBadRequest, "invalid request body, missing workflowId")
		return
	}
	if err := a.service.UpdateWorkflow(&wf); err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "workflow updated", "workflowId": wf.WorkflowID})
}

// DeleteWorkflow deletes a workflow definition
// DELETE /api/v1/workflows?workflowId=<id>
func (a *WorkflowAPI) DeleteWorkflow(w http.ResponseWriter, r *http.Request) {
	wfID := r.URL.Query().Get("workflowId")
	if wfID == "" {
		errJSON(w, http.StatusBadRequest, "workflowId is required")
		return
	}
	if err := a.service.DeleteWorkflow(wfID); err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "workflow deleted", "workflowId": wfID})
}

// ToggleWorkflow enables or disables a workflow
// POST /api/v1/workflows/toggle
func (a *WorkflowAPI) ToggleWorkflow(w http.ResponseWriter, r *http.Request) {
	var req domain.ToggleWorkflowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.WorkflowID == "" {
		errJSON(w, http.StatusBadRequest, "workflowId is required")
		return
	}
	if err := a.service.ToggleWorkflow(req.WorkflowID, req.Enabled); err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"status": "workflow toggle updated", "workflowId": req.WorkflowID, "enabled": req.Enabled})
}

// RunWorkflow handles manually running a workflow
// POST /api/v1/workflows/run
func (a *WorkflowAPI) RunWorkflow(w http.ResponseWriter, r *http.Request) {
	var req domain.RunWorkflowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.WorkflowID == "" || req.VaultID == "" {
		errJSON(w, http.StatusBadRequest, "workflowId and vaultId are required")
		return
	}
	run, err := a.service.RunWorkflow(r.Context(), &req)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, run)
}

// GetHistory returns execution history runs and logs
// GET /api/v1/workflows/history?vaultId=<id>&workflowId=<id>
func (a *WorkflowAPI) GetHistory(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	wfID := r.URL.Query().Get("workflowId")

	if wfID != "" {
		runs, err := a.service.ListRuns(wfID)
		if err != nil {
			errJSON(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, map[string]interface{}{"runs": runs, "total": len(runs)})
		return
	}

	if vaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId or workflowId is required")
		return
	}
	runs, err := a.service.ListAllRuns(vaultID)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"runs": runs, "total": len(runs)})
}

// GetStats returns vault observability metrics
// GET /api/v1/workflows/stats?vaultId=<id>
func (a *WorkflowAPI) GetStats(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	if vaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	stats, err := a.service.GetStats(vaultID)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, stats)
}
