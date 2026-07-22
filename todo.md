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
- [x] Build Shared Design System Foundation (`shadcn/ui`, Radix UI, Dark Mode default, Theme tokens)
- [x] Build Application Shell & Navigation Foundation (Responsive AppShell, Route Placeholders, PageShell, Global Providers)
- [ ] Set up Docker Compose for local development (Go API + PostgreSQL + pgvector + Redis)
- [ ] Establish CI/CD structures (GitHub Actions, linting, testing checks)

---

## 🔒 Phase 1: Authentication & Identity Foundation
- [x] Build Authentication UI & Onboarding Experience (Welcome, Sign In, Register, Forgot Password, Reset Password, Verify Email, Magic Link, Auth Error views)
- [x] Implement user email hashing (Argon2id) before persistence
- [x] Define SRP-6a domain types and protocol state models
- [x] Implement single-use SRP challenge generation & expiration validation
- [x] Add SRP proof verification interface & engine foundation
- [x] Add session model and database table
- [x] Create session token generator (JWT-like or secure random byte tokens)
- [x] Complete SRP-6a Authentication flow endpoints on Go API:
  - `POST /api/v1/auth/register`
  - `POST /api/v1/auth/login/start`
  - `POST /api/v1/auth/login/finish`
- [x] Implement JWT Session rotation and Refresh token flow
- [x] Web Client Integration (Next.js):
  - Integrate client-side SRP calculation & payload exchange (`AuthService`, TanStack Query mutations)
  - Connect login and registration forms to Go SRP backend APIs
- [ ] Mobile Client Integration (Expo / React Native):
  - Integrate native cryptographic library for SRP calculations
  - Implement secure local storage (e.g. `expo-secure-store`) for session tokens and key wrapping
- [x] Implement User Onboarding, Profile & Preferences (Interactive 8-step `/onboarding`, tabbed `/settings`, `ProfileService`, `003_create_user_profiles_and_preferences.sql`)
- [ ] Integrate WebAuthn / Passkey setup and validation flow (both Web and Native Mobile)

---

## 📦 Phase 2: Cryptographic Vaults & Direct Ingestion
- [x] Vault Workspace Foundation (Default "My Life Vault" creation, `vaults` database table, `VaultService`, Vault overview UI)
- [ ] Database Schema: Vaults and Digital Items
  - Implement migration for `vaults`, `digital_items`, and `item_versions` tables
  - Add Row-Level Security (RLS) policies on PostgreSQL for user-isolation
- [ ] Vault Management services on Go API:
  - Vault creation (`POST /v1/vaults`) and key wrapping handling (wrapping VK with MVK)
- [ ] Client-Side AES-256-GCM encryption utilities for files (both Next.js and React Native)
- [x] Content Ingestion & Upload Pipeline (`StorageProvider` abstraction, `005_create_assets_and_ingestion_tables.sql`, `UploadModal` drag & drop, SHA-256 WebCrypto checksums)
- [ ] Direct S3/Cloudflare R2 Chunked Upload implementation:
  - Backend pre-signed URL generator API (`POST /v1/files/upload/initiate`)
  - Web & Mobile multipart chunked upload handlers with auto-resume support
  - Upload commit endpoint (`POST /v1/files/upload/commit`)

---

## 🧠 Phase 3: Background Jobs & Ingestion Pipeline
- [x] Set up Go background worker framework (Processing Engine, Queue interfaces, Scheduler)
- [x] Ingestion job state machine monitoring (`QUEUED` ➔ `RUNNING` ➔ `COMPLETED`/`FAILED`)
- [x] Implement capability-driven Content Ingestion Pipeline:
  - [x] File Validation & mime-type magic-number checks
  - [x] Content parsing (plaintext extraction, PDF structures)
  - [ ] OCR extraction service (Tesseract wrapper, local ONNX OCR model, or cloud API wrapper)
  - Audio transcription service (Whisper execution)

---

## 🔗 Phase 4: AI Context, Intent & Relationships
- [x] AI Analysis & Understanding Engine (Provider Abstraction, Prompt Manager, Structured JSON Output)
- [x] Database Schema: Relational tables
  - Add tables for `relationships`, `contexts`, `context_assets`, `context_events`, `intent_types`, `intent_predictions`, `context_versions`
