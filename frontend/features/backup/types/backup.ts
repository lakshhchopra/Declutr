// Backup, Disaster Recovery & Business Continuity TypeScript types

export type BackupType =
  | 'MANUAL'
  | 'SCHEDULED'
  | 'INCREMENTAL'
  | 'FULL'
  | 'ENCRYPTED'
  | 'OFFLINE'
  | 'COLD_STORAGE';

export type BackupStatus = 'PENDING' | 'IN_PROGRESS' | 'COMPLETED' | 'FAILED' | 'CORRUPTED';

export type RestoreMode =
  | 'FULL_VAULT'
  | 'SELECTIVE'
  | 'ASSETS_ONLY'
  | 'METADATA_ONLY'
  | 'AI_STATE_ONLY'
  | 'WORKFLOWS_ONLY'
  | 'SETTINGS_ONLY';

export type RestoreStrategy =
  | 'OVERWRITE_EXISTING'
  | 'RESTORE_AS_NEW_VAULT'
  | 'MERGE_RESTORE'
  | 'DRY_RUN';

export type ScheduleFrequency = 'MANUAL' | 'DAILY' | 'WEEKLY' | 'MONTHLY' | 'CUSTOM_CRON';

export interface Backup {
  backupId: string;
  vaultId: string;
  backupType: BackupType;
  status: BackupStatus;
  sizeBytes: number;
  compressedSizeBytes: number;
  checksum: string;
  storagePath: string;
  isEncrypted: boolean;
  createdAt: string;
}

export interface BackupManifest {
  manifestId: string;
  backupId: string;
  vaultId: string;
  totalAssets: number;
  totalMemories: number;
  totalWorkflows: number;
  manifestData: Record<string, unknown>;
  createdAt: string;
}

export interface BackupSchedule {
  vaultId: string;
  frequency: ScheduleFrequency;
  cronExpression?: string;
  retentionDays: number;
  maxBackupCount: number;
  encryptBackups: boolean;
  isEnabled: boolean;
  nextScheduledAt: string;
}

export interface RestoreJob {
  jobId: string;
  vaultId: string;
  backupId: string;
  restoreMode: RestoreMode;
  restoreStrategy: RestoreStrategy;
  status: 'RUNNING' | 'SUCCESS' | 'FAILED' | 'CANCELLED';
  restoredBy: string;
  startedAt: string;
  completedAt?: string;
}

export interface BackupStats {
  vaultId: string;
  totalBackups: number;
  totalSizeBytes: number;
  totalRestores: number;
  lastBackupAt: string;
  lastVerifySuccess: boolean;
  compressionRatioPct: number;
}
