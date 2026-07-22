package transport

import (
	"encoding/json"
	"net/http"

	"github.com/diablovocado/declutr/modules/backup/application"
	"github.com/diablovocado/declutr/modules/backup/domain"
)

// BackupAPI handles HTTP endpoints for Backup, Disaster Recovery & Business Continuity
type BackupAPI struct {
	service *application.BackupService
}

// NewBackupAPI creates a new BackupAPI instance
func NewBackupAPI(service *application.BackupService) *BackupAPI {
	return &BackupAPI{service: service}
}

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func errJSON(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, map[string]string{"error": msg})
}

// ListBackups returns all backups for a vault
// GET /api/v1/backups?vaultId=<id>
func (a *BackupAPI) ListBackups(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	if vaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	backups, err := a.service.ListBackups(vaultID)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"backups": backups, "total": len(backups)})
}

// CreateBackup handles manual snapshot backup creation
// POST /api/v1/backups
func (a *BackupAPI) CreateBackup(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateBackupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.VaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	if req.BackupType == "" {
		req.BackupType = domain.BackupManual
	}
	b, err := a.service.CreateBackup(r.Context(), &req)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, b)
}

// GetBackupDetails returns backup details and its manifest
// GET /api/v1/backups/detail?backupId=<id>
func (a *BackupAPI) GetBackupDetails(w http.ResponseWriter, r *http.Request) {
	backupID := r.URL.Query().Get("backupId")
	if backupID == "" {
		errJSON(w, http.StatusBadRequest, "backupId is required")
		return
	}
	b, manifest, err := a.service.GetBackupDetails(backupID)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"backup": b, "manifest": manifest})
}

// ConfigureSchedule configures vault backup automation schedule and retention policy
// POST /api/v1/backups/schedule
func (a *BackupAPI) ConfigureSchedule(w http.ResponseWriter, r *http.Request) {
	var req domain.ScheduleBackupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.VaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	sched, err := a.service.ConfigureSchedule(&req)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, sched)
}

// RestoreBackup triggers disaster recovery vault restoration
// POST /api/v1/backups/restore
func (a *BackupAPI) RestoreBackup(w http.ResponseWriter, r *http.Request) {
	var req domain.RestoreBackupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.VaultID == "" || req.BackupID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId and backupId are required")
		return
	}
	job, err := a.service.RestoreBackup(&req)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, job)
}

// VerifyIntegrity runs SHA-256 integrity check on backup payload
// POST /api/v1/backups/verify
func (a *BackupAPI) VerifyIntegrity(w http.ResponseWriter, r *http.Request) {
	var req domain.VerifyBackupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.BackupID == "" {
		errJSON(w, http.StatusBadRequest, "backupId is required")
		return
	}
	ok, msg, err := a.service.VerifyIntegrity(&req)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"isValid": ok, "message": msg, "backupId": req.BackupID})
}

// CancelJob cancels an active backup or restore job
// POST /api/v1/backups/cancel?jobId=<id>
func (a *BackupAPI) CancelJob(w http.ResponseWriter, r *http.Request) {
	jobID := r.URL.Query().Get("jobId")
	if jobID == "" {
		errJSON(w, http.StatusBadRequest, "jobId is required")
		return
	}
	if err := a.service.CancelJob(jobID); err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "job cancelled", "jobId": jobID})
}

// GetStats returns vault backup and disaster recovery stats
// GET /api/v1/backups/stats?vaultId=<id>
func (a *BackupAPI) GetStats(w http.ResponseWriter, r *http.Request) {
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
