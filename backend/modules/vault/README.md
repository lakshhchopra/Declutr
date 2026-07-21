# Vault Module

## Responsibility
Manages physical vault scopes, shared crypt-keys, zero-knowledge vault boundaries, and folder hierarchies.

## Module Boundaries
- Domain: Defines Vault entities and key wrapping types.
- Application: Encapsulates vault CRUD orchestration.
- Repository: Handles PostgreSQL persistence for vault records.
- Transport: Exposes REST APIs for vault management.
