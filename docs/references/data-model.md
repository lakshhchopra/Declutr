# Declutr Data Model

## User

- id (UUID)
- email
- encrypted_master_key
- created_at
- updated_at

## Vault

- id (UUID)
- owner_id
- encrypted_vault_key
- vault_name
- created_at
- updated_at

## File

- id (UUID)
- vault_id
- encrypted_file_key
- encrypted_metadata
- file_size
- mime_type
- storage_path
- created_at

## File Version

- id (UUID)
- file_id
- encrypted_file_key
- storage_path
- version_number
- created_at
