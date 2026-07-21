-- Migration 007: Create Metadata Engine Tables

CREATE TABLE IF NOT EXISTS asset_metadata (
    asset_id UUID PRIMARY KEY REFERENCES assets(asset_id) ON DELETE CASCADE,
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    filename TEXT NOT NULL,
    extension VARCHAR(50),
    mime_type VARCHAR(100),
    file_size BIGINT NOT NULL,
    checksum VARCHAR(64),
    hash VARCHAR(64),
    encoding VARCHAR(50),
    created_date TIMESTAMPTZ,
    modified_date TIMESTAMPTZ,
    upload_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_extracted_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS asset_properties (
    asset_id UUID PRIMARY KEY REFERENCES asset_metadata(asset_id) ON DELETE CASCADE,
    properties JSONB NOT NULL DEFAULT '{}'::jsonb
);

CREATE TABLE IF NOT EXISTS asset_exif (
    asset_id UUID PRIMARY KEY REFERENCES asset_metadata(asset_id) ON DELETE CASCADE,
    camera_make VARCHAR(100),
    camera_model VARCHAR(100),
    lens VARCHAR(150),
    gps_lat DOUBLE PRECISION,
    gps_long DOUBLE PRECISION,
    iso INT,
    exposure VARCHAR(50),
    f_stop DOUBLE PRECISION,
    focal_length DOUBLE PRECISION,
    date_taken TIMESTAMPTZ,
    raw_data JSONB
);

CREATE TABLE IF NOT EXISTS metadata_versions (
    version_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_id UUID NOT NULL REFERENCES asset_metadata(asset_id) ON DELETE CASCADE,
    source VARCHAR(100) NOT NULL, -- e.g., 'SYSTEM_EXTRACTOR', 'USER_OVERRIDE'
    extractor_version VARCHAR(50) NOT NULL,
    confidence FLOAT,
    snapshot JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_asset_metadata_vault_id ON asset_metadata(vault_id);
CREATE INDEX IF NOT EXISTS idx_asset_metadata_mime_type ON asset_metadata(mime_type);
CREATE INDEX IF NOT EXISTS idx_metadata_versions_asset_id ON metadata_versions(asset_id);
