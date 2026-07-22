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

## 🚀 Phase 7: Deployment & Optimization
- [ ] Performance Optimizations:
  - PostgreSQL indexes and pgvector HNSW index configurations
  - Redis cache policy setups for metadata querying
- [ ] Production-ready Dockerfiles & Helm charts
- [ ] Deploy staging environment
- [ ] Expo/Mobile app build profiles (eas build setup for iOS and Android)
- [ ] Final end-to-end system verification
