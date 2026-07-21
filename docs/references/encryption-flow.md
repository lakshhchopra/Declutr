# Declutr Encryption Flow

Password
↓
Argon2id
↓
Master Encryption Key (MEK)
↓
Decrypt Master Vault Key (MVK)
↓
Decrypt Vault Key (VK)
↓
Decrypt File Key (FK)
↓
Decrypt File

## Upload Flow

File
↓
Generate File Key (FK)
↓
AES-256-GCM Encrypt File
↓
Encrypt FK with Vault Key
↓
Upload Ciphertext

## Password Change

Old Password
↓
Decrypt MVK
↓
New Password
↓
Re-wrap MVK

No file re-encryption required.
