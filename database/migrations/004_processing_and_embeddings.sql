-- Consolidated Migration 004: Processing & Embeddings
CREATE TABLE IF NOT EXISTS asset_processing_jobs (
    id VARCHAR(255) PRIMARY KEY,
    asset_id VARCHAR(255) NOT NULL REFERENCES assets(id) ON DELETE CASCADE,
    job_type VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    extracted_text TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS vector_embeddings (
    id VARCHAR(255) PRIMARY KEY,
    asset_id VARCHAR(255) NOT NULL REFERENCES assets(id) ON DELETE CASCADE,
    vector_data JSONB NOT NULL,
    dimensions INT NOT NULL DEFAULT 1536,
    model_name VARCHAR(100) DEFAULT 'text-embedding-3-small',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