- [x] Entity Extraction: Parser targeting dates, locations, merchants, transaction values, and names
- [x] Relationship Discovery Engine:
  - Direct relationship modeling: `RELATED_TO`, `PART_OF`, `MENTIONS`, `SAME_EVENT`, `SAME_LOCATION`
- [x] Context & Intent Engine:
  - Dynamic real-world context auto-discovery (*Japan Vacation*, *Buying a Car*, *Tax Filing*, *Medical Treatment*, *Stanford Admission*)
  - 12 Intent dimensions (*Travel*, *Finance*, *Health*, *Legal*, *Identity*, *Education*, *Business*, *Shopping*, *Personal*, *Entertainment*, *Research*, *Knowledge*)
  - Automatic event extraction (*Trips*, *Meetings*, *Purchases*, *Hospital Visits*, *Flights*, *Contract Signings*)
  - Multi-context asset membership scoring & version auditing
  - Web & Mobile Context Dashboard, Timeline, Intent Card, Suggested Contexts, and Detail View
- [ ] User Feedback loop:
  - API endpoint `POST /v1/feedback/verify` to confirm or correct AI-generated metadata/tags

---

## 🔍 Phase 5: Semantic Retrieval & Persona Intelligence
- [x] Add `pgvector` extension support to PostgreSQL migration (`015_create_embedding_tables.sql`)
- [x] Build vector embedding tables (`embeddings`, `embedding_chunks`, `embedding_versions`, `embedding_jobs`, `embedding_providers`, `vector_metadata`)
- [x] Embeddings generation pipeline & Knowledge Representation Layer:
  - Rich structured representation input (Title + Summary + Content + Entities + Relationships + Contexts + Intent + Memory Score + Tags)
  - Provider Abstraction: OpenAI, Gemini, Voyage, Cohere, Ollama, Local Deterministic
  - Vector Store Repository Abstraction: PGVector, Qdrant, Weaviate, Pinecone, Milvus, InMemory
  - Intelligent Chunking Engine: Semantic, Heading-aware, Page-aware, Hierarchical, Document-aware
  - Deduplication via SHA-256 content hashes
  - Model versioning & upgrade re-indexing (`RebuildForVersion`)
  - Web UI: EmbeddingDashboard, EmbeddingStatus, ModelInformation, GenerationHistory, `/embedding` page route
  - Mobile UI: EmbeddingStatus, ProcessingProgress
  - 6/6 Go tests passing
- [x] Hybrid Search Engine (`POST /api/v1/search/query`):
  - Database migration `016_create_search_tables.sql` (5 tables: search_history, saved_searches, search_statistics, search_preferences, search_index_versions)
  - Query Parser detecting intent, entities, locations, file types, quoted terms, excluded terms, and date ranges
  - Search Planner dynamically selecting strategy combinations (Keyword, Vector, Entity, Context, Relationship, Memory, Metadata)
  - Weighted Score Fusion combining keyword FTS, vector embeddings, entity match, context intent match, and recency
  - Complete match explainability (WhyMatched, ContributingStrategies, MatchedEntities, MatchedContexts, RelatedMemories)
  - 8 REST API endpoints (query, saved searches, history, suggestions, stats, preferences)
  - Web UI: GlobalSearch, SearchResults, AdvancedFilters, SavedSearches, `/search` page route
  - Mobile UI: GlobalSearch, SearchResults, SearchFilters, SavedSearches
- [x] Knowledge Insights & Timeline Engine:
  - Database migration `017_create_timeline_and_insights_tables.sql` (6 tables: timeline_events, timeline_groups, knowledge_insights, insight_history, insight_preferences, milestones)
  - TimelineEngine generating chronological event streams for Travel, Education, Medical, Financial, Projects, Legal, Purchases, Subscriptions, and Custom Contexts
  - InsightEngine & PatternDetector proactively identifying Expiration Warnings, Recurring Expenses, Top Places, Important/Missing Docs, and Knowledge Growth
  - MilestoneDetector tracking passport expirations, visa completions, tax filings, medical completions, and project milestones
  - Complete explainability (Why generated, evidence rationale, confidence)
  - 8 REST API endpoints (timeline, insights, milestones, dismiss, refresh, stats, preferences)
  - Web UI: TimelineView, InsightDashboard, MilestoneCards, ActivityFeed, `/insights` page route
