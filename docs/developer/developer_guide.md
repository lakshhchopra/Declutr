# Declutr Public API & Developer Platform Guide

Declutr provides a versioned, secure REST API platform enabling developers, automations, and AI agents to build applications on top of Declutr.

## Authentication Modes

1. **Scoped API Keys**:
   - Format: `declutr_live_...`
   - Header: `Authorization: Bearer <API_KEY>`
   - Granular Scopes: `vault.read`, `vault.write`, `asset.read`, `asset.write`, `workflow.execute`, `ai.chat`, `search.query`, `backup.manage`, `admin.manage`.

2. **OAuth 2.1 PKCE**:
   - Standard PKCE Authorization Code Grant for public and confidential third-party client apps.

3. **Personal Access Tokens (PATs) & Service Accounts**:
   - For automated CI/CD and system integrations.

## API Versioning

All public endpoints are versioned under `/api/v1/`. Future major versions will coexist without breaking existing clients.
