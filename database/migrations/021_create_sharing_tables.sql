-- Migration 021: Create Secure Sharing & Collaboration Tables

-- 1. shares - core shared resource definition
CREATE TABLE IF NOT EXISTS shares (
    share_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    resource_type VARCHAR(50) NOT NULL, -- ASSET, FOLDER, COLLECTION, CONTEXT, PROJECT, TIMELINE_VIEW, SEARCH_RESULT
    resource_id VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    access_type VARCHAR(50) NOT NULL DEFAULT 'PRIVATE', -- PRIVATE, INVITE_ONLY, LINK_SHARING, TEMPORARY_ACCESS
    created_by VARCHAR(50) NOT NULL DEFAULT 'USER',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 2. share_permissions - role permission definitions
CREATE TABLE IF NOT EXISTS share_permissions (
    permission_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    share_id UUID NOT NULL REFERENCES shares(share_id) ON DELETE CASCADE,
    role VARCHAR(50) NOT NULL, -- READ_ONLY, COMMENT_ONLY, EDIT, OWNER, CO_OWNER
    can_view BOOLEAN NOT NULL DEFAULT TRUE,
    can_download BOOLEAN NOT NULL DEFAULT TRUE,
    can_edit BOOLEAN NOT NULL DEFAULT FALSE,
    can_delete BOOLEAN NOT NULL DEFAULT FALSE,
    can_comment BOOLEAN NOT NULL DEFAULT TRUE,
    can_share BOOLEAN NOT NULL DEFAULT FALSE,
    can_manage_members BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 3. share_members - member membership and roles
CREATE TABLE IF NOT EXISTS share_members (
    member_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    share_id UUID NOT NULL REFERENCES shares(share_id) ON DELETE CASCADE,
    user_id VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'READ_ONLY',
    joined_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ
);

-- 4. share_links - link sharing tokens and rules
CREATE TABLE IF NOT EXISTS share_links (
    link_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    share_id UUID NOT NULL REFERENCES shares(share_id) ON DELETE CASCADE,
    link_token VARCHAR(255) NOT NULL UNIQUE,
    is_password_protected BOOLEAN NOT NULL DEFAULT FALSE,
    password_hash VARCHAR(255) NOT NULL DEFAULT '',
    disable_download BOOLEAN NOT NULL DEFAULT FALSE,
    disable_reshare BOOLEAN NOT NULL DEFAULT TRUE,
    view_count INT NOT NULL DEFAULT 0,
    max_views INT NOT NULL DEFAULT 0, -- 0 = unlimited
    download_count INT NOT NULL DEFAULT 0,
    max_downloads INT NOT NULL DEFAULT 0,
    expires_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 5. share_comments - threaded resource comments
CREATE TABLE IF NOT EXISTS share_comments (
    comment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    share_id UUID NOT NULL REFERENCES shares(share_id) ON DELETE CASCADE,
    user_id VARCHAR(100) NOT NULL,
    user_name VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    parent_comment_id UUID REFERENCES share_comments(comment_id) ON DELETE CASCADE,
    is_resolved BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 6. share_activity - audit log of collaboration events
CREATE TABLE IF NOT EXISTS share_activity (
    activity_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    share_id UUID NOT NULL REFERENCES shares(share_id) ON DELETE CASCADE,
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    actor_id VARCHAR(100) NOT NULL,
    actor_name VARCHAR(255) NOT NULL,
    action_type VARCHAR(50) NOT NULL, -- VIEWED, DOWNLOADED, EDITED, COMMENTED, SHARED, PERMISSION_CHANGED, ACCESS_REVOKED, INVITE_ACCEPTED
    details JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 7. share_invitations - pending invitations
CREATE TABLE IF NOT EXISTS share_invitations (
    invitation_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    share_id UUID NOT NULL REFERENCES shares(share_id) ON DELETE CASCADE,
    inviter_id VARCHAR(100) NOT NULL,
    invitee_email VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'READ_ONLY',
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING', -- PENDING, ACCEPTED, REJECTED, REVOKED
    token VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_shares_vault ON shares(vault_id);
CREATE INDEX IF NOT EXISTS idx_share_members_user ON share_members(user_id);
CREATE INDEX IF NOT EXISTS idx_share_comments_share ON share_comments(share_id);
