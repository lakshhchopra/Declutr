-- Migration 010: Create Entity Extraction Tables

CREATE TABLE IF NOT EXISTS entity_types (
    entity_type VARCHAR(50) PRIMARY KEY,
    description TEXT
);

INSERT INTO entity_types (entity_type, description) VALUES
('Person', 'A human being'),
('Organization', 'A company, institution, or group'),
('Location', 'A physical place, city, country, or address'),
('Date', 'A specific date or time period'),
('Amount', 'A monetary value or currency'),
('Product', 'A commercial product or brand'),
('Identifier', 'A passport number, vehicle registration, etc.')
ON CONFLICT DO NOTHING;

CREATE TABLE IF NOT EXISTS entities (
    entity_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    entity_type VARCHAR(50) NOT NULL REFERENCES entity_types(entity_type),
    canonical_name VARCHAR(255) NOT NULL,
    normalized_value VARCHAR(255),
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(vault_id, entity_type, canonical_name)
);

CREATE TABLE IF NOT EXISTS entity_aliases (
    alias_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    entity_id UUID NOT NULL REFERENCES entities(entity_id) ON DELETE CASCADE,
    alias_name VARCHAR(255) NOT NULL,
    UNIQUE(entity_id, alias_name)
);

CREATE TABLE IF NOT EXISTS entity_occurrences (
    occurrence_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    entity_id UUID NOT NULL REFERENCES entities(entity_id) ON DELETE CASCADE,
    asset_id UUID NOT NULL REFERENCES assets(asset_id) ON DELETE CASCADE,
    analysis_id UUID REFERENCES ai_analysis(analysis_id) ON DELETE SET NULL,
    original_value VARCHAR(255) NOT NULL,
    confidence_score FLOAT NOT NULL DEFAULT 1.0,
    extractor_version VARCHAR(50),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_entities_vault ON entities(vault_id);
CREATE INDEX IF NOT EXISTS idx_entities_canonical ON entities(canonical_name);
CREATE INDEX IF NOT EXISTS idx_entity_aliases_name ON entity_aliases(alias_name);
CREATE INDEX IF NOT EXISTS idx_entity_occurrences_asset ON entity_occurrences(asset_id);