- [x] Declutr AI Copilot (RAG & Personal Intelligence):
  - Database migration `018_create_copilot_tables.sql` (6 tables: conversations, messages, conversation_context, conversation_feedback, prompt_versions, response_history)
  - IntentParser classifying questions into SUMMARY, TIMELINE_QUERY, MEMORY_RECALL, ENTITY_EXPLORE, GENERAL_QA
  - GroundedRAGEngine & CopilotService implementing zero-hallucination grounded RAG answer synthesis strictly using retrieved vault documents
  - Full evidence citations (asset ID, title, type, snippet, confidence score, matched entities/contexts)
  - Multi-turn conversation management & history carry-over
  - 7 REST API endpoints (conversations, messages, feedback, SSE streaming)
  - Web UI: AIWorkspace, ConversationSidebar, ChatInterface, CitationViewer, SuggestedQuestions, `/copilot` page route
  - Mobile UI: ChatInterface, SourcePanel, ConversationHistory
- [x] Workflow Automation Engine:
  - Database migration `019_create_workflow_tables.sql` (7 tables: workflows, workflow_triggers, workflow_conditions, workflow_actions, workflow_runs, workflow_logs, workflow_history)
  - 12 Trigger Types (uploads, updates, deletions, context, memory, entity, document expiring, daily schedule, manual)
  - Rule Evaluator with AND, OR, NOT combinators across file types, entities, contexts, confidence, dates, storage size
  - Executable Actions (apply tags, create collection, move asset, generate summary, archive asset, create reminder, pin memory, notify user)
  - Async Execution Engine with step logging, retry handling, and metrics calculation
  - 8 REST API endpoints (workflows CRUD, toggle, manual run, history, stats)
  - Web UI: WorkflowDashboard, VisualRuleBuilder, ExecutionHistory, `/workflows` page route
  - Mobile UI: WorkflowList, WorkflowDetails, ExecutionHistory
- [x] Notification Center & Proactive Intelligence:
  - Database migration `020_create_notification_tables.sql` (6 tables: notifications, notification_rules, notification_preferences, notification_delivery, notification_history, digest_reports)
  - 10 Notification Categories (INFORMATION, SUCCESS, WARNING, CRITICAL, REMINDER, RECOMMENDATION, AI_INSIGHT, WORKFLOW, SECURITY, SYSTEM)
  - Dynamic PriorityEngine (LOW, MEDIUM, HIGH, URGENT)
  - Actionable alerts (Open Asset, View Context, Run Workflow, Retry Job, Dismiss, Pin, Archive, Snooze)
  - Proactive Digest Generator (Daily Intelligence Summaries & Weekly Recaps)
  - Deduplication engine & Preference channels manager (In-App, Email, Push, Desktop)
  - 7 REST API endpoints (notifications, mark read, dismiss, action, digests, preferences, stats)
  - Web UI: NotificationCenter, DigestView, NotificationPreferencesView, `/notifications` page route
  - Mobile UI: NotificationList, NotificationDetail, NotificationPreferences
- [x] Secure Sharing & Collaboration Platform:
  - Database migration `021_create_sharing_tables.sql` (7 tables: shares, share_permissions, share_members, share_links, share_comments, share_activity, share_invitations)
  - Granular Resource Types (ASSET, FOLDER, COLLECTION, CONTEXT, PROJECT, TIMELINE_VIEW, SEARCH_RESULT)
  - Role-based Access Control (READ_ONLY, COMMENT_ONLY, EDIT, OWNER, CO_OWNER)
  - Password-protected link sharing with expiration dates and download limits
  - Threaded discussion comments, replies, and resolution state
  - Auditable Activity Logging (viewed, downloaded, edited, commented, shared, permission changed, revoked)
  - 11 REST API endpoints (shares, invites, links, comments, activity, stats)
  - Web UI: ShareDialog, PermissionManager, CommentPanel, ActivityFeed, `/collaboration` page route
  - Mobile UI: ShareSheet, PermissionList, CommentThread
