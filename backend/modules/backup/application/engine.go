package application

import (
	"context"
	"fmt"
	"log"

	"github.com/diablovocado/declutr/modules/backup/domain"
)

// DisasterRecoveryEngine orchestrates background backup scheduling and automated disaster recovery verification
type DisasterRecoveryEngine struct {
	service *BackupService
}

// NewDisasterRecoveryEngine creates a new DisasterRecoveryEngine instance
func NewDisasterRecoveryEngine(service *BackupService) *DisasterRecoveryEngine {
	return &DisasterRecoveryEngine{service: service}
}

// RunScheduledBackup triggers a scheduled vault backup operation
func (e *DisasterRecoveryEngine) RunScheduledBackup(ctx context.Context, vaultID string) (*domain.Backup, error) {
	log.Printf("[DisasterRecoveryEngine] Executing scheduled backup for vault %s", vaultID)

	b, err := e.service.CreateBackup(ctx, &domain.CreateBackupRequest{
		VaultID:    vaultID,
		BackupType: domain.BackupScheduled,
	})
	if err != nil {
		return nil, fmt.Errorf("scheduled backup failed: %w", err)
	}
	return b, nil
}
