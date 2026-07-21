-- Migration 009: Create AI Analysis Tables

CREATE TABLE IF NOT EXISTS ai_analysis (
    analysis_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    document_id UUID NOT NULL REFERENCES extracted_documents(document_id) ON DELETE CASCADE,
    asset_id UUID NOT NULL REFERENCES assets(asset_id) ON DELETE CASCADE,
    title TEXT,
    short_summary TEXT,
    detailed_summary TEXT,
    language VARCHAR(50),
    writing_style VARCHAR(100),
    sentiment VARCHAR(50),
    complexity VARCHAR(50),
    reading_level VARCHAR(50),
    estimated_reading_time INT,
    document_purpose TEXT,
    confidence_score FLOAT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS ai_classification (
    classification_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    analysis_id UUID NOT NULL REFERENCES ai_analysis(analysis_id) ON DELETE CASCADE,
    document_category VARCHAR(100) NOT NULL, -- e.g., 'Receipt', 'Contract'
    document_type VARCHAR(100),
    is_scanned BOOLEAN DEFAULT false,
    is_corrupted BOOLEAN DEFAULT false,
    is_incomplete BOOLEAN DEFAULT false,
    quality_score FLOAT,
    confidence_score FLOAT
);

CREATE TABLE IF NOT EXISTS ai_tags (
    tag_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    analysis_id UUID NOT NULL REFERENCES ai_analysis(analysis_id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    confidence_score FLOAT
);

CREATE TABLE IF NOT EXISTS ai_topics (
    topic_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    analysis_id UUID NOT NULL REFERENCES ai_analysis(analysis_id) ON DELETE CASCADE,
    name VARCHAR(200) NOT NULL,
    confidence_score FLOAT
);

CREATE TABLE IF NOT EXISTS analysis_versions (
    version_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    analysis_id UUID NOT NULL REFERENCES ai_analysis(analysis_id) ON DELETE CASCADE,
    provider VARCHAR(50) NOT NULL, -- 'openai', 'anthropic', 'mock'
    model_name VARCHAR(100) NOT NULL,
    prompt_version VARCHAR(50) NOT NULL,
    prompt_tokens INT DEFAULT 0,
    completion_tokens INT DEFAULT 0,
    cost_usd DOUBLE PRECISION DEFAULT 0.0,
    latency_ms INT DEFAULT 0,
    raw_output JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_ai_analysis_doc_id ON ai_analysis(document_id);
CREATE INDEX IF NOT EXISTS idx_ai_tags_name ON ai_tags(name);
CREATE INDEX IF NOT EXISTS idx_ai_classification_cat ON ai_classification(document_category);
