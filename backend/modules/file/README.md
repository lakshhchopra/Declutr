# File Module

## Responsibility
Manages metadata properties of digital items, file upload allocations, S3 pre-signed link generation, and chunk composition.

## Module Boundaries
- Domain: Defines File and Version entities.
- Application: Orchestrates direct-to-S3 uploads and commit verification.
- Repository: Saves storage paths and versions in database.
- Transport: Exposes upload and commit APIs.
