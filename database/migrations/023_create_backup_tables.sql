-- Migration 023: Create Backup, Disaster Recovery & Business Continuity Tables

-- 1. backups - core vault backup packages
CREATE TABLE IF NOT EXISTS backups (
    backup_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    backup_type VARCHAR(50) NOT NULL DEFAULT 'MANUAL', -- MANUAL, SCHEDULED, INCREMENTAL, FULL, ENCRYPTED, OFFLINE, COLD_STORAGE
    status VARCHAR(50) NOT NULL DEFAULT 'COMPLETED', -- PENDING, IN_PROGRESS, COMPLETED, FAILED, CORRUPTED
    size_bytes BIGINT NOT NULL DEFAULT 0,
    compressed_size_bytes BIGINT NOT NULL DEFAULT 0,
    checksum VARCHAR(255) NOT NULL DEFAULT '',
    storage_path VARCHAR(500) NOT NULL DEFAULT '',
    is_encrypted BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 2. backup_jobs - async backup & restore execution jobs
CREATE TABLE IF NOT EXISTS backup_jobs (
    job_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    job_type VARCHAR(50) NOT NULL, -- BACKUP, RESTORE, INTEGRITY_CHECK
    status VARCHAR(50) NOT NULL DEFAULT 'RUNNING', -- RUNNING, SUCCESS, FAILED, CANCELLED
    progress_pct INT NOT NULL DEFAULT 0,
    error_msg TEXT NOT NULL DEFAULT '',
    started_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMPTZ
);

-- 3. backup_files - individual files inside backup payload
CREATE TABLE IF NOT EXISTS backup_files (
    file_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    backup_id UUID NOT NULL REFERENCES backups(backup_id) ON DELETE CASCADE,
    file_path VARCHAR(500) NOT NULL,
    size_bytes BIGINT NOT NULL DEFAULT 0,
    checksum VARCHAR(255) NOT NULL DEFAULT '',
    content_type VARCHAR(100) NOT NULL DEFAULT 'application/octet-stream',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 4. backup_manifests - JSON manifests tracking full vault snapshot contents
CREATE TABLE IF NOT EXISTS backup_manifests (
    manifest_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    backup_id UUID NOT NULL REFERENCES backups(backup_id) ON DELETE CASCADE,
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    total_assets INT NOT NULL DEFAULT 0,
    total_memories INT NOT NULL DEFAULT 0,
    total_workflows INT NOT NULL DEFAULT 0,
    manifest_data JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 5. backup_history - audit trail for backup creation, verification, and purging
CREATE TABLE IF NOT EXISTS backup_history (
    history_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    backup_id UUID NOT NULL REFERENCES backups(backup_id) ON DELETE CASCADE,
    event_type VARCHAR(50) NOT NULL, -- CREATED, RESTORED, VERIFIED, PURGED, EXPIRED
    details JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 6. restore_jobs - tracking vault disaster recovery execution jobs
CREATE TABLE IF NOT EXISTS restore_jobs (
    job_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    backup_id UUID NOT NULL REFERENCES backups(backup_id) ON DELETE CASCADE,
    restore_mode VARCHAR(50) NOT NULL DEFAULT 'FULL_VAULT', -- FULL_VAULT, SELECTIVE, ASSETS_ONLY, METADATA_ONLY, AI_STATE_ONLY, WORKFLOWS_ONLY, SETTINGS_ONLY
    restore_strategy VARCHAR(50) NOT NULL DEFAULT 'RESTORE_AS_NEW_VAULT', -- OVERWRITE_EXISTING, RESTORE_AS_NEW_VAULT, MERGE_RESTORE, DRY_RUN
    status VARCHAR(50) NOT NULL DEFAULT 'SUCCESS',
    restored_by VARCHAR(100) NOT NULL DEFAULT 'USER',
    started_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMPTZ
);

-- 7. restore_history - disaster recovery log entries
CREATE TABLE IF NOT EXISTS restore_history (
    restore_hist_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    backup_id UUID NOT NULL REFERENCES backups(backup_id) ON DELETE CASCADE,
    restored_items_count INT NOT NULL DEFAULT 0,
    details JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_backups_vault ON backups(vault_id);
CREATE INDEX IF NOT EXISTS idx_backup_jobs_vault ON backup_jobs(vault_id);
