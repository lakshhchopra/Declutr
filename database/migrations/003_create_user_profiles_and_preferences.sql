-- Migration 003: Create User Profiles and Preferences Tables
CREATE TABLE IF NOT EXISTS user_profiles (
    user_id UUID PRIMARY KEY REFERENCES users(user_id) ON DELETE CASCADE,
    display_name VARCHAR(100),
    avatar_url TEXT,
    timezone VARCHAR(50) DEFAULT 'UTC',
    language VARCHAR(10) DEFAULT 'en',
    region VARCHAR(10) DEFAULT 'US',
    bio TEXT,
    onboarding_completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_preferences (
    user_id UUID PRIMARY KEY REFERENCES users(user_id) ON DELETE CASCADE,
    theme VARCHAR(20) DEFAULT 'dark',
    ai_behavior VARCHAR(30) DEFAULT 'balanced', -- 'minimal', 'balanced', 'proactive'
    privacy_mode VARCHAR(30) DEFAULT 'local_first', -- 'local_first', 'encrypted_cloud', 'hybrid'
    email_notifications BOOLEAN DEFAULT TRUE,
    push_notifications BOOLEAN DEFAULT TRUE,
    desktop_notifications BOOLEAN DEFAULT TRUE,
    security_alerts BOOLEAN DEFAULT TRUE,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
