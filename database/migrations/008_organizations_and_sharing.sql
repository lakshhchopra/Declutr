-- Consolidated Migration 008: Enterprise Organizations & Sharing
CREATE TABLE IF NOT EXISTS organizations (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    domain VARCHAR(255) UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS org_members (
    organization_id VARCHAR(255) NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    user_id VARCHAR(255) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role VARCHAR(50) NOT NULL DEFAULT 'MEMBER',
    PRIMARY KEY (organization_id, user_id)
);

CREATE TABLE IF NOT EXISTS shares (
    id VARCHAR(255) PRIMARY KEY,
    asset_id VARCHAR(255) NOT NULL REFERENCES assets(id) ON DELETE CASCADE,
    shared_by_user_id VARCHAR(255) NOT NULL,
    recipient_email VARCHAR(255) NOT NULL,
    permission VARCHAR(50) NOT NULL DEFAULT 'VIEW',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
