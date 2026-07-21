# ADR 0002: Secure Remote Password (SRP-6a) Zero-Knowledge Auth

## Context
Standard password hashing (Bcrypt/Argon2id sent over wire) exposes user credentials to server memory during login. A zero-knowledge protocol ensures the backend never learns or handles the plaintext password.

## Decision
Implement **SRP-6a** as the default zero-knowledge authentication mechanism, combined with WebAuthn/Passkeys for biometric passwordless sign-in.

## Status
Accepted

## Consequences
- Passwords are never sent to the server.
- Server database stores salt and verifier `v = g^x mod N`.
- Protects vault master keys from server credential leaks.
