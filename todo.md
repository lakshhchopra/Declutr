# Declutr Project To-Do List

This document tracks the comprehensive roadmap and action items to complete the Declutr project, including the Go backend, Next.js web application, and React Native/Expo mobile application.

---

## 🛠️ Phase 0: Project Setup & Synchronization
- [x] Scaffold Go backend workspace in [/backend](file:///f:/Github/Declutr/backend)
- [x] Scaffold Next.js web application in [/frontend](file:///f:/Github/Declutr/frontend)
- [x] Scaffold React Native/Expo mobile application in [frontend/declutr-mobile](file:///f:/Github/Declutr/frontend/declutr-mobile)
- [x] Refactor repository into a domain-oriented modular monorepo (`modules/`, `shared/`, `platform/`, `features/`, `docs/`, `infrastructure/`, `scripts/`, `tests/`, `security/`)
- [x] Initialize PostgreSQL database connection wrapper
- [x] Setup health check endpoints and routing configurations
- [ ] Set up Docker Compose for local development (Go API + PostgreSQL + pgvector + Redis)
- [ ] Establish CI/CD structures (GitHub Actions, linting, testing checks)

---

## 🔒 Phase 1: Authentication & Identity Foundation
- [x] Implement user email hashing (Argon2id) before persistence
- [x] Define SRP-6a domain types and protocol state models
- [x] Implement single-use SRP challenge generation & expiration validation
- [x] Add SRP proof verification interface & engine foundation
- [x] Add session model and database table
- [x] Create session token generator (JWT-like or secure random byte tokens)
- [ ] Complete SRP-6a Authentication flow endpoints on Go API:
  - `POST /v1/auth/srp/initiate`
  - `POST /v1/auth/srp/verify`
- [ ] Implement JWT Session rotation and Refresh token flow
- [ ] Web Client Integration (Next.js):
  - Integrate client-side SRP calculation (WASM-compiled `libsodium` or WebCrypto)
  - Implement credentials caching and Master Vault Key wrapping
- [ ] Mobile Client Integration (Expo / React Native):
  - Integrate native cryptographic library for SRP calculations
  - Implement secure local storage (e.g. `expo-secure-store`) for session tokens and key wrapping
- [ ] Integrate WebAuthn / Passkey setup and validation flow (both Web and Native Mobile)

---

## 📦 Phase 2: Cryptographic Vaults & Direct Ingestion
- [ ] Database Schema: Vaults and Digital Items
  - Implement migration for `vaults`, `digital_items`, and `item_versions` tables
  - Add Row-Level Security (RLS) policies on PostgreSQL for user-isolation
- [ ] Vault Management services on Go API:
  - Vault creation (`POST /v1/vaults`) and key wrapping handling (wrapping VK with MVK)
- [ ] Client-Side AES-256-GCM encryption utilities for files (both Next.js and React Native)
- [ ] Direct S3/Cloudflare R2 Chunked Upload implementation:
  - Backend pre-signed URL generator API (`POST /v1/files/upload/initiate`)
  - Web & Mobile multipart chunked upload handlers with auto-resume support
  - Upload commit endpoint (`POST /v1/files/upload/commit`)

---

## 🧠 Phase 3: Background Jobs & Ingestion Pipeline
- [ ] Set up Go background worker framework (e.g. `Asynq` or `River` backed by Redis)
- [ ] Implement capability-driven Content Ingestion Pipeline:
  - File Validation & mime-type magic-number checks
  - Content parsing (plaintext extraction, PDF structures)
  - OCR extraction service (Tesseract wrapper, local ONNX OCR model, or cloud API wrapper)
  - Audio transcription service (Whisper execution)
- [ ] Ingestion job state machine monitoring (`UPLOADED` ➔ `QUEUED` ➔ `PROCESSING` ➔ `COMPLETED`/`FAILED`)

---

## 🔗 Phase 4: AI Context, Intent & Relationships
- [ ] Database Schema: Metadata & Relational tables
  - Add tables for `item_metadata`, `ai_metadata`, and `relationships`
- [ ] Entity Extraction: Parser targeting dates, locations, merchants, transaction values, and names
- [ ] Intent-Aware Organization Engine:
  - Probabilistic classification of files into Category vs. Intent vs. Context
  - Direct relationship modeling: `RELATED_TO`, `PART_OF`, `MENTIONS`, `SAME_EVENT`, `SAME_LOCATION`
- [ ] User Feedback loop:
  - API endpoint `POST /v1/feedback/verify` to confirm or correct AI-generated metadata/tags

---

## 🔍 Phase 5: Semantic Retrieval & Persona Intelligence
- [ ] Add `pgvector` extension support to PostgreSQL migration
- [ ] Build vector embedding table (`embeddings`)
- [ ] Embeddings generation pipeline:
  - Connect text chunks to vector models (e.g. local ONNX embeddings or cloud APIs)
- [ ] Hybrid Search query processor (`POST /v1/search/query`):
  - Combine traditional PostgreSQL full-text search (keyword) and pgvector semantic distance
- [ ] Reverse Persona Engine:
  - Collect user interaction signals (clicks, search terms, category usage)
  - Build personalization profile with time-based recency decay
  - Utilize persona parameters to rank search results contextually

---

## 🛡️ Phase 6: Behavioral Security, Hardening & Audit
- [ ] Behavioral Authentication Engine:
  - Passive session signal collector (IP subnet, client device fingerprint, navigation anomalies)
  - Risk evaluator model producing real-time session scores
- [ ] Adaptive security interceptors:
  - Trigger MFA/Passkey challenge when session risk score exceeds threshold
- [ ] Cryptographic Audit Trail:
  - HMAC-chained append-only database transaction log for critical user actions
- [ ] Isolated file parsing sandbox (executing extraction engines in isolated Docker/WASM layers)
- [ ] Penetration testing and vulnerability scanning

---

## 🚀 Phase 7: Deployment & Optimization
- [ ] Performance Optimizations:
  - PostgreSQL indexes and pgvector HNSW index configurations
  - Redis cache policy setups for metadata querying
- [ ] Production-ready Dockerfiles & Helm charts
- [ ] Deploy staging environment
- [ ] Expo/Mobile app build profiles (eas build setup for iOS and Android)
- [ ] Final end-to-end system verification
