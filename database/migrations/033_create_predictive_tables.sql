-- Migration: 033_create_predictive_tables.sql
-- Description: Predictive Intelligence & Life Intelligence Engine Database Tables

-- 1. Proactive Predictions Table
CREATE TABLE IF NOT EXISTS predictions (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    confidence NUMERIC(3,2) NOT NULL DEFAULT 0.85,
    priority VARCHAR(20) NOT NULL DEFAULT 'MEDIUM', -- HIGH, MEDIUM, LOW
    evidence JSONB NOT NULL DEFAULT '{}'::jsonb,
    affected_assets JSONB NOT NULL DEFAULT '[]'::jsonb,
    suggested_action TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING', -- PENDING, ACCEPTED, DISMISSED, EXPIRED
    expiration TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_predictions_user ON predictions (user_id);
CREATE INDEX IF NOT EXISTS idx_predictions_type ON predictions (type);
CREATE INDEX IF NOT EXISTS idx_predictions_status ON predictions (status);

-- 2. Prediction History Table
CREATE TABLE IF NOT EXISTS prediction_history (
    id VARCHAR(255) PRIMARY KEY,
    prediction_id VARCHAR(255) NOT NULL REFERENCES predictions(id) ON DELETE CASCADE,
    user_id VARCHAR(255) NOT NULL,
    action VARCHAR(50) NOT NULL,
    executed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 3. Prediction Feedback & Tuning Table
CREATE TABLE IF NOT EXISTS prediction_feedback (
    id VARCHAR(255) PRIMARY KEY,
    prediction_id VARCHAR(255) NOT NULL REFERENCES predictions(id) ON DELETE CASCADE,
    user_id VARCHAR(255) NOT NULL,
    action VARCHAR(50) NOT NULL, -- ACCEPTED, DISMISSED
    reason TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 4. Predictive Machine Learning Models Table
CREATE TABLE IF NOT EXISTS prediction_models (
    id VARCHAR(255) PRIMARY KEY,
    model_name VARCHAR(100) NOT NULL,
    version VARCHAR(50) NOT NULL,
    accuracy_score NUMERIC(3,2) DEFAULT 0.95,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 5. Prediction Scores Telemetry Table
CREATE TABLE IF NOT EXISTS prediction_scores (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    total_generated INT DEFAULT 0,
    accepted_count INT DEFAULT 0,
    dismissed_count INT DEFAULT 0,
    acceptance_rate NUMERIC(3,2) DEFAULT 0.0,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
