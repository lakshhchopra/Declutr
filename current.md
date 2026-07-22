# Project Status - Declutr

This document provides a summary of the current status of the Declutr codebase, its architecture, and its git history.

## 🛠️ Codebase Overview

Declutr is structured as a production-grade modular monorepo:

1. **Backend ([/backend](file:///f:/Github/Declutr/backend))**
   - Written in Go.
   - Refactored into a Domain-Oriented Modular Monolith:
     - `cmd/server/`: Main application entrypoint (`main.go`).
     - `modules/`: Feature modules (`auth`, `vault`, `file`, `search`, `persona`, `behavior`) each owning their `domain`, `application`, `repository`, `transport`, and `validators` layers.
     - `shared/`: Cross-cutting concerns (`crypto`, `database`, `middleware`, `config`, `logger`, `errors`, `constants`, `utils`, `types`).
     - `platform/`: Platform drivers (`postgres`, `redis`, `storage`).
     - `pkg/`: Public utility packages (`health`).
2. **Web Frontend ([/frontend](file:///f:/Github/Declutr/frontend))**
   - Next.js application using TypeScript.
   - Restructured into feature-first architecture:
     - `app/`: Next.js App router pages.
     - `features/`: Web feature modules (`auth`, `vault`, `search`).
     - `shared/`: Shared components (`ui`, `layout`, `feedback`, `forms`), `hooks`, `lib`, `providers`, `services`, `api`, `types`, `constants`.
     - `styles/`: Global CSS styling.
3. **Mobile Frontend ([/frontend/declutr-mobile](file:///f:/Github/Declutr/frontend/declutr-mobile))**
   - React Native application managed via **Expo** (with TypeScript).
   - Core directories:
     - `app/`: Expo Router pages (`(tabs)`, `login`, `register`, `vault`, `modal`).
     - `features/`: Mobile feature modules.
     - `shared/`: Native components, constants, hooks, providers, services, api, and utils.
     - `navigation/`: Router navigation helpers.
4. **Database ([/database](file:///f:/Github/Declutr/database))**
   - Database project containing `migrations/`, `seeds/`, and `scripts/`.
5. **Docs & Supporting Infrastructure**
   - Categorized into `docs/architecture/`, `docs/api/`, `docs/development/`, `docs/references/`, `docs/adr/`, and `docs/images/`.
   - Infrastructure configurations under `infrastructure/` (`docker`, `compose`, `github`, `monitoring`, `deployment`, `terraform`, `kubernetes`).
   - Helper scripts under `scripts/` (`setup`, `dev`, `build`, `release`, `database`, `maintenance`).
   - Testing suites under `tests/` (`unit`, `integration`, `e2e`, `fixtures`, `helpers`).
   - Security documentation under `security/` (`policies`, `audits`, `documentation`).

---

## 📜 Dev History (Commit Log Summary)

- **Backup, Disaster Recovery & Business Continuity (Issue #027)**:
  - Created PostgreSQL database migration `database/migrations/023_create_backup_tables.sql` (`backups`, `backup_jobs`, `backup_files`, `backup_manifests`, `backup_history`, `restore_jobs`, `restore_history`).
  - Implemented domain models for `Backup`, `BackupJob`, `BackupManifest`, `BackupSchedule`, `RestoreJob`, `BackupStats`, `CreateBackupRequest`, `ScheduleBackupRequest`, `RestoreBackupRequest`, `VerifyBackupRequest`.
  - Built `BackupService` & `DisasterRecoveryEngine` managing encrypted full & incremental snapshot backups, automated scheduler policies (`DAILY`, `WEEKLY`, `MONTHLY`), SHA-256 integrity validation, and catastrophe recovery with customizable restore modes (`FULL_VAULT`, `SELECTIVE`) and strategies (`OVERWRITE_EXISTING`, `RESTORE_AS_NEW_VAULT`, `MERGE_RESTORE`, `DRY_RUN`).
  - Added 8 REST API endpoints (`POST /backups`, `GET /backups`, `GET /backups/detail`, `POST /backups/schedule`, `POST /backups/restore`, `POST /backups/verify`, `POST /backups/cancel`, `GET /backups/stats`).
  - Created Web UI module (`frontend/features/backup/components/`) featuring `BackupDashboard`, `RestoreWizard`, `BackupHistory`, and Next.js page route (`/backup`).
  - Created Mobile UI components (`frontend/declutr-mobile/features/backup/components/`): `BackupStatus.tsx`, `RestoreHistory.tsx`, `ManualBackup.tsx`.
  - Added comprehensive Go test suite (`backup_test.go`) — 6/6 tests passing: Create Manual Backup, Backup Schedule Policy, Integrity Verification, Disaster Recovery Restore, Cancel Backup Job, Backup Stats.

- **Version History, Recovery & Time Machine (Issue #026)**:
  - Created PostgreSQL database migration `database/migrations/022_create_versioning_tables.sql` (`resource_versions`, `version_snapshots`, `change_history`, `recycle_bin`, `restore_jobs`, `version_diffs`).
  - Implemented domain models for `ResourceVersion`, `VersionSnapshot`, `RecycleItem`, `VersionDiff`, `VersioningStats`, `CreateSnapshotRequest`, `CompareVersionsRequest`, `RestoreVersionRequest`.
  - Built `ComputeDiff` engine calculating added, removed, and modified key-value pairs between version snapshots.
  - Built `VersioningService` & `TimeMachineRecoveryEngine` managing automated snapshot capture, version timeline listing, point-in-time state restoration, Recycle Bin soft deletes, bulk restoration, and permanent purges.
  - Added 8 REST API endpoints (`GET /versions`, `POST /versions/snapshot`, `POST /versions/compare`, `POST /versions/restore`, `GET /recyclebin`, `POST /recyclebin/restore`, `DELETE /recyclebin/purge`, `GET /versions/stats`).
  - Created Web UI module (`frontend/features/versioning/components/`) featuring `VersionHistoryPanel`, `DiffViewer`, `RecycleBin`, and Next.js page route (`/versioning`).
  - Created Mobile UI components (`frontend/declutr-mobile/features/versioning/components/`): `VersionList.tsx`, `VersionDetail.tsx`, `RecycleBin.tsx`.
  - Added comprehensive Go test suite (`versioning_test.go`) — 6/6 tests passing: Version Snapshot Creation, Diff Engine Comparison, Restore Version, Recycle Bin Soft Delete, Recycle Bin Restore & Purge, Versioning Stats.

- **Secure Sharing & Collaboration Platform (Issue #025)**:
  - Created PostgreSQL database migration `database/migrations/021_create_sharing_tables.sql` (`shares`, `share_permissions`, `share_members`, `share_links`, `share_comments`, `share_activity`, `share_invitations`).
  - Implemented domain models for `Share`, `SharePermission`, `ShareMember`, `ShareLink`, `ShareComment`, `ShareActivity`, `ShareInvitation`, `ShareStats`, `CreateShareRequest`, `InviteRequest`, `CreateLinkRequest`, `AddCommentRequest`.
  - Built `PermissionEngine` & `PermissionValidationEngine` validating role-based permissions (`READ_ONLY`, `COMMENT_ONLY`, `EDIT`, `OWNER`, `CO_OWNER`) across Assets, Folders, Collections, Contexts, Projects, and Timeline Views.
  - Built `CollaborationService` implementing share creation, member invitations, password-protected link generation, threaded comments with resolution, and audit trail logging.
  - Added 11 REST API endpoints (`POST /shares`, `GET /shares`, `DELETE /shares`, `POST /shares/invite`, `POST /shares/invite/accept`, `POST /shares/links`, `POST /shares/links/revoke`, `POST /shares/comments`, `GET /shares/comments`, `GET /shares/activity`, `GET /shares/stats`).
  - Created Web UI module (`frontend/features/collaboration/components/`) featuring `ShareDialog`, `PermissionManager`, `CommentPanel`, `ActivityFeed`, and Next.js page route (`/collaboration`).
  - Created Mobile UI components (`frontend/declutr-mobile/features/collaboration/components/`): `ShareSheet.tsx`, `PermissionList.tsx`, `CommentThread.tsx`.
  - Added comprehensive Go test suite (`collaboration_test.go`) — 6/6 tests passing: Share Creation & Permissions, Invitation Lifecycle, Link Sharing, Threaded Comments, Audit Activity Logging, Revoke Share.

- **Notification Center & Proactive Intelligence (Issue #024)**:
  - Created PostgreSQL database migration `database/migrations/020_create_notification_tables.sql` (`notifications`, `notification_rules`, `notification_preferences`, `notification_delivery`, `notification_history`, `digest_reports`).
  - Implemented domain models for `Notification`, `NotificationPreferences`, `DigestReport`, `NotificationStats`, `MarkReadRequest`, `ActionRequest`.
  - Built `PriorityEngine` dynamically calculating priority levels (`LOW`, `MEDIUM`, `HIGH`, `URGENT`) across 10 notification categories.
  - Built `NotificationService` implementing event subscription, deduplication, priority scoring, read/dismiss status, actionable button execution, and digest generation.
  - Added 7 REST API endpoints (`GET /notifications`, `POST /notifications/read`, `POST /notifications/dismiss`, `POST /notifications/action`, `GET /notifications/digests`, `GET/PUT /notifications/preferences`, `GET /notifications/stats`).
  - Created Web UI module (`frontend/features/notification/components/`) featuring `NotificationCenter`, `DigestView`, `NotificationPreferencesView`, and Next.js page route (`/notifications`).
  - Created Mobile UI components (`frontend/declutr-mobile/features/notification/components/`): `NotificationList.tsx`, `NotificationDetail.tsx`, `NotificationPreferences.tsx`.
  - Added comprehensive Go test suite (`notification_test.go`) — 6/6 tests passing: Notification Generation, Priority Engine, Mark Read & Dismiss, Deduplication, Digest Generation, Preferences Update.

- **Workflow Automation & Intelligent Actions Engine (Issue #023)**:
  - Created PostgreSQL database migration `database/migrations/019_create_workflow_tables.sql` (`workflows`, `workflow_triggers`, `workflow_conditions`, `workflow_actions`, `workflow_runs`, `workflow_logs`, `workflow_history`).
  - Implemented domain models for `Workflow`, `WorkflowTrigger`, `WorkflowCondition`, `WorkflowAction`, `WorkflowRun`, `WorkflowLog`, `WorkflowStats`, `RunWorkflowRequest`, `ToggleWorkflowRequest`.
  - Built `ConditionEvaluator` evaluating rule conditions (`EQUALS`, `CONTAINS`, `GREATER_THAN`, `LESS_THAN`, `IN`) with AND, OR, NOT combinators.
  - Built `ActionExecutor` & `WorkflowExecutionEngine` executing sequential & parallel actions (`APPLY_TAGS`, `CREATE_COLLECTION`, `MOVE_ASSET`, `GENERATE_SUMMARY`, `ARCHIVE_ASSET`, `CREATE_REMINDER`, `PIN_MEMORY`, `REFRESH_SEARCH_INDEX`, `NOTIFY_USER`).
  - Added 8 REST API endpoints (`POST /workflows`, `GET /workflows`, `PUT /workflows`, `DELETE /workflows`, `POST /workflows/toggle`, `POST /workflows/run`, `GET /workflows/history`, `GET /workflows/stats`).
  - Created Web UI module (`frontend/features/workflow/components/`) featuring `WorkflowDashboard`, `VisualRuleBuilder`, `ExecutionHistory`, and Next.js page route (`/workflows`).
  - Created Mobile UI components (`frontend/declutr-mobile/features/workflow/components/`): `WorkflowList.tsx`, `WorkflowDetails.tsx`, `ExecutionHistory.tsx`.
  - Added comprehensive Go test suite (`workflow_test.go`) — 6/6 tests passing: Workflow Creation & Toggle, Condition Evaluation, Workflow Execution Success, Condition Failure Handling, Execution History & Logs, Workflow Stats.

- **Declutr AI Copilot RAG & Personal Intelligence (Issue #022)**:
  - Created PostgreSQL database migration `database/migrations/018_create_copilot_tables.sql` (`conversations`, `messages`, `conversation_context`, `conversation_feedback`, `prompt_versions`, `response_history`).
  - Implemented domain models for `Conversation`, `Message`, `Citation`, `RAGContext`, `PromptVersion`, `SendMessageRequest`, `SendMessageResponse`, `FeedbackRequest`, `StreamChunk`.
  - Built `IntentParser` classifying question intent into `SUMMARY`, `TIMELINE_QUERY`, `MEMORY_RECALL`, `ENTITY_EXPLORE`, and `GENERAL_QA`.
  - Built `ContextBuilder` & `PromptBuilder` constructing structured versioned prompts with grounding rules, conversation history, and citations.
  - Built `GroundedRAGEngine` & `CopilotService` implementing zero-hallucination grounded RAG answer synthesis strictly using retrieved vault documents, confidence scoring, and reasoning overviews.
  - Added 7 REST API endpoints (`POST /copilot/conversations`, `GET /copilot/conversations`, `DELETE /copilot/conversations`, `POST /copilot/messages`, `GET /copilot/messages`, `POST /copilot/feedback`, `GET /copilot/messages/stream` SSE).
  - Created Web UI module (`frontend/features/copilot/components/`) featuring `AIWorkspace`, `ConversationSidebar`, `ChatInterface`, `CitationViewer`, `SuggestedQuestions`, and Next.js page route (`/copilot`).
  - Created Mobile UI components (`frontend/declutr-mobile/features/copilot/components/`): `ChatInterface.tsx`, `SourcePanel.tsx`, `ConversationHistory.tsx`.
  - Added comprehensive Go test suite (`copilot_test.go`) — 6/6 tests passing: Intent Parsing, Grounded RAG Answering, Multi-Turn Context Carry-Over, Hallucination Prevention, Conversation History, Feedback Ratings.

- **Knowledge Insights & Timeline Engine (Issue #021)**:
  - Created PostgreSQL database migration `database/migrations/017_create_timeline_and_insights_tables.sql` (`timeline_events`, `timeline_groups`, `knowledge_insights`, `insight_history`, `insight_preferences`, `milestones`).
  - Implemented domain models for `TimelineEvent`, `TimelineGroup`, `KnowledgeInsight`, `Milestone`, `InsightStats`, `InsightPreferences`, `TimelineFilter`.
  - Built `TimelineEngine` automatically generating chronological event streams for Travel, Education, Medical, Financial, Projects, Legal, Purchases, Subscriptions, and Custom Contexts.
  - Built `InsightEngine` & `PatternDetector` proactively scanning vault knowledge to identify Upcoming Expirations (Passport, Visa, Insurance), Recurring Expenses, Top Visited Places, Important/Missing Docs, and Knowledge Growth.
  - Built `MilestoneDetector` tracking passport expirations, visa completions, tax filings, medical completions, and project milestones.
  - Added 8 REST API endpoints (`GET /insights/timeline`, `GET /insights`, `GET /insights/milestones`, `POST /insights/dismiss`, `POST /insights/refresh`, `GET /insights/stats`, `GET/PUT /insights/preferences`).
  - Created Web UI module (`frontend/features/insights/components/`) featuring `TimelineView`, `InsightDashboard`, `MilestoneCards`, `ActivityFeed`, and Next.js page route (`/insights`).
  - Created Mobile UI components (`frontend/declutr-mobile/features/insights/components/`): `TimelineView.tsx`, `InsightsDashboard.tsx`, `MilestoneCards.tsx`, `ActivityFeed.tsx`.
  - Added comprehensive Go test suite (`insights_test.go`) — 6/6 tests passing: Timeline Generation, Insight Detection, Milestone Tracking, Pattern Recognition, Incremental Refresh, Insight Stats.

- **Hybrid Knowledge Search Engine (Issue #020)**:
  - Created PostgreSQL database migration `database/migrations/016_create_search_tables.sql` (`search_history`, `saved_searches`, `search_statistics`, `search_preferences`, `search_index_versions`).
  - Implemented domain models for `ParsedQuery`, `SearchPlan`, `SearchQueryRequest`, `SearchResultItem`, `SearchQueryResponse`, `SavedSearch`, `SearchHistoryItem`, `SearchStats`, `SearchPreferences`, `RankingWeights`, `SearchFilters`.
  - Built `QueryParser` detecting intent, entities, locations, file types, quoted exact terms, excluded terms (`-term`), and year/date ranges.
  - Built `SearchPlanner` dynamically selecting strategy combinations (Keyword, Vector, Entity, Context, Relationship, Memory, Metadata).
  - Built `SearchService` & `HybridSearchEngine` executing parallel retrievers, weighted score fusion, deduplication, and complete match explainability (`WhyMatched`, `ContributingStrategies`, `MatchedEntities`, `MatchedContexts`, `RelatedMemories`).
  - Added 8 REST API endpoints (`POST /search/query`, `POST /search/saved`, `GET /search/saved`, `DELETE /search/saved`, `GET /search/history`, `GET /search/suggestions`, `GET /search/stats`, `GET/PUT /search/preferences`).
  - Created Web UI module (`frontend/features/search/components/`) featuring `GlobalSearch`, `SearchResults`, `AdvancedFilters`, `SavedSearches`, and Next.js page route (`/search`).
  - Created Mobile UI components (`frontend/declutr-mobile/features/search/components/`): `GlobalSearch.tsx`, `SearchResults.tsx`, `SearchFilters.tsx`, `SavedSearches.tsx`.
  - Added comprehensive Go test suite (`search_test.go`) — 7/7 tests passing: Query Parsing, Keyword Search, Hybrid Search & Fusion, Search Filtering, Match Explainability, Saved Searches, Search History & Stats.

- **Embedding Engine & Knowledge Representation Layer (Issue #019)**:
  - Created PostgreSQL database migration `database/migrations/015_create_embedding_tables.sql` (`embeddings`, `embedding_chunks`, `embedding_versions`, `embedding_jobs`, `embedding_providers`, `vector_metadata`).
  - Implemented domain models for `Embedding`, `EmbeddingChunk`, `EmbeddingVersion`, `EmbeddingJob`, `EmbeddingProviderConfig`, `VectorMetadata`, `StructuredRepresentationInput`, `ChunkResult`, `GenerationOptions`, `EmbeddingStats`.
  - Built Provider Abstraction (`providers/provider.go`) supporting OpenAI, Gemini, Voyage, Cohere, Ollama, and Local deterministic provider.
  - Built Vendor-Independent Vector Database Repository (`repository/repository.go`) supporting PGVector, Qdrant, Weaviate, Pinecone, Milvus, and thread-safe InMemory driver with Cosine Similarity search.
  - Built Intelligent Chunking Engine (`chunking/chunker.go`) supporting Semantic, Heading-aware, Page-aware, Hierarchical, and Document-aware strategies.
  - Built `EmbeddingService` & `EmbeddingEngine` performing rich structured knowledge vectorization, SHA-256 deduplication, incremental refresh, and model version re-indexing (`RebuildForVersion`).
  - Registered `EmbeddingWorker` into the processing pipeline: `Memory Engine` ↓ `Embedding Engine` ↓ `Vector Storage`.
  - Added 7 REST API endpoints (`POST /embedding/generate`, `POST /embedding/refresh`, `GET /embedding/status`, `GET /embedding/stats`, `GET /embedding/history`, `PUT /embedding/provider`, `POST /embedding/rebuild`).
  - Created Web UI module (`frontend/features/embedding/components/`) featuring `EmbeddingDashboard`, `EmbeddingStatus`, `ModelInformation`, `GenerationHistory`, and Next.js page route (`/embedding`).
  - Created Mobile UI components (`frontend/declutr-mobile/features/embedding/components/`): `EmbeddingStatus.tsx`, `ProcessingProgress.tsx`.
  - Added comprehensive Go test suite (`embedding_test.go`) — 6/6 tests passing: Structured Embedding Generation, Intelligent Chunking (5 strategies), Provider Switching (6 providers), Vector Store Nearest Search, Incremental Updates, Version Upgrades.

- **Memory Engine & Knowledge Memory Foundation (Issue #018)**:
  - Created PostgreSQL database migration `database/migrations/014_create_memory_tables.sql` (`memories`, `memory_sources`, `memory_scores`, `memory_events`, `memory_history`, `memory_settings`, `memory_clusters`).
  - Implemented domain models for `Memory`, `MemorySource`, `MemoryScore`, `MemoryEvent`, `MemoryHistory`, `MemorySettings`, `MemoryCluster`, `MemoryDetail`, `MemoryStats`, `MemoryFormationRequest`.
  - Built `MemoryService` supporting: automatic memory formation from Context, Persona, Entities, and Assets; composite scoring (`0.4×Importance + 0.3×Recency + 0.2×LogFreq + 0.1×Confidence`); configurable exponential decay (`e^(−λ × days)`); soft auto-archiving/forgetting thresholds; incremental consolidation into topic clusters; duplicate merging; full lifecycle tracking.
  - Built `MemoryEngine` orchestrating incremental vault processing (Decay → Consolidation → Promotion/Demotion) without full rebuilds.
  - Registered `MemoryWorker` into processing pipeline: `Context Engine` ↓ `Persona Engine` ↓ `Memory Formation` ↓ `Knowledge Memory`.
  - Added 10 REST API endpoints (`GET/DELETE /memory`, `/memory/timeline`, `/memory/detail`, `/memory/refresh`, `/memory/pin`, `/memory/archive`, `/memory/stats`, `/memory/reset`, `GET/PUT /memory/settings`).
  - Created Web UI module (`frontend/features/memory/components/`) featuring `MemoryDashboard`, `TimelineView`, `PinnedMemories`, `MemoryExplorer`, and Next.js page route (`/memory`).
  - Created Mobile UI components (`frontend/declutr-mobile/features/memory/components/`): `MemoryFeed.tsx`, `MemoryTimeline.tsx`, `PinnedMemories.tsx`, `MemoryDetails.tsx`.
  - Added comprehensive Go test suite (`memory_test.go`) — 10/10 tests passing: Creation, Consolidation, Decay, Timeline, Pinning, Retrieval, Reset, Stats, Context Memory Formation, Persona Memory Formation.

- **Reverse Persona Engine (Issue #017)**:
  - Created PostgreSQL database migration `database/migrations/013_create_persona_tables.sql` (`persona_profiles`, `persona_signals`, `persona_scores`, `persona_interests`, `persona_recommendations`, `persona_settings`, `persona_history`).
  - Implemented domain models for `PersonaProfile`, `PersonaSignal`, `PersonaScore`, `PersonaInterest`, `PersonaRecommendation`, `PersonaSettings`, `PersonaHistory`, `PersonaExport`, `PersonaKnowledgeModel`.
  - Built `PersonaService` implementing: signal collection (privacy-aware), exponential recency decay scoring (`Recency = e^(−λ × days)`), persona type inference (12 types), recommendation generation with full explainability (reason, confidence, evidence, contributing signals), interest tracking, knowledge model builder.
  - Built `PersonaEngine` orchestrating incremental vault processing (Score → Profile → Recommendations) without full rebuilds.
  - Registered `PersonaWorker` into the processing pipeline: `Context Detection` ↓ `Behaviour Signals` ↓ `Persona Learning` ↓ `Recommendation Generation`.
  - Added 9 REST API endpoints (`GET /persona`, `/persona/recommendations`, `/persona/settings`, `PUT /persona/settings`, `POST /persona/signal`, `POST /persona/reset`, `GET /persona/export`, `DELETE /persona`, `GET /persona/history`).
  - Created Web UI module (`frontend/features/persona/components/`) featuring `PersonaDashboard`, `LearningInsights`, `RecommendationsPanel`, `InterestOverview`, `SignalSettings`, `PrivacyControls`, and Next.js page route (`/persona`).
  - Created Mobile UI components: `PersonaDashboard.tsx`, `RecommendationsCard.tsx`, `PrivacyControls.tsx`, `LearningInsights.tsx`.
  - Full privacy model: pause learning, disable individual signal types, reset persona, export (JSON), full GDPR deletion — all via UI and API.
  - Added comprehensive Go test suite (`persona_test.go`) — 9/9 tests passing: Privacy-Aware Signals, Developer Persona, Traveller Persona, Recommendation Explainability, Interest Tracking, Reset, GDPR Delete, Export Bundle, Signal Type Disabling.

- **Context & Intent Engine (Issue #016)**:
  - Created PostgreSQL database migration `database/migrations/012_create_context_and_intent_tables.sql` (`contexts`, `context_assets`, `context_events`, `intent_types`, `intent_predictions`, `context_versions`).
  - Implemented Domain models for `Context`, `ContextAsset`, `ContextEvent`, `IntentType`, `IntentPrediction`, `ContextVersion`, and `ContextStats`.
  - Built `ContextService` & `ContextEngine` performing LLM-backed dynamic context resolution, deduplication, multi-context asset membership scoring, and event extraction (*Trips*, *Meetings*, *Purchases*, *Hospital Visits*, *Flights*, *Interviews*).
  - Registered `ContextWorker` into the processing pipeline (`Relationship Discovery` ↓ `Context Detection` ↓ `Intent Prediction` ↓ `Context Graph`).
  - Added REST API endpoints (`/api/v1/context`, `/api/v1/context/details`, `/api/v1/context/refresh`, `/api/v1/context/intent`, `/api/v1/context/stats`).
  - Created Web UI module (`frontend/features/context/components/`) featuring `ContextDashboard`, `ContextTimeline`, `IntentCard`, `SuggestedContexts`, `ContextDetailView`, and Next.js page route (`/context`).
  - Created Mobile UI components (`ContextsScreen.tsx`, `ContextTimeline.tsx`, `IntentSummaryCard.tsx`).
  - Added comprehensive Go unit and integration test suite (`context_test.go`) validating Travel, Medical, Invoices, Receipts, Education, Legal, Projects, and Notes domain scenarios.
- **Relationship Discovery Engine (Issue #015)**:
  - Created PostgreSQL database migration `database/migrations/011_create_knowledge_graph_tables.sql` (`graph_nodes`, `graph_edges`, `graph_edge_evidence`, `graph_versions`).
  - Implemented Domain models for `GraphNode`, `GraphEdge`, `EdgeEvidence`, `RelationshipType` (`MENTIONS`, `RELATED_TO`, `LOCATED_AT`, etc.).
  - Built `MockRelationshipDiscoverer` generating high-confidence relational edges based on entities found in the document.
  - Built `GraphDiscoveryWorker` to integrate with the Processing Engine orchestration queue.
  - Added REST API endpoints (`/api/v1/graph/relationships/:nodeId`).
  - Created Web UI Component (`frontend/features/graph/components/relationship-panel.tsx`) featuring Relationship lists and detailed Edge Evidence viewers.
  - Created Mobile UI Component (`RelationshipViewer.tsx`) for native display of graph edges.
- **Entity Extraction & Knowledge Foundation (Issue #014)**:
  - Created PostgreSQL database migration `database/migrations/010_create_entity_extraction_tables.sql` (`entity_types`, `entities`, `entity_aliases`, `entity_occurrences`).
  - Implemented Domain models for `Entity`, `EntityType`, `EntityOccurrence`.
  - Built `MockEntityExtractor` simulating NLP extraction of Organizations, Locations, Amounts, and Dates.
  - Built Canonical Entity Resolution logic tying `OriginalValue` aliases to a single resolved `CanonicalName` securely within a `VaultID`.
  - Built `EntityExtractionWorker` to integrate with the Processing Engine orchestration queue.
  - Added REST API endpoints (`/api/v1/entities`, `/api/v1/entities/asset/:assetId`).
  - Created Web UI Component (`frontend/features/entities/components/entity-panel.tsx`) featuring hover-inspectable Entity Chips grouped by type.
  - Created Mobile UI Component (`EntityViewer.tsx`) for native display of entity cards.
- **AI Analysis & Understanding Engine (Issue #013)**:
  - Created PostgreSQL database migration `database/migrations/009_create_ai_analysis_tables.sql` (`ai_analysis`, `ai_classification`, `ai_tags`, `ai_topics`, `analysis_versions`).
  - Implemented Domain models for `AIAnalysis`, `Classification`, `Tag`, `Topic`, and `AnalysisVersion`.
  - Built LLM Provider Abstraction (`providers.LLMProvider`, `MockProvider`, and `OpenAIProvider` skeleton).
  - Built `PromptManager` to standardize extraction-to-prompt pipelines.
  - Built `AIAnalysisWorker` to integrate with the Processing Engine orchestration queue.
  - Added REST API endpoints (`/api/v1/analysis/:assetId`, `/history`, `/refresh`).
  - Created Web UI Component (`frontend/features/ai/components/analysis-panel.tsx`) with classification, tagging, and confidence indicators.
  - Created Mobile UI Component (`AnalysisViewer.tsx`) for native React Native display.
- **Universal Content Extraction Engine (Issue #012)**:
  - Created PostgreSQL database migration `database/migrations/008_create_content_extraction_tables.sql` (`extracted_documents`, `document_sections`, `document_blocks`, `document_versions`).
  - Implemented Domain models for the Normalized Document Model (`Document`, `Block`, `Section`, `Version`).
  - Built extensible `ExtractorRegistry` with `TextExtractor` (Markdown, TXT) and `StubDocumentExtractor` for PDFs/DOCX.
  - Built `ContentExtractionWorker` to integrate with the Processing Engine orchestration queue.
  - Added REST API endpoints (`/api/v1/content/:assetId`, `/history`, `/refresh`).
  - Created Web UI Component (`frontend/features/extraction/components/document-viewer.tsx`) to natively render extracted text blocks and headings.
  - Created Mobile UI Component (`ContentPreview.tsx`) for native React Native display of content blocks.
- **Metadata Extraction Engine (Issue #011)**:
  - Created PostgreSQL database migration `database/migrations/007_create_metadata_tables.sql` (`asset_metadata`, `asset_properties`, `asset_exif`, `metadata_versions`).
  - Implemented Domain models for `AssetMetadata`, `AssetProperties`, `AssetExif`, and `CompleteMetadata`.
  - Built extensible `ExtractorRegistry` with `ImageExtractor`, `TextExtractor`, and `MockComplexExtractor` for PDFs/Video/Audio.
  - Built `MetadataExtractionWorker` to integrate with the Processing Engine orchestration queue.
  - Added REST API endpoints (`/api/v1/metadata/:assetId`, `/history`, `/refresh`).
  - Created Web UI Metadata Panel (`frontend/features/metadata/components/metadata-panel.tsx`) with categorized factual metadata display.
  - Created Mobile UI component `MetadataViewer.tsx` for tracking metadata on native devices.
- **Content Processing Engine & Background Jobs (Issue #010)**:
  - Created PostgreSQL database migration `database/migrations/006_create_processing_tables.sql` (`processing_jobs`, `processing_workers`, `processing_events`, `processing_attempts`).
  - Implemented Domain models for `Job`, `Worker`, and lifecycle `Event`s with state machine statuses (Queued, Running, Failed, Retrying, etc.).
  - Built orchestration Engine (`backend/modules/processing/application/engine.go`), `JobScheduler`, and `RetryManager` with exponential backoff calculation.
  - Built `WorkerManager` to track worker node health, capabilities, and dynamic allocation.
  - Added REST API endpoints (`/api/v1/processing/jobs`, `/api/v1/processing/stats`, `/api/v1/processing/retry`) mapped in `transport/api.go`.
  - Created Web UI Processing Dashboard (`frontend/features/processing/components/processing-dashboard.tsx`) with real-time stats and `JobQueue` view.
  - Created Mobile UI component `ProcessingStatusCard.tsx` for tracking background processing on native devices.
- **Content Ingestion & Upload Pipeline (Issue #009)**:
  - Created PostgreSQL database migration `database/migrations/005_create_assets_and_ingestion_tables.sql` (`assets` and `upload_jobs` tables with status indexes).
  - Built storage provider abstraction layer `StorageProvider` (`backend/shared/storage/storage.go`) supporting S3, Cloudflare R2, and local file storage providers.
  - Implemented `Asset` domain model (`backend/modules/file/domain/asset.go`) featuring extensible pipeline status states (`QUEUED` ➔ `UPLOADING` ➔ `UPLOADED` ➔ `VALIDATING` ➔ `METADATA_PENDING` ➔ `AI_PENDING` ➔ `INDEXED_PENDING` ➔ `READY` / `FAILED`).
  - Added unit test suite `ingestion_test.go` covering status state transitions and job progress percentage logic.
  - Built client upload service `UploadService` (`frontend/features/upload/services/upload-service.ts`) computing WebCrypto SHA-256 checksums and file size validations.
  - Implemented interactive `UploadModal` (`frontend/features/upload/components/upload-modal.tsx`) supporting drag & drop, file browser, ingestion queue list, status badges, progress indicators, and cancellation.
- **Vault Workspace Foundation (Issue #008)**:
  - Created PostgreSQL database migration `database/migrations/004_create_vaults_table.sql` (`vaults` and `vault_settings` tables with owner_id foreign keys).
  - Built backend `Vault` domain model (`backend/modules/vault/domain/vault.go`) and unit tests (`application/vault_test.go`).
  - Implemented `VaultService` client (`frontend/features/vault/services/vault-service.ts`) managing default vault creation ("My Life Vault") and workspace metadata updates.
  - Enhanced Vault Overview view (`/vault`) with storage usage card, digital asset counters, collection counters, and premium zero-knowledge empty state.
- **User Onboarding, Profile & Preferences (Issue #007)**:
  - Created PostgreSQL database migration `database/migrations/003_create_user_profiles_and_preferences.sql` (`user_profiles` and `user_preferences` tables).
  - Implemented client-side Zod validation schemas (`profileSchema`, `preferencesSchema`) in `frontend/features/user/schemas/profile-schema.ts`.
  - Built `ProfileService` (`frontend/features/user/services/profile-service.ts`) managing `getProfile`, `updateProfile`, `getPreferences`, `updatePreferences`, and `completeOnboarding`.
  - Built interactive 8-step Onboarding flow page (`/onboarding`): Welcome, Display Name, Avatar Accent, Theme Mode, AI Behavior, Privacy Architecture Mode, Notifications, and Setup Completion.
  - Enhanced Settings page (`/settings`) with tabbed sections for General & Profile, Appearance, AI Behavior, Privacy Mode, and Notifications.
- **Session Management & Authentication Persistence (Issue #006)**:
  - Created PostgreSQL database migration `database/migrations/002_create_sessions_table.sql` (`user_sessions` table with indexes on `user_id`, `refresh_token_hash`, and `expires_at`).
  - Implemented session domain entity `UserSession` (`backend/modules/auth/domain/session.go`) with `IsActive()` expiration and revocation checkers.
  - Added unit test suite `session_test.go` covering active, expired, and revoked session states.
  - Enhanced `SessionProvider` (`frontend/shared/providers/session-provider.tsx`) with session persistence (survives page refreshes), status tracking (`loading`, `authenticated`, `unauthenticated`, `refreshing`), and `logout` / `logoutAll` handlers.
- **Authentication Integration (Issue #005)**:
  - Created typed API service client `AuthService` (`frontend/features/auth/services/auth-service.ts`) encapsulating `register`, `loginStart`, and `loginFinish` endpoints.
  - Connected Next.js authentication forms (`/login`, `/register`) to backend SRP APIs using TanStack React Query `useMutation`.
  - Integrated client-side SRP payload exchange:
    - Registration: `email`, `srpSalt`, `srpVerifier`, `mvk` ciphertext payload.
    - Login Start: `email` request ➔ returns `challengeId`, `srpSalt`, `serverPublicKey B`.
    - Login Finish: `challengeId`, `clientPublicKey A`, `clientProof M1` ➔ returns `serverProof M2` & token.
  - Handled network errors, server unavailable fallbacks, and user feedback toasts via `ToastProvider`.
- **SRP Authentication Backend Foundation (Issue #004)**:
  - Implemented production-ready zero-knowledge SRP-6a authentication backend architecture in `backend/modules/auth/`.
    - Domain: User credentials model, SRP challenge entity, session types.
    - Application: `Service`, `LoginStart`, `LoginFinish`, `Engine` math, and `ChallengeStore` single-use challenge expiration validator.
    - Repository: `UserRepository` interface and PostgreSQL `PostgresUserRepository` queries.
    - Transport & Endpoints: `POST /api/v1/auth/register`, `POST /api/v1/auth/login/start`, `POST /api/v1/auth/login/finish`.
  - Added unit test suite in `modules/auth/application/engine_test.go` and `store_test.go`.
- **Authentication UI & Onboarding Experience (Issue #003)**:
  - Created complete authentication user interface for Web and Mobile following zero-trust design guidelines.
  - Implemented client-side Zod validation schemas (`loginSchema`, `registerSchema`, `forgotPasswordSchema`, `resetPasswordSchema`) integrated with React Hook Form.
  - Built reusable auth components: `PasswordStrengthMeter`, `AuthCardLayout`, `SocialAuthButtons`.
  - Created complete Next.js App Router auth pages:
    - `/welcome` (Hero onboarding, product description, primary/secondary CTAs, zero-knowledge privacy statement)
    - `/login` (Email, password, show/hide toggle, remember device, social & passkey placeholders)
    - `/register` (Name, email, password, confirm password, password strength meter, terms acceptance)
    - `/forgot-password` (Passphrase recovery email dispatch)
    - `/reset-password` (New passphrase setup with strength validation)
    - `/verify-email` (6-digit code entry with resend timer)
    - `/magic-link-waiting` (Animated polling state UI)
    - `/auth-error` (SRP challenge proof failure view with retry actions)
- **Application Shell & Navigation Foundation (Issue #002)**:
  - Built responsive multi-device application shell (`AppShell`) with Desktop Sidebar, Top Navigation, Tablet Collapsible Sidebar, and Mobile Bottom Navigation bar.
  - Configured global application providers in `frontend/shared/providers/`.
  - Built reusable `PageShell` page template component with breadcrumbs, title, subtitle, and actions header.
  - Implemented application route structure with clean placeholder pages (`/dashboard`, `/vault`, `/search`, `/collections`, `/ai`, `/persona`, `/security`, `/settings`).
  - Added loading framework skeleton (`loading.tsx`), custom 404 page (`not-found.tsx`), and error boundary fallback (`error.tsx`).
- **Shared Design System Foundation (Issue #001)**:
  - Established centralized theme system (`ThemeProvider`) supporting Dark Mode (default), Light Mode, System Theme detection, and persistent `localStorage` preference.
  - Implemented semantic CSS design tokens in `globals.css`.
  - Built reusable `shadcn/ui` & Radix UI component primitives.

