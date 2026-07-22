package repository

import (
	"fmt"
	"sync"
	"time"

	"github.com/diablovocado/declutr/modules/backup/domain"
)

// BackupRepository defines persistence contract for backups, manifests, schedules, jobs, and restore records
type BackupRepository interface {
	CreateBackup(b *domain.Backup, manifest *domain.BackupManifest) error
	GetBackup(backupID string) (*domain.Backup, error)
	GetManifest(backupID string) (*domain.BackupManifest, error)
	ListBackups(vaultID string) ([]*domain.Backup, error)

	SaveSchedule(sched *domain.BackupSchedule) error
	GetSchedule(vaultID string) (*domain.BackupSchedule, error)

	CreateJob(job *domain.BackupJob) error
	UpdateJob(jobID string, status domain.JobStatus, progress int, errMsg string) error
	GetJob(jobID string) (*domain.BackupJob, error)

	CreateRestoreJob(job *domain.RestoreJob) error
	GetRestoreJob(jobID string) (*domain.RestoreJob, error)

	GetStats(vaultID string) (*domain.BackupStats, error)
	ClearAllData(vaultID string) error
}

// InMemoryBackupRepository is a thread-safe in-memory store
type InMemoryBackupRepository struct {
	mu          sync.RWMutex
	backups     map[string]*domain.Backup         // backupID -> Backup
	manifests   map[string]*domain.BackupManifest // backupID -> Manifest
	schedules   map[string]*domain.BackupSchedule // vaultID -> Schedule
	jobs        map[string]*domain.BackupJob      // jobID -> Job
	restoreJobs map[string]*domain.RestoreJob     // jobID -> RestoreJob
}

// NewInMemoryBackupRepository creates a new in-memory backup repository
func NewInMemoryBackupRepository() *InMemoryBackupRepository {
	return &InMemoryBackupRepository{
		backups:     make(map[string]*domain.Backup),
		manifests:   make(map[string]*domain.BackupManifest),
		schedules:   make(map[string]*domain.BackupSchedule),
		jobs:        make(map[string]*domain.BackupJob),
		restoreJobs: make(map[string]*domain.RestoreJob),
	}
}

func (r *InMemoryBackupRepository) CreateBackup(b *domain.Backup, manifest *domain.BackupManifest) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.backups[b.BackupID] = b
	if manifest != nil {
		r.manifests[b.BackupID] = manifest
	}
	return nil
}

func (r *InMemoryBackupRepository) GetBackup(backupID string) (*domain.Backup, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	b, ok := r.backups[backupID]
	if !ok {
		return nil, fmt.Errorf("backup %s not found", backupID)
	}
	return b, nil
}

func (r *InMemoryBackupRepository) GetManifest(backupID string) (*domain.BackupManifest, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	m, ok := r.manifests[backupID]
	if !ok {
		return nil, fmt.Errorf("manifest for backup %s not found", backupID)
	}
	return m, nil
}

func (r *InMemoryBackupRepository) ListBackups(vaultID string) ([]*domain.Backup, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var list []*domain.Backup
	for _, b := range r.backups {
		if b.VaultID == vaultID {
			list = append(list, b)
		}
	}
	if len(list) == 0 {
		list = defaultSampleBackups(vaultID)
		for _, b := range list {
			r.backups[b.BackupID] = b
			r.manifests[b.BackupID] = &domain.BackupManifest{
				ManifestID:     "mnf-" + b.BackupID,
				BackupID:       b.BackupID,
				VaultID:        vaultID,
				TotalAssets:    142,
				TotalMemories:  38,
				TotalWorkflows: 6,
				ManifestData:   map[string]interface{}{"vaultId": vaultID, "checksum": b.Checksum},
				CreatedAt:      b.CreatedAt,
			}
		}
	}
	return list, nil
}

func (r *InMemoryBackupRepository) SaveSchedule(sched *domain.BackupSchedule) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.schedules[sched.VaultID] = sched
	return nil
}

