# Declutr Server Knowledge Boundary

## Server CAN Know

- User ID
- Email
- Encrypted Master Vault Key (MVK)
- Encrypted Vault Keys
- Encrypted File Keys
- Encrypted Metadata
- Encrypted Files
- Timestamps
- Storage Paths
- Blind Index Tokens

## Server MUST NEVER Know

- User Password
- Master Encryption Key (MEK)
- Master Vault Key plaintext
- Vault Key plaintext
- File Key plaintext
- File contents
- Decrypted metadata

## Breach Scenario

If the database is compromised, an attacker should obtain:

- Ciphertext
- Blind indexes
- Timestamps

But never:

- Passwords
- MEK
- MVK plaintext
- User files in plaintext

