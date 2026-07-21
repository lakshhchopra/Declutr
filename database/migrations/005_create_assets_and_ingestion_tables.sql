-- Migration 005: Create Assets and Ingestion Pipeline Job Tables
CREATE TABLE IF NOT EXISTS assets (
    asset_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    owner_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    filename VARCHAR(255) NOT NULL,
    mime_type VARCHAR(100) NOT NULL,
    size_bytes BIGINT NOT NULL,
    checksum_sha256 VARCHAR(64) NOT NULL,
    storage_key VARCHAR(512) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'QUEUED', 
    -- Status states: QUEUED, UPLOADING, UPLOADED, VALIDATING, METADATA_PENDING, AI_PENDING, INDEXED_PENDING, READY, FAILED
    error_message TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS upload_jobs (
    job_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_id UUID NOT NULL REFERENCES assets(asset_id) ON DELETE CASCADE,
    job_type VARCHAR(50) NOT NULL, -- 'FILE_VALIDATION', 'OCR_PARSING', 'WHISPER_AUDIO', 'EMBEDDING', 'RELATIONSHIP'
    status VARCHAR(30) NOT NULL DEFAULT 'PENDING',
    progress_percentage INT NOT NULL DEFAULT 0,
    attempts INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_assets_vault_id ON assets(vault_id);
CREATE INDEX IF NOT EXISTS idx_assets_checksum ON assets(checksum_sha256);
CREATE INDEX IF NOT EXISTS idx_assets_status ON assets(status);
CREATE INDEX IF NOT EXISTS idx_upload_jobs_asset_id ON upload_jobs(asset_id);
