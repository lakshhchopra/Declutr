package transport

import (
	"encoding/json"
	"net/http"

	"github.com/diablovocado/declutr/modules/collaboration/application"
	"github.com/diablovocado/declutr/modules/collaboration/domain"
)

// CollaborationAPI handles HTTP endpoints for Secure Sharing & Collaboration Platform
type CollaborationAPI struct {
	service *application.CollaborationService
}

// NewCollaborationAPI creates a new CollaborationAPI instance
func NewCollaborationAPI(service *application.CollaborationService) *CollaborationAPI {
	return &CollaborationAPI{service: service}
}

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func errJSON(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, map[string]string{"error": msg})
}

// CreateShare handles creating a shared resource container
// POST /api/v1/shares
func (a *CollaborationAPI) CreateShare(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateShareRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.VaultID == "" || req.Title == "" {
		errJSON(w, http.StatusBadRequest, "invalid request body, missing vaultId or title")
		return
	}
	created, err := a.service.CreateShare(&req)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, created)
}

// ListShares returns all shares for a vault
// GET /api/v1/shares?vaultId=<id>
func (a *CollaborationAPI) ListShares(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	if vaultID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId is required")
		return
	}
	list, err := a.service.ListShares(vaultID)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"shares": list, "total": len(list)})
}

// DeleteShare revokes a share and all access
// DELETE /api/v1/shares?shareId=<id>
func (a *CollaborationAPI) DeleteShare(w http.ResponseWriter, r *http.Request) {
	shareID := r.URL.Query().Get("shareId")
	if shareID == "" {
		errJSON(w, http.StatusBadRequest, "shareId is required")
		return
	}
	if err := a.service.DeleteShare(shareID); err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "share deleted", "shareId": shareID})
}

// InviteUser sends an invitation to join a share
// POST /api/v1/shares/invite
func (a *CollaborationAPI) InviteUser(w http.ResponseWriter, r *http.Request) {
	var req domain.InviteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.ShareID == "" || req.InviteeEmail == "" {
		errJSON(w, http.StatusBadRequest, "shareId and inviteeEmail are required")
		return
	}
	inv, err := a.service.InviteUser(&req)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, inv)
}

// AcceptInvitation handles accepting an invitation token
// POST /api/v1/shares/invite/accept
func (a *CollaborationAPI) AcceptInvitation(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	userID := r.URL.Query().Get("userId")
	if token == "" || userID == "" {
		errJSON(w, http.StatusBadRequest, "token and userId are required")
		return
	}
	if err := a.service.AcceptInvitation(token, userID); err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "invitation accepted", "token": token})
}

// CreateLink creates a share link with password & expiry options
// POST /api/v1/shares/links
func (a *CollaborationAPI) CreateLink(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateLinkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.ShareID == "" {
		errJSON(w, http.StatusBadRequest, "shareId is required")
		return
	}
	link, err := a.service.CreateLink(&req)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, link)
}

// RevokeLink revokes a share link
// POST /api/v1/shares/links/revoke?linkId=<id>
func (a *CollaborationAPI) RevokeLink(w http.ResponseWriter, r *http.Request) {
	linkID := r.URL.Query().Get("linkId")
	if linkID == "" {
		errJSON(w, http.StatusBadRequest, "linkId is required")
		return
	}
	if err := a.service.RevokeLink(linkID); err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "link revoked", "linkId": linkID})
}

// AddComment handles adding a threaded comment
// POST /api/v1/shares/comments
func (a *CollaborationAPI) AddComment(w http.ResponseWriter, r *http.Request) {
	var req domain.AddCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.ShareID == "" || req.Content == "" {
		errJSON(w, http.StatusBadRequest, "shareId and content are required")
		return
	}
	cmnt, err := a.service.AddComment(&req)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, cmnt)
}

// ListComments returns comments for a share
// GET /api/v1/shares/comments?shareId=<id>
func (a *CollaborationAPI) ListComments(w http.ResponseWriter, r *http.Request) {
	shareID := r.URL.Query().Get("shareId")
	if shareID == "" {
		errJSON(w, http.StatusBadRequest, "shareId is required")
		return
	}
	list, err := a.service.ListComments(shareID)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"comments": list, "total": len(list)})
}

// GetActivity returns audit log history
// GET /api/v1/shares/activity?vaultId=<id>&shareId=<id>
func (a *CollaborationAPI) GetActivity(w http.ResponseWriter, r *http.Request) {
	vaultID := r.URL.Query().Get("vaultId")
	shareID := r.URL.Query().Get("shareId")

	if vaultID == "" && shareID == "" {
		errJSON(w, http.StatusBadRequest, "vaultId or shareId is required")
		return
	}
	activity, err := a.service.ListActivity(vaultID, shareID)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"activity": activity, "total": len(activity)})
}

// GetStats returns vault collaboration metrics
// GET /api/v1/shares/stats?vaultId=<id>
func (a *CollaborationAPI) GetStats(w http.ResponseWriter, r *http.Request) {
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
