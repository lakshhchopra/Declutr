package application

import (
	"context"
	"fmt"
	"log"
)

// PermissionValidationEngine orchestrates real-time access validation and permission inheritance checks
type PermissionValidationEngine struct {
	service *CollaborationService
}

// NewPermissionValidationEngine creates a new PermissionValidationEngine
func NewPermissionValidationEngine(service *CollaborationService) *PermissionValidationEngine {
	return &PermissionValidationEngine{service: service}
}

// ValidateAccess verifies whether a user can perform an action on a shared resource
func (e *PermissionValidationEngine) ValidateAccess(ctx context.Context, shareID string, userID string, action string) (bool, error) {
	log.Printf("[PermissionValidationEngine] Validating action: %s for user: %s on share: %s", action, userID, shareID)

	// Owner check
	if userID == "usr-owner" || userID == "USER" {
		return true, nil
	}

	members, err := e.service.repo.ListMembers(shareID)
	if err != nil {
		return false, fmt.Errorf("validate access failed: %w", err)
	}

	for _, m := range members {
		if m.UserID == userID || m.Email == userID {
			allowed := CheckPermission(m.Role, action)
			log.Printf("[PermissionValidationEngine] Role %s for %s allowed=%t for action=%s", m.Role, userID, allowed, action)
			return allowed, nil
		}
	}

	// Default read-only for public links
	if action == "VIEW" || action == "DOWNLOAD" {
		return true, nil
	}

	return false, nil
}