- [x] Version History, Recovery & Time Machine:
  - Database migration `022_create_versioning_tables.sql` (6 tables: resource_versions, version_snapshots, change_history, recycle_bin, restore_jobs, version_diffs)
  - Granular Resource Types (ASSET, METADATA, AI_ANALYSIS, CONTEXT, RELATIONSHIP, COLLECTION, MEMORY, WORKFLOW, PREFERENCES)
  - Snapshot Strategies (FULL, INCREMENTAL, DELTA, COMPRESSED, IMMUTABLE)
  - Field-level Compare Diff Engine (added, removed, modified key-values)
  - Recovery Engine with point-in-time restoration & Recycle Bin soft delete
  - 8 REST API endpoints (versions, snapshot, compare, restore, recyclebin, restore, purge, stats)
  - Web UI: VersionHistoryPanel, DiffViewer, RecycleBin, `/versioning` page route
  - Mobile UI: VersionList, VersionDetail, RecycleBin
- [x] Backup, Disaster Recovery & Business Continuity:
  - Database migration `023_create_backup_tables.sql` (7 tables: backups, backup_jobs, backup_files, backup_manifests, backup_history, restore_jobs, restore_history)
  - Backup Types (MANUAL, SCHEDULED, INCREMENTAL, FULL, ENCRYPTED, OFFLINE, COLD_STORAGE)
  - Full Vault Snapshot Engine (Assets, Metadata, AI Analysis, Entities, Memories, Workflows, Settings, Logs)
  - Automated Backup Scheduler (DAILY, WEEKLY, MONTHLY, CUSTOM_CRON) with retention policies
  - SHA-256 Checksum & Manifest Integrity Verification Engine
  - Disaster Recovery Restore Engine with scopes (FULL_VAULT, SELECTIVE) and strategies (OVERWRITE, NEW_VAULT, MERGE, DRY_RUN)
  - 8 REST API endpoints (backups, detail, schedule, restore, verify, cancel, stats)
  - Web UI: BackupDashboard, RestoreWizard, BackupHistory, `/backup` page route
  - Mobile UI: BackupStatus, RestoreHistory, ManualBackup
- [x] Security Center, Audit Hub & Trust Platform:
  - Database migration `024_create_security_tables.sql` (7 tables: security_events, security_scores, device_registry, trusted_devices, audit_events, risk_assessments, security_recommendations)
  - Security Posture Score (0-100 & Letter Grades A-F)
  - Risk Engine evaluating risk signals (new devices, failed logins, mass operations) to compute risk levels (LOW, MEDIUM, HIGH, CRITICAL)
  - Asynchronous Audit Engine logging events across 9 categories (AUTH, ASSET, SHARING, WORKFLOW, AI, SEARCH, BACKUP, VERSIONING, SETTINGS)
  - Session & Device Manager with session termination triggers & device trust toggles
  - Actionable Security Posture Recommendations
  - 8 REST API endpoints (dashboard, audit, sessions, terminate, devices, trust, risk, recommendations)
  - Web UI: SecurityDashboardComponent, AuditViewerComponent, SessionDeviceManagerComponent, `/security` page route
  - Mobile UI: SecurityOverview, SessionList, AuditSummary
- [x] Offline-First Sync Engine & Conflict Resolution:
  - Database migration `025_create_sync_tables.sql` (7 tables: sync_queue, sync_events, sync_conflicts, sync_sessions, device_state, sync_statistics, offline_operations)
  - Change Tracker & Queue Engine with statuses (QUEUED, UPLOADING, DOWNLOADING, RETRY, PAUSED, COMPLETED, FAILED, CANCELLED)
  - Conflict Resolver supporting Last Write Wins & 3-way Field-Level Merge (`MergeFieldLevel`)
  - Bidirectional Push/Pull delta streaming with per-device sequence checkpointing
  - Resume interrupted sync sessions on network reconnection
  - 7 REST API endpoints (push, pull, status, conflicts, resolve, register-device, stats)
  - Web UI: SyncCenterComponent, ConflictResolverComponent, SyncQueueViewerComponent, `/sync` page route
  - Mobile UI: OfflineBanner, SyncStatus, ConflictResolver
