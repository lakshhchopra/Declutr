package transport

import (
	"encoding/json"
	"net/http"

	"github.com/diablovocado/declutr/modules/security/application"
	"github.com/diablovocado/declutr/modules/security/domain"
)

// SecurityAPI handles HTTP endpoints for Security Center, Audit Hub & Trust Platform
type SecurityAPI struct {
	service *application.SecurityCenterService
}

// NewSecurityAPI creates a new SecurityAPI instance
func NewSecurityAPI(service *application.SecurityCenterService) *SecurityAPI {
	return &SecurityAPI{service: service}
}

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func errJSON(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, map[string]string{"error": msg})
}

// GetDashboard returns the complete security dashboard payload
// GET /api/v1/security/dashboard?vaultId=<id>
func (a *SecurityAPI) GetDashboard(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	if vaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	dash, err := a.service.GetDashboard(vaultID)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, dash)
}

// ListAuditEvents returns audit log history
// GET /api/v1/security/audit?vaultId=<id>&category=<cat>
func (a *SecurityAPI) ListAuditEvents(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	category := domain.AuditCategory(r.URL.Query().Get("category"))
	if vaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	events, err := a.service.ListAuditEvents(vaultID, category, 50)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"events": events, "total": len(events)})
}

// ListSessions returns active user sessions
// GET /api/v1/security/sessions?vaultId=<id>
func (a *SecurityAPI) ListSessions(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	if vaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	sessions, err := a.service.ListSessions(vaultID)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"sessions": sessions, "total": len(sessions)})
}

// TerminateSession terminates single or all active user sessions
// POST /api/v1/security/sessions/terminate
func (a *SecurityAPI) TerminateSession(w http.ResponseWriter, r *http.Request) {
	var req domain.TerminateSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.VaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	if err := a.service.TerminateSession(&req); err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "session(s) terminated"})
}

// ListDevices returns device registry
// GET /api/v1/security/devices?vaultId=<id>
func (a *SecurityAPI) ListDevices(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	if vaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	devices, err := a.service.ListDevices(vaultID)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"devices": devices, "total": len(devices)})
}

// SetDeviceTrust updates device trust status
// POST /api/v1/security/devices/trust
func (a *SecurityAPI) SetDeviceTrust(w http.ResponseWriter, r *http.Request) {
	var req domain.TrustDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.DeviceID == "" {
		errJSON(w, http.StatusBadRequest, "deviceId is required")
		return
	}
	if err := a.service.SetDeviceTrust(&req); err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"status": "device trust updated", "deviceId": req.DeviceID, "isTrusted": req.Trust})
}

// GetRiskAssessment returns risk engine evaluation
// GET /api/v1/security/risk?vaultId=<id>
func (a *SecurityAPI) GetRiskAssessment(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	if vaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	risk, err := a.service.GetRiskAssessment(vaultID)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, risk)
}

// GetRecommendations returns actionable security recommendations
// GET /api/v1/security/recommendations?vaultId=<id>
func (a *SecurityAPI) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	if vaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	recs, err := a.service.GetRecommendations(vaultID)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"recommendations": recs, "total": len(recs)})
}