func (r *InMemoryBackupRepository) GetSchedule(vaultID string) (*domain.BackupSchedule, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	sched, ok := r.schedules[vaultID]
	if !ok {
		return &domain.BackupSchedule{
			VaultID:         vaultID,
			Frequency:       domain.FreqWeekly,
			RetentionDays:   30,
			MaxBackupCount:  10,
			EncryptBackups:  true,
			IsEnabled:       true,
			NextScheduledAt: time.Now().Add(3 * 24 * time.Hour),
		}, nil
	}
	return sched, nil
}

func (r *InMemoryBackupRepository) CreateJob(job *domain.BackupJob) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.jobs[job.JobID] = job
	return nil
}

func (r *InMemoryBackupRepository) UpdateJob(jobID string, status domain.JobStatus, progress int, errMsg string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	job, ok := r.jobs[jobID]
	if !ok {
		return fmt.Errorf("job %s not found", jobID)
	}
	job.Status = status
	job.ProgressPct = progress
	job.ErrorMsg = errMsg
	if status == domain.JobSuccess || status == domain.JobFailed || status == domain.JobCancelled {
		now := time.Now()
		job.CompletedAt = &now
	}
	return nil
}

func (r *InMemoryBackupRepository) GetJob(jobID string) (*domain.BackupJob, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	job, ok := r.jobs[jobID]
	if !ok {
		return nil, fmt.Errorf("job %s not found", jobID)
	}
	return job, nil
}

func (r *InMemoryBackupRepository) CreateRestoreJob(job *domain.RestoreJob) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.restoreJobs[job.JobID] = job
	return nil
}

func (r *InMemoryBackupRepository) GetRestoreJob(jobID string) (*domain.RestoreJob, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	job, ok := r.restoreJobs[jobID]
	if !ok {
		return nil, fmt.Errorf("restore job %s not found", jobID)
	}
	return job, nil
}

func (r *InMemoryBackupRepository) GetStats(vaultID string) (*domain.BackupStats, error) {
	list, _ := r.ListBackups(vaultID)

	var totalSize int64
	var lastBackup time.Time
	for _, b := range list {
		totalSize += b.CompressedSizeBytes
		if b.CreatedAt.After(lastBackup) {
			lastBackup = b.CreatedAt
		}
	}

	return &domain.BackupStats{
		VaultID:             vaultID,
		TotalBackups:        len(list),
		TotalSizeBytes:      totalSize,
		TotalRestores:       1,
		LastBackupAt:        lastBackup,
		LastVerifySuccess:   true,
		CompressionRatioPct: 42.5,
	}, nil
}

func (r *InMemoryBackupRepository) ClearAllData(vaultID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for id, b := range r.backups {
		if b.VaultID == vaultID {
			delete(r.backups, id)
			delete(r.manifests, id)
		}
	}
	return nil
}

// Sample Data Generators
func defaultSampleBackups(vaultID string) []*domain.Backup {
	now := time.Now()
	return []*domain.Backup{
		{
			BackupID:            "bkp-weekly-auto-001",
			VaultID:             vaultID,
			BackupType:          domain.BackupScheduled,
			Status:              domain.BackupCompleted,
			SizeBytes:           1024 * 1024 * 350,
			CompressedSizeBytes: 1024 * 1024 * 180,
			Checksum:            "sha256-full-weekly-987abc",
			StoragePath:         "/storage/backups/bkp-weekly-auto-001.declutr",
			IsEncrypted:         true,
			CreatedAt:           now.Add(-4 * 24 * time.Hour),
		},
		{
			BackupID:            "bkp-pre-upgrade-002",
			VaultID:             vaultID,
			BackupType:          domain.BackupManual,
			Status:              domain.BackupCompleted,
			SizeBytes:           1024 * 1024 * 360,
			CompressedSizeBytes: 1024 * 1024 * 185,
			Checksum:            "sha256-preupgrade-123def",
			StoragePath:         "/storage/backups/bkp-pre-upgrade-002.declutr",
			IsEncrypted:         true,
			CreatedAt:           now.Add(-1 * 24 * time.Hour),
		},
	}
}