- [x] Integration Platform & Connector Framework:
  - Database migration `026_create_integration_tables.sql` (7 tables: connectors, connector_configs, connector_credentials, connector_sync_jobs, connector_webhooks, connector_logs, connector_health)
  - Connector SDK (`Initialize`, `Authenticate`, `Validate`, `Sync`, `Import`, `Export`, `Webhook`, `HealthCheck`, `Disconnect`)
  - Reference Connector Implementations (`GoogleDriveConnector` & `WebDAVConnector`)
  - Connector Marketplace (Google Drive, Dropbox, Notion, GitHub, AWS S3, WebDAV / Nextcloud)
  - Authentication Modes (OAuth 2.0, OAuth PKCE, API Keys, PAT, Service Accounts)
  - Inbound Webhook Ingestion & Event Bus Publishing
  - 8 REST API endpoints (integrations, install, configure, enable, sync, status, logs, webhooks)
  - Web UI: IntegrationMarketplaceComponent, InstalledConnectorsComponent, ConnectorLogsComponent, `/integrations` page route
  - Mobile UI: ConnectedServices, ConnectorCard, SyncTrigger
  - 6/6 Go tests passing
- [x] Reverse Persona Engine:
  - [x] Collect user interaction signals (ASSET_OPEN, SEARCH, PIN, UPLOAD, EDIT, CONTEXT_SWITCH, RELATIONSHIP_EXPLORE, COLLECTION_USE, TIME_OF_DAY, SEARCH_REFINEMENT, DASHBOARD_USAGE, FAVOURITE)
  - [x] Build personalization profile with time-based recency decay (exponential: `e^(−λ × days)`)
  - [x] Persona type inference: Traveller, Developer, Researcher, Healthcare Professional, Student, Entrepreneur, Designer, Photographer, Project Manager, Content Creator, Writer, Finance Professional
  - [x] Recommendation engine with full explainability (reason, confidence, evidence, contributing signals)
  - [x] Personal knowledge model: entities, locations, projects, interests, contacts, workflows
  - [x] Privacy controls: pause, disable signal types, reset, export, full GDPR deletion
  - [x] Web UI: PersonaDashboard, LearningInsights, RecommendationsPanel, InterestOverview, SignalSettings, PrivacyControls, /persona page route
  - [x] Mobile UI: PersonaDashboard, RecommendationsCard, PrivacyControls, LearningInsights
  - [x] Memory Engine & Knowledge Memory Foundation:
  - [x] Database migration `014_create_memory_tables.sql` (7 tables: memories, memory_sources, memory_scores, memory_events, memory_history, memory_settings, memory_clusters)
  - [x] 9 memory types (SHORT_TERM, WORKING, LONG_TERM, ARCHIVED, FORGOTTEN, PINNED, GENERATED, USER, AI)
  - [x] Dynamic composite scoring (`0.4×Importance + 0.3×Recency + 0.2×LogFreq + 0.1×Confidence`)
  - [x] Exponential recency decay (`e^(−λ × days)`) with configurable auto-archive and auto-forget thresholds
  - [x] Incremental consolidation into topic clusters and duplicate merging
  - [x] MemoryWorker pipeline registration (Context Engine → Persona Engine → Memory Formation → Knowledge Memory)
  - [x] 10 REST API endpoints (memories, timeline, detail, refresh, pin, archive, stats, reset, delete, settings)
  - [x] Web UI: MemoryDashboard, TimelineView, PinnedMemories, MemoryExplorer, `/memory` page route
  - [x] Mobile UI: MemoryFeed, MemoryTimeline, PinnedMemories, MemoryDetails
  - [x] 10/10 Go tests passing

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

## 🚀 Phase 7: Deployment & Optimization (Issue #031 Complete)
- [x] Performance Optimizations:
  - PostgreSQL indexes and pgvector HNSW index configurations (`027_production_hardening.sql`)
  - Redis cache policy setups & Cache Abstraction Layer (`backend/shared/cache/cache.go`)
