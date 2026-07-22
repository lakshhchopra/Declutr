-- Migration 016: Create Hybrid Knowledge Search Tables

-- 1. search_history - audit log of user search queries and performance
CREATE TABLE IF NOT EXISTS search_history (
    history_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    query_text TEXT NOT NULL,
    parsed_query JSONB NOT NULL DEFAULT '{}',
    result_count INT NOT NULL DEFAULT 0,
    latency_ms INT NOT NULL DEFAULT 0,
    search_type VARCHAR(50) NOT NULL DEFAULT 'HYBRID', -- HYBRID, KEYWORD, VECTOR, ENTITY, CONTEXT, MEMORY
    searched_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 2. saved_searches - user pinned and bookmarked search queries with filters
CREATE TABLE IF NOT EXISTS saved_searches (
    saved_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    search_name VARCHAR(255) NOT NULL,
    query_text TEXT NOT NULL,
    filters JSONB NOT NULL DEFAULT '{}',
    is_pinned BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 3. search_statistics - vault-level query analytics and top terms
CREATE TABLE IF NOT EXISTS search_statistics (
    stats_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE UNIQUE,
    total_searches INT NOT NULL DEFAULT 0,
    top_queries JSONB NOT NULL DEFAULT '[]',
    avg_latency_ms FLOAT NOT NULL DEFAULT 0.0,
    strategy_usage JSONB NOT NULL DEFAULT '{}',
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 4. search_preferences - per-vault ranking weights and search defaults
CREATE TABLE IF NOT EXISTS search_preferences (
    preference_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE UNIQUE,
    ranking_weights JSONB NOT NULL DEFAULT '{
        "keyword": 0.25,
        "vector": 0.25,
        "entity": 0.15,
        "context": 0.15,
        "memory": 0.10,
        "recency": 0.10
    }',
    enable_auto_suggestions BOOLEAN NOT NULL DEFAULT TRUE,
    max_results_per_page INT NOT NULL DEFAULT 20,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 5. search_index_versions - tracks search index sync status
CREATE TABLE IF NOT EXISTS search_index_versions (
    index_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    fts_version VARCHAR(50) NOT NULL DEFAULT 'v1',
    vector_version VARCHAR(50) NOT NULL DEFAULT 'v1',
    last_indexed_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_search_history_vault ON search_history(vault_id);
CREATE INDEX IF NOT EXISTS idx_saved_searches_vault ON saved_searches(vault_id);
