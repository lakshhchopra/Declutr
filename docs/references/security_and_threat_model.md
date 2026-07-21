# Declutr Security & Threat Model

> **Source of Truth:** [declutr_architecture_document.html](file:///f:/Github/Declutr/docs/architecture/declutr_architecture_document.html)  
> **Sections:** 6 (SRP Auth), 15 (Behavioral Auth), 21 (Security Model), 22 (Threat Model)

---

## 1. Zero-Knowledge Authentication (SRP-6a)

Declutr implements Secure Remote Password (SRP-6a) to verify user credentials without transmitting or storing plaintext passwords or pass-hashes.

```
Client (User)                                          Server (Backend)
  │                                                           │
  ├─── 1. Initiate (email_hash, A = g^a mod N) ──────────────>│
  │                                                           │ (Fetch Salt, Verifier v)
  │<── 2. Challenge (Salt, B = (kv + g^b) mod N) ─────────────┤
  │                                                           │
(Derive S, K, M1)                                           (Derive S, K, M1)
  │                                                           │
  ├─── 3. Verify (Client Proof M1) ──────────────────────────>│
  │                                                           │ (Verify M1)
  │<── 4. Token Issue (Server Proof M2, Access JWT) ──────────┤
```

- **Salt & Verifier:** Salt `s` and verifier `v = g^x mod N` stored in PostgreSQL.
- **Client Key Wrapping:** Password derives Master Encryption Key (MEK) to unwrap Master Vault Key (MVK). MVK unwraps Vault Key (VK).

---

## 2. Behavioral Risk Engine & Adaptive Security

Session risk is calculated passively and separately from personalization profiling:

```
[Session Signals (IP, Device Fingerprint, Access Velocity)]
                          │
                          ▼
              [Risk Evaluator Processor]
                          │
            ┌─────────────┴─────────────┐
            ▼                           ▼
   Score < 0.70                Score ≥ 0.70
 (Standard Access)       (Trigger Adaptive MFA Challenge)
```

---

## 3. STRIDE Threat Model & Mitigations

| Risk Vector | STRIDE Category | Risk Level | Architectural Mitigation |
| :--- | :--- | :--- | :--- |
| **Malicious File Injection** | Tampering / Elevation | Critical | Parse uploads inside isolated backend sandboxes (WASM/Docker runtimes) with strict memory limits and no network access. Sanitize extracted text before processing. |
| **Vector Embedding Leakage** | Information Disclosure | High | In Private Mode, client-encrypt vector embeddings and AI metadata. In Enhanced AI Mode, restrict database access via PostgreSQL Row-Level Security (RLS). |
| **Behavioral False Positives** | Denial of Service | Medium | Implement progressive grace periods and prompt with WebAuthn/Passkey challenges rather than executing instant session lockouts. |
| **Cryptographic Replay Attacks** | Spoofing / Replay | High | SRP challenges are single-use with a strict 5-minute expiration timestamp and cryptographic nonce verification. |
