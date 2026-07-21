# ADR 0003: Unified Storage with PostgreSQL 16 + pgvector

## Context
Declutr requires relational database guarantees (foreign keys, row-level security) alongside high-dimensional vector similarity search for hybrid retrieval.

## Decision
Use **PostgreSQL 16 with the `pgvector` extension** as the single primary datastore, consolidating relational metadata, session tokens, and 512-dimensional semantic embeddings. Dedicated graph databases (Neo4j) or separate vector databases (Qdrant/Pinecone) are excluded from MVP.

## Status
Accepted

## Consequences
- Single datastore simplifies transactions and backups.
- Native Row-Level Security (RLS) restricts access contextually per user.
- Cost-effective scaling with HNSW vector indexing.
