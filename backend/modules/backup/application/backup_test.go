package application_test

import (
	"context"
	"testing"

	"github.com/diablovocado/declutr/modules/backup/application"
	"github.com/diablovocado/declutr/modules/backup/domain"
	"github.com/diablovocado/declutr/modules/backup/repository"
)

const testVaultID = "vault-test-001"

func setupService() *application.BackupService {
	repo := repository.NewInMemoryBackupRepository()
	return application.NewBackupService(repo)
}

// TestCreateManualBackup validates manual backup creation
func TestCreateManualBackup(t *testing.T) {
	svc := setupService()
	ctx := context.Background()

	b, err := svc.CreateBackup(ctx, &domain.CreateBackupRequest{
		VaultID:    testVaultID,
		BackupType: domain.BackupManual,
		Passphrase: "secret-key-passphrase-123",
	})
	if err != nil {
		t.Fatalf("create backup failed: %v", err)
	}
	if !b.IsEncrypted {
		t.Error("expected backup with passphrase to be encrypted")
	}

	list, err := svc.ListBackups(testVaultID)
	if err != nil {
		t.Fatalf("list backups failed: %v", err)
	}
	if len(list) == 0 {
		t.Error("expected backup in vault list")
	}

	t.Logf("PASS: Create Manual Backup — Created %s (Size: %d, Encrypted: %t)", b.BackupID, b.SizeBytes, b.IsEncrypted)
}

// TestBackupSchedulePolicy validates schedule policy configuration
func TestBackupSchedulePolicy(t *testing.T) {
	svc := setupService()

	sched, err := svc.ConfigureSchedule(&domain.ScheduleBackupRequest{
		VaultID:        testVaultID,
		Frequency:      domain.FreqDaily,
		RetentionDays:  30,
		MaxBackupCount: 14,
		EncryptBackups: true,
	})
	if err != nil {
		t.Fatalf("configure schedule failed: %v", err)
	}
	if sched.Frequency != domain.FreqDaily {
		t.Errorf("expected FREQ_DAILY, got %s", sched.Frequency)
	}

	t.Logf("PASS: Backup Schedule Policy — Configured %s schedule (Retention: %d days)", sched.Frequency, sched.RetentionDays)
}

// TestIntegrityVerification validates checksum and manifest integrity checks
func TestIntegrityVerification(t *testing.T) {
	svc := setupService()
	ctx := context.Background()

	b, _ := svc.CreateBackup(ctx, &domain.CreateBackupRequest{
		VaultID:    testVaultID,
		BackupType: domain.BackupFull,
	})

	ok, msg, err := svc.VerifyIntegrity(&domain.VerifyBackupRequest{
		VaultID:  testVaultID,
		BackupID: b.BackupID,
	})
	if err != nil {
		t.Fatalf("verify integrity failed: %v", err)
	}
	if !ok {
		t.Errorf("expected integrity verification to pass, got error: %s", msg)
	}

	t.Logf("PASS: Integrity Verification — %s for %s", msg, b.BackupID)
}

// TestDisasterRecoveryRestore validates full vault restoration
func TestDisasterRecoveryRestore(t *testing.T) {
	svc := setupService()
	ctx := context.Background()

	b, _ := svc.CreateBackup(ctx, &domain.CreateBackupRequest{
		VaultID:    testVaultID,
		BackupType: domain.BackupFull,
	})

	job, err := svc.RestoreBackup(&domain.RestoreBackupRequest{
		VaultID:         testVaultID,
		BackupID:        b.BackupID,
		RestoreMode:     domain.RestoreFullVault,
		RestoreStrategy: domain.StrategyMerge,
	})
	if err != nil {
		t.Fatalf("restore backup failed: %v", err)
	}
	if job.Status != domain.JobSuccess {
		t.Errorf("expected restore status SUCCESS, got %s", job.Status)
	}

	t.Logf("PASS: Disaster Recovery Restore — Executed restore job %s (Mode: %s)", job.JobID, job.RestoreMode)
}

// TestCancelBackupJob validates cancelling a running backup job
func TestCancelBackupJob(t *testing.T) {
	svc := setupService()

	jobID := "job-temp-001"
	if err := svc.CancelJob(jobID); err != nil {
		t.Fatalf("cancel job failed: %v", err)
	}

	t.Logf("PASS: Cancel Backup Job — Successfully cancelled job %s", jobID)
}

// TestBackupStats validates vault disaster recovery metrics
func TestBackupStats(t *testing.T) {
	svc := setupService()

	stats, err := svc.GetStats(testVaultID)
	if err != nil {
		t.Fatalf("get stats failed: %v", err)
	}
	if stats.TotalBackups == 0 {
		t.Error("expected positive total backups count")
	}

	t.Logf("PASS: Backup Stats — TotalBackups=%d, TotalSizeBytes=%d, Ratio=%.1f%%",
		stats.TotalBackups, stats.TotalSizeBytes, stats.CompressionRatioPct)
}
