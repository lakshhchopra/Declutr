# Declutr Authentication Flow

## Registration

User Password
↓
Argon2id
↓
Master Encryption Key (MEK)
↓
Generate Master Vault Key (MVK)
↓
Encrypt MVK with MEK
↓
Store encrypted MVK on server

Server never sees:
- Password
- MEK
- MVK plaintext

## Login

User Password
↓
Argon2id
↓
Recover MEK
↓
Download encrypted MVK
↓
Decrypt MVK locally

Server never sees:
- Password
- MEK
- MVK plaintext

## Future

- SRP Authentication
- Passkeys
- MFA
