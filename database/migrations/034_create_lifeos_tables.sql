-- Migration: 034_create_lifeos_tables.sql
-- Description: Life Operating System (LifeOS) Database Tables

-- 1. Life Areas Table
CREATE TABLE IF NOT EXISTS life_areas (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    icon VARCHAR(100) DEFAULT 'Compass',
    color VARCHAR(50) DEFAULT '#6366F1',
    is_custom BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_life_areas_user ON life_areas (user_id);

-- 2. Projects Master Table
CREATE TABLE IF NOT EXISTS projects (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    life_area_id VARCHAR(255) NOT NULL REFERENCES life_areas(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'IN_PROGRESS', -- PLANNING, IN_PROGRESS, BLOCKED, COMPLETED, ARCHIVED
    budget NUMERIC(12,2) DEFAULT 0.0,
    associated_docs JSONB NOT NULL DEFAULT '[]'::jsonb,
    people JSONB NOT NULL DEFAULT '[]'::jsonb,
    target_date TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_projects_area ON projects (life_area_id);
CREATE INDEX IF NOT EXISTS idx_projects_user ON projects (user_id);

-- 3. Project Goals Table
CREATE TABLE IF NOT EXISTS project_goals (
    id VARCHAR(255) PRIMARY KEY,
    project_id VARCHAR(255) NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    user_id VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    progress_pct INT NOT NULL DEFAULT 0,
    is_completed BOOLEAN DEFAULT false,
    missing_assets JSONB NOT NULL DEFAULT '[]'::jsonb,
    due_date TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_project_goals_prj ON project_goals (project_id);

-- 4. Goal Progress Log Table
CREATE TABLE IF NOT EXISTS goal_progress (
    id VARCHAR(255) PRIMARY KEY,
    goal_id VARCHAR(255) NOT NULL REFERENCES project_goals(id) ON DELETE CASCADE,
    progress_pct INT NOT NULL,
    comment TEXT,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 5. LifeOS Unified Dashboard Telemetry Table
CREATE TABLE IF NOT EXISTS life_dashboard (
    user_id VARCHAR(255) PRIMARY KEY,
    health_score INT DEFAULT 90,
    active_projects_count INT DEFAULT 0,
    active_goals_count INT DEFAULT 0,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 6. Life Balance Metrics Telemetry Table
CREATE TABLE IF NOT EXISTS life_metrics (
    user_id VARCHAR(255) PRIMARY KEY,
    active_projects INT DEFAULT 0,
    goal_completion NUMERIC(3,2) DEFAULT 0.0,
    life_balance INT DEFAULT 85,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
