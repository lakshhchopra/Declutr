package domain

import "time"

// BackupType specifies the nature of the backup package
type BackupType string

const (
	BackupManual      BackupType = "MANUAL"
	BackupScheduled   BackupType = "SCHEDULED"
	BackupIncremental BackupType = "INCREMENTAL"
	BackupFull        BackupType = "FULL"
	BackupEncrypted   BackupType = "ENCRYPTED"
	BackupOffline     BackupType = "OFFLINE"
	BackupColdStorage BackupType = "COLD_STORAGE"
)

// BackupStatus defines the execution status of a backup
type BackupStatus string

const (
	BackupPending    BackupStatus = "PENDING"
	BackupInProgress BackupStatus = "IN_PROGRESS"
	BackupCompleted  BackupStatus = "COMPLETED"
	BackupFailed     BackupStatus = "FAILED"
	BackupCorrupted  BackupStatus = "CORRUPTED"
)

// JobType defines async task classification
type JobType string

const (
	JobBackup         JobType = "BACKUP"
	JobRestore        JobType = "RESTORE"
	JobIntegrityCheck JobType = "INTEGRITY_CHECK"
)

// JobStatus defines execution status of async jobs
type JobStatus string

const (
	JobRunning   JobStatus = "RUNNING"
	JobSuccess   JobStatus = "SUCCESS"
	JobFailed    JobStatus = "FAILED"
	JobCancelled JobStatus = "CANCELLED"
)

// RestoreMode defines scope of recovery
type RestoreMode string

const (
	RestoreFullVault    RestoreMode = "FULL_VAULT"
	RestoreSelective    RestoreMode = "SELECTIVE"
	RestoreAssetsOnly   RestoreMode = "ASSETS_ONLY"
	RestoreMetadataOnly RestoreMode = "METADATA_ONLY"
	RestoreAIStateOnly  RestoreMode = "AI_STATE_ONLY"
	RestoreWorkflows    RestoreMode = "WORKFLOWS_ONLY"
	RestoreSettingsOnly RestoreMode = "SETTINGS_ONLY"
)

// RestoreStrategy specifies how restored data is merged/applied
type RestoreStrategy string

const (
	StrategyOverwriteExisting RestoreStrategy = "OVERWRITE_EXISTING"
	StrategyNewVault            RestoreStrategy = "RESTORE_AS_NEW_VAULT"
	StrategyMerge               RestoreStrategy = "MERGE_RESTORE"
	StrategyDryRun              RestoreStrategy = "DRY_RUN"
)

// ScheduleFrequency defines backup automation frequency
type ScheduleFrequency string

const (
	FreqManual  ScheduleFrequency = "MANUAL"
	FreqDaily   ScheduleFrequency = "DAILY"
	FreqWeekly  ScheduleFrequency = "WEEKLY"
	FreqMonthly ScheduleFrequency = "MONTHLY"
	FreqCustom  ScheduleFrequency = "CUSTOM_CRON"
)

// Backup model
type Backup struct {
	BackupID           string       `json:"backupId"`
	VaultID            string       `json:"vaultId"`
	BackupType         BackupType   `json:"backupType"`
	Status             BackupStatus `json:"status"`
	SizeBytes          int64        `json:"sizeBytes"`
	CompressedSizeBytes int64        `json:"compressedSizeBytes"`
	Checksum           string       `json:"checksum"`
	StoragePath        string       `json:"storagePath"`
	IsEncrypted        bool         `json:"isEncrypted"`
	CreatedAt          time.Time    `json:"createdAt"`
}

// BackupJob model
type BackupJob struct {
	JobID       string     `json:"jobId"`
	VaultID     string     `json:"vaultId"`
	JobType     JobType    `json:"jobType"`
	Status      JobStatus  `json:"status"`
	ProgressPct int        `json:"progressPct"`
	ErrorMsg    string     `json:"errorMsg,omitempty"`
	StartedAt   time.Time  `json:"startedAt"`
	CompletedAt *time.Time `json:"completedAt,omitempty"`
}

// BackupManifest model
type BackupManifest struct {
	ManifestID     string                 `json:"manifestId"`
	BackupID       string                 `json:"backupId"`
	VaultID        string                 `json:"vaultId"`
	TotalAssets    int                    `json:"totalAssets"`
	TotalMemories  int                    `json:"totalMemories"`
	TotalWorkflows int                    `json:"totalWorkflows"`
	ManifestData   map[string]interface{} `json:"manifestData"`
	CreatedAt      time.Time              `json:"createdAt"`
}

// BackupSchedule configuration payload
type BackupSchedule struct {
	VaultID         string            `json:"vaultId"`
	Frequency       ScheduleFrequency `json:"frequency"`
	CronExpression  string            `json:"cronExpression,omitempty"`
	RetentionDays   int               `json:"retentionDays"`
	MaxBackupCount  int               `json:"maxBackupCount"`
	EncryptBackups  bool              `json:"encryptBackups"`
	IsEnabled       bool              `json:"isEnabled"`
	NextScheduledAt time.Time         `json:"nextScheduledAt"`
}

// RestoreJob execution model
type RestoreJob struct {
	JobID           string          `json:"jobId"`
	VaultID         string          `json:"vaultId"`
	BackupID        string          `json:"backupId"`
	RestoreMode     RestoreMode     `json:"restoreMode"`
	RestoreStrategy RestoreStrategy `json:"restoreStrategy"`
	Status          JobStatus       `json:"status"`
	RestoredBy      string          `json:"restoredBy"`
	StartedAt       time.Time       `json:"startedAt"`
	CompletedAt     *time.Time      `json:"completedAt,omitempty"`
}

// BackupStats metrics model
type BackupStats struct {
	VaultID             string    `json:"vaultId"`
	TotalBackups        int       `json:"totalBackups"`
	TotalSizeBytes      int64     `json:"totalSizeBytes"`
	TotalRestores       int       `json:"totalRestores"`
	LastBackupAt        time.Time `json:"lastBackupAt"`
	LastVerifySuccess   bool      `json:"lastVerifySuccess"`
	CompressionRatioPct float64   `json:"compressionRatioPct"`
}

// Request DTOs

type CreateBackupRequest struct {
	VaultID    string     `json:"vaultId"`
	BackupType BackupType `json:"backupType"`
	Passphrase string     `json:"passphrase,omitempty"`
}

type ScheduleBackupRequest struct {
	VaultID        string            `json:"vaultId"`
	Frequency      ScheduleFrequency `json:"frequency"`
	CronExpression string            `json:"cronExpression,omitempty"`
	RetentionDays  int               `json:"retentionDays"`
	MaxBackupCount int               `json:"maxBackupCount"`
	EncryptBackups bool              `json:"encryptBackups"`
}

type RestoreBackupRequest struct {
	VaultID         string          `json:"vaultId"`
	BackupID        string          `json:"backupId"`
	RestoreMode     RestoreMode     `json:"restoreMode"`
	RestoreStrategy RestoreStrategy `json:"restoreStrategy"`
	Passphrase      string          `json:"passphrase,omitempty"`
}

type VerifyBackupRequest struct {
	VaultID  string `json:"vaultId"`
	BackupID string `json:"backupId"`
}
