-- Migration 015: Create Embedding Engine Tables & Vector Abstraction

-- 1. embeddings - core stored vector representations for knowledge items
CREATE TABLE IF NOT EXISTS embeddings (
    embedding_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    source_type VARCHAR(50) NOT NULL, -- DOCUMENT, SUMMARY, ENTITY, CONTEXT, RELATIONSHIP, MEMORY, COLLECTION, NOTE, CHAT
    source_id VARCHAR(255) NOT NULL,
    provider_name VARCHAR(50) NOT NULL DEFAULT 'default',
    model_name VARCHAR(100) NOT NULL DEFAULT 'text-embedding-3-small',
    model_version VARCHAR(50) NOT NULL DEFAULT 'v1',
    dimensions INT NOT NULL DEFAULT 1536,
    representation_text TEXT NOT NULL, -- Rich structured text representation
    content_hash VARCHAR(64) NOT NULL, -- SHA-256 for deduplication
    vector_data JSONB NOT NULL DEFAULT '[]', -- Vector floats array
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 2. embedding_chunks - intelligent chunked representations for large documents
CREATE TABLE IF NOT EXISTS embedding_chunks (
    chunk_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    embedding_id UUID NOT NULL REFERENCES embeddings(embedding_id) ON DELETE CASCADE,
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    chunk_index INT NOT NULL DEFAULT 0,
    chunk_strategy VARCHAR(50) NOT NULL DEFAULT 'SEMANTIC', -- SEMANTIC, HIERARCHICAL, DOCUMENT_AWARE, PAGE_AWARE, HEADING_AWARE
    chunk_text TEXT NOT NULL,
    token_count INT NOT NULL DEFAULT 0,
    heading_path VARCHAR(500) DEFAULT '',
    page_number INT DEFAULT 0,
    vector_data JSONB NOT NULL DEFAULT '[]',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 3. embedding_versions - model and provider upgrade version tracking
CREATE TABLE IF NOT EXISTS embedding_versions (
    version_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    provider_name VARCHAR(50) NOT NULL,
    model_name VARCHAR(100) NOT NULL,
    dimensions INT NOT NULL,
    version_tag VARCHAR(50) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    total_embedded_items INT NOT NULL DEFAULT 0,
    upgraded_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 4. embedding_jobs - batch and background job tracking
CREATE TABLE IF NOT EXISTS embedding_jobs (
    job_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    target_type VARCHAR(50) NOT NULL,
    target_id VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'QUEUED', -- QUEUED, PROCESSING, COMPLETED, FAILED
    error_message TEXT,
    processed_chunks INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMPTZ
);

-- 5. embedding_providers - active provider configurations per vault
CREATE TABLE IF NOT EXISTS embedding_providers (
    provider_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    provider_name VARCHAR(50) NOT NULL, -- OPENAI, GEMINI, VOYAGE, COHERE, OLLAMA, LOCAL
    model_name VARCHAR(100) NOT NULL,
    dimensions INT NOT NULL DEFAULT 1536,
    batch_size INT NOT NULL DEFAULT 32,
    is_default BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 6. vector_metadata - arbitrary key-value metadata attributes for vectors
CREATE TABLE IF NOT EXISTS vector_metadata (
    metadata_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    embedding_id UUID NOT NULL REFERENCES embeddings(embedding_id) ON DELETE CASCADE,
    meta_key VARCHAR(100) NOT NULL,
    meta_value TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_embeddings_vault ON embeddings(vault_id);
CREATE INDEX IF NOT EXISTS idx_embeddings_source ON embeddings(source_type, source_id);
CREATE INDEX IF NOT EXISTS idx_embeddings_hash ON embeddings(content_hash);
CREATE INDEX IF NOT EXISTS idx_embedding_chunks_embedding ON embedding_chunks(embedding_id);
CREATE INDEX IF NOT EXISTS idx_embedding_jobs_vault ON embedding_jobs(vault_id);
CREATE INDEX IF NOT EXISTS idx_vector_metadata_embedding ON vector_metadata(embedding_id);
