-- Migration 024: Create Security Center, Audit Hub & Trust Platform Tables

-- 1. security_events - raw security events log
CREATE TABLE IF NOT EXISTS security_events (
    event_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    event_type VARCHAR(50) NOT NULL, -- AUTHENTICATION, SESSION_TERMINATED, DEVICE_ADDED, RISK_DETECTED, PERMISSION_CHANGED, WORKFLOW_ALERT
    severity VARCHAR(20) NOT NULL DEFAULT 'LOW', -- LOW, MEDIUM, HIGH, CRITICAL
    actor_id VARCHAR(100) NOT NULL DEFAULT 'USER',
    ip_address VARCHAR(45) NOT NULL DEFAULT '',
    device_info VARCHAR(255) NOT NULL DEFAULT '',
    details JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 2. security_scores - aggregated security posture score snapshots
CREATE TABLE IF NOT EXISTS security_scores (
    score_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    score INT NOT NULL DEFAULT 100, -- 0 to 100
    grade VARCHAR(5) NOT NULL DEFAULT 'A', -- A, B, C, D, F
    status VARCHAR(50) NOT NULL DEFAULT 'HEALTHY', -- HEALTHY, ATTENTION_REQUIRED, CRITICAL
    factors JSONB NOT NULL DEFAULT '{}',
    calculated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 3. device_registry - registered user devices
CREATE TABLE IF NOT EXISTS device_registry (
    device_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    device_name VARCHAR(255) NOT NULL,
    browser VARCHAR(100) NOT NULL DEFAULT '',
    os VARCHAR(100) NOT NULL DEFAULT '',
    platform VARCHAR(50) NOT NULL DEFAULT 'WEB',
    ip_address VARCHAR(45) NOT NULL DEFAULT '',
    location VARCHAR(100) NOT NULL DEFAULT 'Unknown',
    first_seen_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_seen_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    is_trusted BOOLEAN NOT NULL DEFAULT FALSE
);

-- 4. trusted_devices - trusted device whitelist log
CREATE TABLE IF NOT EXISTS trusted_devices (
    trust_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    device_id UUID NOT NULL REFERENCES device_registry(device_id) ON DELETE CASCADE,
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    trusted_by VARCHAR(100) NOT NULL DEFAULT 'USER',
    trusted_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 5. audit_events - comprehensive vault audit log hub
CREATE TABLE IF NOT EXISTS audit_events (
    audit_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    category VARCHAR(50) NOT NULL, -- AUTH, ASSET, SHARING, WORKFLOW, AI, SEARCH, BACKUP, VERSIONING, SETTINGS
    action VARCHAR(100) NOT NULL,
    actor_id VARCHAR(100) NOT NULL DEFAULT 'USER',
    actor_name VARCHAR(255) NOT NULL DEFAULT 'Vault Owner',
    ip_address VARCHAR(45) NOT NULL DEFAULT '',
    resource_id VARCHAR(255) NOT NULL DEFAULT '',
    details JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 6. risk_assessments - risk engine scoring assessments
CREATE TABLE IF NOT EXISTS risk_assessments (
    assessment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    risk_score INT NOT NULL DEFAULT 0, -- 0 to 100
    risk_level VARCHAR(20) NOT NULL DEFAULT 'LOW', -- LOW, MEDIUM, HIGH, CRITICAL
    signals JSONB NOT NULL DEFAULT '[]',
    assessed_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 7. security_recommendations - actionable security posture recommendations
CREATE TABLE IF NOT EXISTS security_recommendations (
    rec_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    category VARCHAR(50) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    action_type VARCHAR(100) NOT NULL DEFAULT 'CONFIGURE',
    priority VARCHAR(20) NOT NULL DEFAULT 'MEDIUM', -- LOW, MEDIUM, HIGH, URGENT
    is_dismissed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_audit_vault ON audit_events(vault_id);
CREATE INDEX IF NOT EXISTS idx_audit_category ON audit_events(category);
CREATE INDEX IF NOT EXISTS idx_device_vault ON device_registry(vault_id);
