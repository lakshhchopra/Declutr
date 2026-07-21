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

