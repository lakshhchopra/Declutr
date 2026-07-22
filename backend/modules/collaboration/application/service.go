package application

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/diablovocado/declutr/modules/collaboration/domain"
	"github.com/diablovocado/declutr/modules/collaboration/repository"
)

// CollaborationService manages secure resource sharing, permissions, link tokens, comments, and audit trails
type CollaborationService struct {
	repo repository.CollaborationRepository
}

// NewCollaborationService creates a new CollaborationService
func NewCollaborationService(repo repository.CollaborationRepository) *CollaborationService {
	return &CollaborationService{repo: repo}
}

// CheckPermission evaluates if a role possesses permission for a given action
func CheckPermission(role domain.MemberRole, action string) bool {
	switch role {
	case domain.RoleOwner, domain.RoleCoOwner:
		return true // Full permissions
	case domain.RoleEdit:
		return action == "VIEW" || action == "DOWNLOAD" || action == "EDIT" || action == "COMMENT" || action == "SHARE"
	case domain.RoleCommentOnly:
		return action == "VIEW" || action == "DOWNLOAD" || action == "COMMENT"
	case domain.RoleReadOnly:
		return action == "VIEW" || action == "DOWNLOAD"
	default:
		return false
	}
}

// CreateShare creates a new shared resource container
func (s *CollaborationService) CreateShare(req *domain.CreateShareRequest) (*domain.Share, error) {
	if req.VaultID == "" || req.ResourceID == "" || req.Title == "" {
		return nil, fmt.Errorf("collaboration: vaultId, resourceId, and title are required")
	}

	shareID := "share-" + uuid.New().String()[:8]
	now := time.Now()

	share := &domain.Share{
		ShareID:      shareID,
		VaultID:      req.VaultID,
		ResourceType: req.ResourceType,
		ResourceID:   req.ResourceID,
		Title:        req.Title,
		AccessType:   req.AccessType,
		CreatedBy:    "USER",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	// Add owner member
	_ = s.repo.AddMember(&domain.ShareMember{
		MemberID: "mem-" + uuid.New().String()[:8],
		ShareID:  shareID,
		UserID:   "usr-owner",
		Email:    "owner@declutr.local",
		Role:     domain.RoleOwner,
		JoinedAt: now,
	})

	// Log audit event
	_ = s.repo.AddActivity(&domain.ShareActivity{
		ActivityID: "act-" + uuid.New().String()[:8],
		ShareID:    shareID,
		VaultID:    req.VaultID,
		ActorID:    "usr-owner",
		ActorName:  "Vault Owner",
		ActionType: domain.ActionShared,
		Details:    map[string]interface{}{"title": req.Title, "resourceType": req.ResourceType},
		CreatedAt:  now,
	})

	if err := s.repo.CreateShare(share); err != nil {
		return nil, err
	}
	return share, nil
}

// ListShares returns all shares for a vault
func (s *CollaborationService) ListShares(vaultID string) ([]*domain.Share, error) {
	if vaultID == "" {
		return nil, fmt.Errorf("collaboration: vaultId is required")
	}
	return s.repo.ListShares(vaultID)
}

// DeleteShare revokes a share and all associated access
func (s *CollaborationService) DeleteShare(shareID string) error {
	if shareID == "" {
		return fmt.Errorf("collaboration: shareId is required")
	}
	return s.repo.DeleteShare(shareID)
}

// InviteUser sends an invitation to join a share
func (s *CollaborationService) InviteUser(req *domain.InviteRequest) (*domain.ShareInvitation, error) {
	if req.ShareID == "" || req.InviteeEmail == "" {
		return nil, fmt.Errorf("collaboration: shareId and inviteeEmail are required")
	}

	token := "inv-tok-" + uuid.New().String()[:12]
	now := time.Now()
	exp := now.Add(7 * 24 * time.Hour)

	inv := &domain.ShareInvitation{
		InvitationID: "inv-" + uuid.New().String()[:8],
		ShareID:      req.ShareID,
		InviterID:    req.InviterID,
		InviteeEmail: req.InviteeEmail,
		Role:         req.Role,
		Status:       domain.InvitePending,
		Token:        token,
		CreatedAt:    now,
		ExpiresAt:    &exp,
	}

	if err := s.repo.CreateInvitation(inv); err != nil {
		return nil, err
	}
	return inv, nil
}

// AcceptInvitation handles accepting an invitation token
func (s *CollaborationService) AcceptInvitation(token string, userID string) error {
	inv, err := s.repo.GetInvitationByToken(token)
	if err != nil {
		return err
	}
	if inv.Status != domain.InvitePending {
		return fmt.Errorf("invitation is no longer pending")
	}

	now := time.Now()
	_ = s.repo.UpdateInvitationStatus(inv.InvitationID, domain.InviteAccepted)

	_ = s.repo.AddMember(&domain.ShareMember{
		MemberID: "mem-" + uuid.New().String()[:8],
		ShareID:  inv.ShareID,
		UserID:   userID,
		Email:    inv.InviteeEmail,
		Role:     inv.Role,
		JoinedAt: now,
	})

	_ = s.repo.AddActivity(&domain.ShareActivity{
		ActivityID: "act-" + uuid.New().String()[:8],
		ShareID:    inv.ShareID,
		VaultID:    "vault-demo",
		ActorID:    userID,
		ActorName:  inv.InviteeEmail,
		ActionType: domain.ActionInviteAccepted,
		Details:    map[string]interface{}{"role": inv.Role},
		CreatedAt:  now,
	})

	return nil
}

// CreateLink creates a share link with optional password protection
func (s *CollaborationService) CreateLink(req *domain.CreateLinkRequest) (*domain.ShareLink, error) {
	if req.ShareID == "" {
		return nil, fmt.Errorf("collaboration: shareId is required")
	}

	now := time.Now()
	var exp *time.Time
	if req.ExpiresInDays > 0 {
		e := now.Add(time.Duration(req.ExpiresInDays) * 24 * time.Hour)
		exp = &e
	}

	link := &domain.ShareLink{
		LinkID:              "link-" + uuid.New().String()[:8],
		ShareID:             req.ShareID,
		LinkToken:           "link-tok-" + uuid.New().String()[:12],
		IsPasswordProtected: req.IsPasswordProtected,
		DisableDownload:     req.DisableDownload,
		ExpiresAt:           exp,
		CreatedAt:           now,
	}

	if err := s.repo.CreateLink(link); err != nil {
		return nil, err
	}
	return link, nil
}

// RevokeLink revokes a share link
func (s *CollaborationService) RevokeLink(linkID string) error {
	return s.repo.RevokeLink(linkID)
}

// AddComment adds a threaded comment or reply
func (s *CollaborationService) AddComment(req *domain.AddCommentRequest) (*domain.ShareComment, error) {
	if req.ShareID == "" || req.Content == "" {
		return nil, fmt.Errorf("collaboration: shareId and content are required")
	}

	now := time.Now()
	comment := &domain.ShareComment{
		CommentID:       "cmnt-" + uuid.New().String()[:8],
		ShareID:         req.ShareID,
		UserID:          req.UserID,
		UserName:        req.UserName,
		Content:         req.Content,
		ParentCommentID: req.ParentCommentID,
		IsResolved:      false,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	_ = s.repo.AddActivity(&domain.ShareActivity{
		ActivityID: "act-" + uuid.New().String()[:8],
		ShareID:    req.ShareID,
		VaultID:    "vault-demo",
		ActorID:    req.UserID,
		ActorName:  req.UserName,
		ActionType: domain.ActionCommented,
		Details:    map[string]interface{}{"content": req.Content},
		CreatedAt:  now,
	})

	if err := s.repo.AddComment(comment); err != nil {
		return nil, err
	}
	return comment, nil
}

// ListComments returns comments for a share
func (s *CollaborationService) ListComments(shareID string) ([]*domain.ShareComment, error) {
	return s.repo.ListComments(shareID)
}

// ListActivity returns audit activity history for a share or vault
func (s *CollaborationService) ListActivity(vaultID string, shareID string) ([]*domain.ShareActivity, error) {
	if shareID != "" {
		return s.repo.GetActivity(shareID)
	}
	return s.repo.ListAllActivity(vaultID)
}

// GetStats returns vault collaboration stats
func (s *CollaborationService) GetStats(vaultID string) (*domain.ShareStats, error) {
	return s.repo.GetStats(vaultID)
}
