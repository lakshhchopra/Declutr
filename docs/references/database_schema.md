# Declutr Database Schema & Architecture

> **Source of Truth:** [declutr_architecture_document.html](file:///f:/Github/Declutr/docs/architecture/declutr_architecture_document.html)  
> **Section:** 18. Database Architecture & Schema

---

## PostgreSQL & pgvector Schema Layout

Declutr isolates user accounts, cryptographic vaults, digital items, AI extractions, vector coordinates, and relationships across PostgreSQL tables.

```
       +-------------------+             +-------------------+
       |       users       | 1         * |      vaults       |
       +-------------------+-------------+-------------------+
       | user_id (PK)      |             | vault_id (PK)     |
       | email_hash        |             | owner_user_id(FK) |
       | srp_verifier      |             | encrypted_vk      |
       | srp_salt          |             +---------+---------+
       | encrypted_mvk     |                       | 1
       +-------------------+                       |
                                                   | *
                                         +---------v---------+
                                         |   digital_items   |
                                         +-------------------+
                                         | item_id (PK)      |
                                         | vault_id (FK)     |
                                         | active_version_id |
                                         +----+----+----+----+
                                              |    |    |
                   +--------------------------+    |    +-------------------------+
                   | 1                             | 1                            | 1
                   v *                             v *                            v *
         +-------------------+           +-------------------+          +-------------------+
         |   item_metadata   |           |    ai_metadata    |          |    embeddings     |
         +-------------------+           +-------------------+          +-------------------+
         | metadata_id (PK)  |           | ai_meta_id (PK)   |          | embedding_id (PK) |
         | item_id (FK)      |           | item_id (FK)      |          | item_id (FK)      |
         | encrypted_payload |           | confidence        |          | vector_data(512)  |
         +-------------------+           | encrypted_outputs |          +-------------------+
                                         +-------------------+
```

---

## Entity Tables

### 1. `users` (Account Registry)
```sql
CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email_hash BYTEA NOT NULL UNIQUE,
    srp_salt BYTEA NOT NULL,
    srp_verifier BYTEA NOT NULL,
    encrypted_mvk BYTEA NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

### 2. `vaults` (Cryptographic Workspaces)
```sql
CREATE TABLE vaults (
    vault_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    encrypted_vault_key BYTEA NOT NULL,
    privacy_mode TEXT NOT NULL DEFAULT 'private', -- 'private' | 'enhanced_ai'
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

### 3. `digital_items` (Atomic Assets)
```sql
CREATE TABLE digital_items (
    item_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    active_version_id UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

### 4. `item_metadata` (Server-Opaque Metadata)
```sql
CREATE TABLE item_metadata (
    metadata_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    item_id UUID NOT NULL REFERENCES digital_items(item_id) ON DELETE CASCADE,
    encrypted_payload BYTEA NOT NULL
);
```

### 5. `ai_metadata` (Extraction Records)
```sql
CREATE TABLE ai_metadata (
    ai_meta_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    item_id UUID NOT NULL REFERENCES digital_items(item_id) ON DELETE CASCADE,
    model_name TEXT NOT NULL,
    model_version TEXT NOT NULL,
    confidence NUMERIC(3,2) NOT NULL,
    user_status TEXT NOT NULL DEFAULT 'unreviewed', -- 'unreviewed'|'accepted'|'corrected'|'rejected'
    encrypted_outputs BYTEA NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

### 6. `embeddings` (pgvector Table)
```sql
CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE embeddings (
    embedding_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    item_id UUID NOT NULL REFERENCES digital_items(item_id) ON DELETE CASCADE,
    vector_data vector(512),
    encrypted_vector_data BYTEA
);
```

### 7. `relationships` (Inferred Relational Graph)
```sql
CREATE TABLE relationships (
    relationship_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    source_item_id UUID NOT NULL REFERENCES digital_items(item_id) ON DELETE CASCADE,
    target_item_id UUID NOT NULL REFERENCES digital_items(item_id) ON DELETE CASCADE,
    relationship_type TEXT NOT NULL, -- 'RELATED_TO', 'PART_OF', 'MENTIONS', 'SAME_EVENT', 'SAME_LOCATION'
    confidence NUMERIC(3,2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```