- [x] Production Observability & Hardening Platform:
  - Structured Logging (`backend/shared/observability/observability.go`)
  - Metrics Collection (`declutr_http_requests_total`, `declutr_http_latency_average_ms`, `/metrics`)
  - Distributed Tracing spans & correlation headers
  - Diagnostic Probes (`/health`, `/ready`, `/live`, `/version`)
  - Token-Bucket Rate Limiter (`backend/shared/ratelimit/ratelimit.go`)
  - Background Worker Supervisor with auto-restart (`backend/shared/supervisor/supervisor.go`)
  - Circuit Breakers & Resilience (`backend/shared/resilience/resilience.go`)
  - Security Middleware & CSP/HSTS Headers (`backend/shared/middleware/security.go`)
- [x] Production-ready Dockerfiles, Docker Compose, Kubernetes manifests & Helm charts
- [x] Web Admin Console (`frontend/app/admin/page.tsx`) & Mobile Status Card
- [x] Automated CI/CD Workflows (`.github/workflows/ci.yml`)
- [x] Comprehensive Platform Go Test Suite (`backend/tests/platform_test.go`)
- [x] Production Documentation Suite (`docs/production/`)

---

## 🏢 Phase 8: Enterprise Organizations, Multi-Tenancy & Administration (Issue #032 Complete)
- [x] Multi-Tenant Architecture & Data Isolation (`backend/shared/middleware/tenant.go`)
- [x] Organization & Workspace Domain Engine (`backend/modules/organization/`)
- [x] Workspaces Classification (`PERSONAL`, `ORGANIZATION`, `DEPARTMENT`, `SHARED`, `ARCHIVED`)
- [x] Member Management & Statuses (`ACTIVE`, `INVITED`, `SUSPENDED`, `DEACTIVATED`) & Ownership Transfer
- [x] Role-Based Access Control (RBAC) & 10 Granular Permissions (`OWNER`, `ADMINISTRATOR`, `MANAGER`, `EDITOR`, `CONTRIBUTOR`, `VIEWER`, `GUEST`)
- [x] Teams & Department Groups with Permission Inheritance
- [x] Enterprise Policy Engine (`PASSWORD`, `SESSION_TIMEOUT`, `MFA`, `SHARING`, `RETENTION`, `AI_USAGE`, `WORKFLOW`)
- [x] SSO Framework Abstraction (SAML 2.0, OIDC, Azure AD, Google Workspace, Okta)
- [x] Web Enterprise Portal (`frontend/app/organization/page.tsx`, `frontend/features/organization/components/`)
- [x] Mobile UI Components (`frontend/declutr-mobile/features/organization/components/`)
- [x] PostgreSQL Migration `028_create_organization_tables.sql`
- [x] Enterprise Test Suite (`backend/tests/organization_test.go`)
- [x] Enterprise Documentation Suite (`docs/enterprise/`)

---

## 💻 Phase 9: Public API, Developer SDK & Developer Platform (Issue #033 Complete)
- [x] Versioned REST API (`/api/v1/`) & OpenAPI 3.0 Specification (`docs/developer/openapi.json`)
- [x] Scoped API Keys (`declutr_live_...`), SHA-256 Hashing, & Key Rotation (`backend/modules/developer/`)
- [x] OAuth 2.1 PKCE Authorization Code Grant & Token Exchange Engine
- [x] Webhook Delivery Engine with HMAC-SHA256 Payload Signing (`X-Declutr-Signature`)
- [x] Webhook Exponential Backoff Retries & Dead Letter Queue (`webhook_dlq`)
- [x] Official TypeScript SDK (`sdks/typescript/`)
- [x] Official Go Client SDK (`sdks/go/`)
- [x] Official Python SDK (`sdks/python/`)
- [x] Official Declutr CLI Binary (`cli/cmd/declutr/main.go`)
- [x] Web Developer Portal (`frontend/app/developer/page.tsx`, `frontend/features/developer/components/`)
- [x] PostgreSQL Migration `029_create_developer_platform_tables.sql`
- [x] Developer Go Test Suite (`backend/tests/developer_test.go`)
- [x] Developer Documentation Suite (`docs/developer/`)



