<div align="center">

  # 🛡️ Declutr
  ### **AI-Powered Intelligent Digital Life Vault**

  [![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](file:///f:/Github/Declutr/backend)
  [![Next.js](https://img.shields.io/badge/Next.js-15.2-000000?style=for-the-badge&logo=nextdotjs&logoColor=white)](file:///f:/Github/Declutr/frontend)
  [![React Native](https://img.shields.io/badge/React_Native-Expo_54-61DAFB?style=for-the-badge&logo=react&logoColor=black)](file:///f:/Github/Declutr/frontend/declutr-mobile)
  [![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16_+_pgvector-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)](file:///f:/Github/Declutr/database)
  [![Architecture](https://img.shields.io/badge/Architecture-Modular_Monolith-8A2BE2?style=for-the-badge)](file:///f:/Github/Declutr/docs/architecture/declutr_architecture_document.html)
  [![Security](https://img.shields.io/badge/Security-Zero--Knowledge_SRP--6a-00E676?style=for-the-badge)](file:///f:/Github/Declutr/docs/references/security_and_threat_model.md)

  <br />

  <!-- Live Status & Quick Action Buttons -->
  <a href="file:///f:/Github/Declutr/docs/architecture/declutr_architecture_document.html">
    <img src="https://img.shields.io/badge/📖_Read_Architecture_Spec-000000?style=for-the-badge&logoColor=white" alt="Architecture Spec" />
  </a>
  <a href="file:///f:/Github/Declutr/docs/api/api_specification.md">
    <img src="https://img.shields.io/badge/⚡_API_Specification-2563EB?style=for-the-badge&logoColor=white" alt="API Spec" />
  </a>
  <a href="file:///f:/Github/Declutr/todo.md">
    <img src="https://img.shields.io/badge/📋_Project_Roadmap-059669?style=for-the-badge&logoColor=white" alt="Roadmap" />
  </a>

  <br /><br />

  <p align="center">
    <b>Declutr</b> transforms fragmented digital memory into an encrypted, intent-aware digital vault.<br />
    Store, organize, contextually connect, and retrieve your files using natural human memory associations.
  </p>

</div>

---

## 🌟 Tech Stack & Infrastructure

<div align="center">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=go,nextjs,react,ts,postgres,redis,docker,tailwind,wasm,githubactions,linux,vscode&perline=6" alt="Tech Stack Icons" />
  </a>
</div>

<br />

| Layer | Technology | Primary Role |
| :--- | :--- | :--- |
| **Backend API** | **Go 1.26** | Domain-Oriented Modular Monolith with clean application layers |
| **Web Frontend** | **Next.js 15 (TypeScript)** | App Router, Tailwind CSS, client-side encryption via WASM |
| **Mobile App** | **React Native (Expo 54)** | Cross-platform iOS/Android native mobile vault interface |
| **Database** | **PostgreSQL 16 + pgvector** | Unified relational metadata store + 512-dim vector embeddings |
| **Queue Workers** | **Redis + Asynq** | Asynchronous OCR parsing, transcription, and embedding tasks |
| **Cloud Storage** | **S3-Compatible (Cloudflare R2)** | Zero-egress direct-to-object chunked file storage |

---

## 🧠 Core Product Pillars

Declutr shifts digital storage from plain folder trees to an **Intelligent Personal Digital Memory System**:

```
 🧠 Content Intelligence       🎯 Intent Intelligence       🔗 Relationship Intelligence
 Extract OCR text, document    Classify item utility        Connect boarding passes, hotel
 layouts, audio transcripts    (receipts, booking references, receipts, and recommendations
 & 512-dim semantic vectors.   expense claims, archives).   into a single "Trip" context.

 👤 Persona Intelligence       🔍 Retrieval Intelligence     🛡 Behavioral Security
 Reverse Persona modeling with Synthesize hybrid FTS         Passive session anomaly scoring
 recency decay preferences.    keyword + pgvector search.   with adaptive MFA prompts.
```

---

## 🏗️ System Architecture & Data Flow

```
  +---------------------------------------------------------------------------------+
  |                               USER / CLIENT APP                                 |
  |         (Next.js Web Client / React Native Expo Mobile / WASM Crypto)          |
  +---------------------------------------+-----------------------------------------+
                                          │
                                          │ HTTPS / SRP-6a Zero-Knowledge Protocol
                                          ▼
  +---------------------------------------------------------------------------------+
  |                        AUTHENTICATION & SESSION LAYER                           |
  |         (SRP-6a Verifier / Passkey Verification / JWT Refresh Rotation)          |
  +---------------------------------------+-----------------------------------------+
                                          │
                                          ▼
  +---------------------------------------------------------------------------------+
  |                                 DIGITAL VAULT                                   |
  |          (Logical Isolation, Client-Side Keys, Encrypted File Metadata)         |
  +---------------------------------------+-----------------------------------------+
                                          │
                                          ▼
  +---------------------------------------------------------------------------------+
  |                       CONTENT INGESTION & AI PIPELINE                           |
  | [File Validation] ➔ [Type Detection] ➔ [OCR/Whisper] ➔ [Embeddings (pgvector)]  |
  +---------------------------------------------------------------------------------+
```

---

## 🧭 Context & Intent Engine

Declutr's **Context & Intent Engine** is a core differentiator that understands **WHY** assets exist together without requiring users to manually organize or create folders.

### 1. Processing Pipeline
```
Upload ➔ Metadata ➔ Content Extraction ➔ AI Understanding ➔ Entities ➔ Relationships ➔ Context & Intent Engine ➔ Context Graph
```

### 2. Context Model
Dynamic, zero-manual creation contexts scoped to the user's Vault:
- **Travel / Vacations** (*Japan Vacation*, *European Tour*)
- **Financial & Property** (*Buying a Car*, *Home Purchase*, *Tax Filing 2025*)
- **Health & Medical** (*Medical Treatment*, *Cardiology Consultation*)
- **Education & Growth** (*University Admission*, *Stanford Application*)
- **Legal & Administrative** (*Visa Application*, *Lease Agreement*)

### 3. Intent Model
Distinguishes the real-world utility of assets across 12 canonical intent dimensions:
`Travel`, `Finance`, `Health`, `Legal`, `Identity`, `Education`, `Business`, `Shopping`, `Personal`, `Entertainment`, `Research`, `Knowledge`.

### 4. Automatic Event Detection
Identifies key events within context timelines:
`Trip`, `Meeting`, `Purchase`, `Hospital Visit`, `Flight`, `Conference`, `Contract Signing`, `Birthday`, `Anniversary`, `Interview`.

### 5. Structured AI Prediction & Audit
Every prediction output includes:
- **Confidence Score** (probabilistic thresholding)
- **Evidence** (extracted document text snippets & entity overlap)
- **Reasoning** (LLM inference rationale)
- **Prompt Versioning** (engine version audit logs)

---

## 📂 Monorepo Repository Structure

```
Declutr Monorepo/
├── 📁 backend/                  # Go Domain-Oriented Modular Monolith
│   ├── 📁 cmd/server/           # Backend entrypoint (main.go)
│   ├── 📁 modules/              # Feature modules (auth, vault, file, search, persona, behavior)
│   ├── 📁 shared/               # Cross-cutting concerns (crypto, database, middleware)
│   ├── 📁 platform/             # Drivers (postgres, redis, storage)
│   └── 📁 pkg/                  # Public shared utilities (health check)
├── 📁 frontend/                 # Next.js Web Client (TypeScript)
│   ├── 📁 app/                  # App router pages
│   ├── 📁 features/             # Web feature modules (auth, vault, search)
│   ├── 📁 shared/               # UI components, hooks, providers, API services
│   └── 📁 declutr-mobile/       # React Native / Expo Mobile Client
│       ├── 📁 app/              # Expo router screens
│       ├── 📁 features/         # Mobile feature modules
│       └── 📁 shared/           # Native components, hooks, services
├── 📁 database/                 # PostgreSQL migrations, seeds, and SQL scripts
├── 📁 docs/                     # Full technical docs (architecture, api, references, adr)
├── 📁 infrastructure/           # Docker, Compose, K8s, Terraform, Monitoring configs
├── 📁 scripts/                  # Development, build, and maintenance automation
├── 📁 security/                 # Security policies and threat model documentation
└── 📁 tests/                    # Unit, Integration, and E2E test suites
```

---

## 🔒 Security & Key Wrapping Architecture

Declutr operates on zero-trust principles. Server databases store no plaintext passwords or unencrypted master keys:

```
  [User Password] ──(Argon2id)──> [Master Encryption Key (MEK)]
                                            │
                                            ▼ (Unwraps)
                                  [Master Vault Key (MVK)]
                                            │
                                            ▼ (Unwraps)
                                     [Vault Key (VK)]
                                            │
                                            ▼ (Encrypts File Block)
                                     [File Key (FK)]
```

- **Zero-Knowledge Auth:** Secure Remote Password (SRP-6a) exchange prevents plaintext credentials or password hashes from hitting the network.
- **Row-Level Security:** PostgreSQL Row-Level Security (RLS) ensures cryptographic user context isolation.

---

## 🚀 Quick Start Guide

### Prerequisites
- **Node.js:** v18+ 
- **Go:** v1.22+
- **Docker & PostgreSQL:** (with `pgvector` enabled)

### 1. Run Backend (Go)
```bash
cd backend
go run ./cmd/server
# Backend starts on http://localhost:8080
```

### 2. Run Web Client (Next.js)
```bash
cd frontend
npm install
npm run dev
# Web app available at http://localhost:3000
```

### 3. Run Mobile Client (React Native / Expo)
```bash
cd frontend/declutr-mobile
npm install
npm run start
```

---

## 📊 Contribution & Activity

<div align="center">
  <img src="https://github-readme-stats.vercel.app/api?username=diablovocado&show_icons=true&theme=dark&hide_border=true" alt="Declutr Stats" height="150" />
  <img src="https://github-readme-stats.vercel.app/api/top-langs/?username=diablovocado&layout=compact&theme=dark&hide_border=true" alt="Top Languages" height="150" />
</div>

<br />

Contributions are strictly governed by our [CONTRIBUTING.md](file:///f:/Github/Declutr/CONTRIBUTING.md) guide. All pull requests must pass strict modular boundary checks and linting suites.

---

## 📄 License

Distributed under the MIT License. See [LICENSE](file:///f:/Github/Declutr/LICENSE) for more information.

---

## 📄 Content Processing & Extraction

Declutr's ingestion pipeline normalizes documents and media into a common structured format using the **Universal Content Extraction Engine**.

### Extractor Interface
Every supported file type is parsed by a specific extractor implementing the `ContentExtractor` interface:
- **TextExtractor**: Handles `text/plain`, `text/markdown`, `text/csv`, and `application/json`.
- **StubDocumentExtractor**: Handles `application/pdf`, `docx`, and `pptx` (stubbed for future native integration).

### Normalized Document Model
Regardless of the original file format, Declutr translates the content into a normalized PostgreSQL schema (`extracted_documents`, `document_sections`, `document_blocks`).
- **Blocks**: The atomic unit of content (Heading, Paragraph, List, Code, Table, Caption, Link).
- **Sections**: Hierarchical groupings mapping to pages or chapters.
- **Attributes**: Language, Word/Char counts, Estimated Reading Time.

This ensures all downstream AI embeddings and search indices consume the exact same structured interface without needing to understand the original file's binary layout.

---

## 🧠 AI Analysis & Understanding Engine

The AI Understanding Engine takes the normalized text from the Extraction Engine and generates structured semantic analysis. It is designed to be fully LLM-agnostic through a **Provider Abstraction** layer (`MockProvider`, `OpenAIProvider`, `AnthropicProvider`, etc.).

### Structured Output
For every document, the engine strictly generates:
- **Summaries**: Short summary, detailed summary, and document purpose.
- **Classification**: Document category (e.g., Receipt, Invoice), document type, and quality markers (scanned, incomplete, corrupted).
- **Metadata**: Sentiment, complexity, reading level, language, and writing style.
- **Tags & Topics**: Deduplicated arrays of semantic tags and topics.
- **Confidence Scores**: Every generated field is backed by an AI confidence score (0.0 - 1.0) for UI transparency.

### Prompt Strategy & Retry Loop
The `PromptManager` compiles normalized blocks into a strict JSON-schema prompt. The engine features an exponential backoff retry loop and full token usage tracking per-request to manage LLM costs.

---

## 🏛️ Entity Extraction & Knowledge Foundation

The Entity Extraction Engine converts the structured AI Analysis into atomic, semantic entities. These entities form the bedrock of the future Knowledge Graph, Relational Engine, and Context-Aware Search.

### Supported Entity Types
The system supports strongly-typed extraction across dozens of categories, including `Person`, `Organization`, `Location`, `Date`, `Amount`, `Product`, and various `Identifier`s (passports, VINs, accounts).

### Canonical Entity Resolution
Entities are deduplicated and resolved to a Canonical Name. For example, the terms *"NYC"*, *"New York"*, and *"New York City"* appearing across dozens of different files are all resolved to a single Canonical Entity (`New York City`).
- The `entity_aliases` table tracks the mapping of historical aliases to the canonical ID.
- The `entity_occurrences` table binds specific entity discoveries to the original `asset_id` and tracks the specific substring (`original_value`) found in the document, along with a `confidence_score`.

### Security Boundaries
Entities are strictly bound to the `vault_id`. Deduplication and canonical resolution happen *within* a user's vault, ensuring no cross-pollination of private knowledge graph data between users.

---

## 🕸️ Relationship Discovery & Knowledge Graph

The Relationship Discovery Engine completes Phase 4 by inferring strictly relational connections between nodes (Assets, Entities, Collections).

### Edge & Evidence Architecture
The `graph_edges` table models typed relationships between a `source_node_id` and a `target_node_id` (e.g., `BELONGS_TO`, `MENTIONS`, `RELATED_TO`). 
Every single edge must include a corresponding entry in the `graph_edge_evidence` table explaining *why* the relationship was inferred (e.g., "Document explicitly mentions Google LLC in the AI Summary"), alongside a computed `confidence_score` and `strength`.

### Graph Incremental Updates
The background `GraphDiscoveryWorker` builds edges incrementally as new entities are detected or metadata changes, avoiding expensive full-graph rebuilds. All operations are strictly bound by the `vault_id` boundary.

---

## 🧬 Reverse Persona Engine

The Reverse Persona Engine is one of Declutr's core differentiators. It understands **how you naturally organise and interact with your own digital life** — without ever leaving your vault.

> **This is NOT advertising. It is NOT tracking. It is personal intelligence.**

### What It Does

The engine observes your natural interaction patterns and builds a dynamic, private profile of who you are digitally:

| Persona Type | Inferred From |
|---|---|
| Traveller | Flight docs, hotel bookings, itineraries, visa paperwork |
| Developer | Code files, technical search terms, debug notes |
| Researcher | Academic papers, citation notes, research journals |
| Healthcare Professional | Medical records, prescriptions, clinical notes |
| Student | Study materials, exam notes, university correspondence |
| Entrepreneur | Business plans, pitch decks, revenue spreadsheets |
| Designer | Figma files, mood boards, typography references |
| Photographer | RAW files, Lightroom presets, shoot briefs |

### Behaviour Signals

The engine collects signals with user consent and control:

```
ASSET_OPEN  ·  SEARCH  ·  PIN  ·  UPLOAD  ·  EDIT
CONTEXT_SWITCH  ·  RELATIONSHIP_EXPLORE  ·  COLLECTION_USE
TIME_OF_DAY  ·  SEARCH_REFINEMENT  ·  DASHBOARD_USAGE  ·  FAVOURITE
```

### Scoring Engine with Recency Decay

Every dimension is scored using an exponential decay model:

```
Recency  = e^(−decayRate × daysSinceLastSeen) × totalWeight
Importance = log(1 + frequency) × totalWeight
Confidence = min((Recency + Importance) / 10, 1.0)
```

Scores **decay naturally** — if you stop using a topic, it fades without requiring explicit action.

### Recommendation Engine

Every recommendation includes full explainability:

```json
{
  "title": "Continue Thesis Chapter 4",
  "reason": "You have 14 research interactions this month. Research is RISING.",
  "confidence": 0.87,
  "evidence": ["14 research interactions", "Trend: RISING"],
  "contributingSignals": ["Research frequency: 14", "Importance: 8.40"]
}
```

Recommendation types: `CONTINUE_PROJECT` · `RESUME_READING` · `RELATED_DOCUMENT` · `SUGGESTED_CONTEXT` · `SUGGESTED_COLLECTION` · `SUGGESTED_ARCHIVE` · `SUGGESTED_RELATIONSHIP`

### Personal Knowledge Model

The engine infers:
- Frequent entities, favourite locations, recurring projects
- Long-term interests, recurring contacts
- Common workflows, frequently accessed vault areas

### Database Schema (Migration 013)

| Table | Purpose |
|---|---|
| `persona_profiles` | Dynamic persona type + confidence + knowledge model |
| `persona_signals` | Raw behaviour events with weight and source |
| `persona_scores` | Scored dimensions with decay, trend, importance, recency |
| `persona_interests` | Long-term inferred interests with personal relevance |
| `persona_recommendations` | Explainable recommendations |
| `persona_settings` | Per-vault privacy controls |
| `persona_history` | Immutable audit snapshots |

### REST API

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/api/v1/persona` | Current persona profile + scores + interests |
| `GET` | `/api/v1/persona/recommendations` | Personalised recommendations |
| `GET` | `/api/v1/persona/settings` | Privacy settings |
| `PUT` | `/api/v1/persona/settings` | Update privacy settings |
| `POST` | `/api/v1/persona/signal` | Record a behaviour signal |
| `POST` | `/api/v1/persona/reset` | Reset all learned data |
| `GET` | `/api/v1/persona/export` | Export persona as JSON |
| `DELETE` | `/api/v1/persona` | Full GDPR deletion |
| `GET` | `/api/v1/persona/history` | Audit history snapshots |

### Privacy Guarantees

- ✅ User can pause learning at any time
- ✅ User can disable individual signal types
- ✅ User can reset the entire persona
- ✅ User can export all data as JSON
- ✅ Full GDPR-style deletion supported
- ✅ All data is vault-scoped — never shared, never sold
- ✅ Every decision is fully explainable
- ✅ No third-party analytics

---

## 🧠 Memory Engine & Knowledge Memory Foundation

The Memory Engine is the heart of Declutr. It transforms isolated assets into persistent knowledge. Unlike a simple vector database, knowledge memory **evolves over time** — remembering what matters, strengthening recurring patterns, and fading stale data.

> **Pipeline position**: `Upload` → `Metadata` → `Content Extraction` → `AI Understanding` → `Entities` → `Relationships` → `Context` → `Reverse Persona` → **Memory Engine** → `Knowledge Memory`

### Memory Types

- `SHORT_TERM` — Recently formed memories with low frequency
- `WORKING` — Active memories being accessed or updated
- `LONG_TERM` — High-strength, consolidated persistent knowledge
- `PINNED` — User-flagged permanent memories (immune to decay)
- `ARCHIVED` — Faded memories below the archive threshold
- `FORGOTTEN` — Stale memories below the forget threshold
- `GENERATED` / `USER` / `AI` — Source attribution types

### Composite Scoring & Configurable Decay

Memory strength is computed dynamically:

```
MemoryStrength = 0.4 × Importance + 0.3 × Recency + 0.2 × log(1 + Frequency)/5 + 0.1 × Confidence
Recency = e^(−decayRate × daysSinceLastSeen)
```

- **Exponential decay**: Stale memories fade naturally unless accessed or pinned.
- **Auto-Archiving**: Memories with strength < `autoArchiveThreshold` (0.10) become `ARCHIVED`.
- **Auto-Forgetting**: Memories with strength < `autoForgetThreshold` (0.01) become `FORGOTTEN`.
- **Never permanently deleted automatically**: Soft transition only; users retain full control.

### Incremental Memory Consolidation

- Automatically groups memories into topic clusters.
- Merges duplicates when matching memories are re-encountered.
- Incremental updates only — never rebuilds the entire memory graph.

### Database Schema (Migration 014)

| Table | Purpose |
|---|---|
| `memories` | Core memory object (type, strength, decay, pin status) |
| `memory_sources` | Contributing assets, entities, contexts, persona signals |
| `memory_scores` | Point-in-time strength audit snapshots |
| `memory_events` | Immutable lifecycle events (formed, strengthened, decayed, pinned, archived) |
| `memory_history` | Audit snapshots for state changes |
| `memory_settings` | Per-vault memory configuration & thresholds |
| `memory_clusters` | Groupings of related memories by topic |

### REST API

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/api/v1/memory` | Retrieve strongest or filtered memories |
| `GET` | `/api/v1/memory/timeline` | Get memories in chronological order |
| `GET` | `/api/v1/memory/detail` | Get full memory detail (sources, scores, events) |
| `POST` | `/api/v1/memory/refresh` | Trigger decay & consolidation cycle |
| `POST` | `/api/v1/memory/pin` | Pin memory (immune to decay) |
| `POST` | `/api/v1/memory/archive` | Archive a memory |
| `GET` | `/api/v1/memory/stats` | Vault memory statistics & clusters |
| `POST` | `/api/v1/memory/reset` | Reset memory model for vault |
| `DELETE` | `/api/v1/memory` | Permanently delete a memory |
| `GET` / `PUT` | `/api/v1/memory/settings` | Get / update memory settings |

---

## 💎 Embedding Engine & Knowledge Representation Layer

Declutr's Embedding Engine transforms enriched, structured knowledge into high-quality semantic vector representations. Rather than vectorising raw unformatted text alone, it builds dense vectors from structured representation inputs containing Title, Summary, Content, Entities, Relationships, Contexts, Intent, Memory Scores, Tags, and Classifications.

> **Pipeline Position**: `Metadata` → `Content Extraction` → `AI Analysis` → `Entities` → `Relationships` → `Context` → `Persona` → `Memory` → **`Embedding Engine`** → `Vector Store`

### Rich Structured Knowledge Vectorization

Embeddings are constructed from enriched input representations:

```
Title: Japan Travel Itinerary 2025
Summary: Three-week vacation covering Tokyo, Kyoto, and Osaka.
Classification: Travel Document
Intent: Vacation Planning
Contexts: Japan Vacation, Family Trip
Entities: Tokyo, Kyoto, Narita Airport
Relationships: Tokyo MENTIONS Narita Airport
Memory Score: 0.88
Tags: travel, japan, flights
```

This maximizes semantic precision while preventing duplicate vectors via SHA-256 content deduplication hashes.

### Provider Abstraction Layer

The engine abstracts model vendors via a unified Go interface:

- **OpenAI**: `text-embedding-3-small` / `text-embedding-3-large` (1536d / 3072d)
- **Google Gemini**: `text-embedding-004` (768d)
- **Voyage AI**: `voyage-3-lite` / `voyage-3` (1024d)
- **Cohere**: `embed-english-v3.0` (1024d)
- **Ollama**: Local privacy-first open models (`nomic-embed-text`, 768d)
- **Local Deterministic**: Zero-dependency L2-normalized synthetic vector generator for tests & local dev

### Vendor-Independent Vector Database Repository

Business layers interact with an abstract `VectorStoreRepository` interface, remaining completely decoupled from the vector storage vendor:

- `PGVector` (default PostgreSQL pgvector integration)
- `Qdrant` · `Weaviate` · `Pinecone` · `Milvus` driver adapters
- `InMemoryVectorRepository` (thread-safe driver with Cosine Similarity math for unit testing)

### Intelligent Chunking Strategies

Avoids naive fixed-character splits by providing 5 strategy implementations:

| Strategy | Behavior |
|---|---|
| `SEMANTIC` | Paragraph & sentence-aware boundary splits |
| `HEADING_AWARE` | Markdown `#` heading splits with breadcrumb heading path (`# Title > ## Section`) |
| `PAGE_AWARE` | Page break (`---` / `\f` / `Page N`) splits preserving page numbers |
| `HIERARCHICAL` | Parent-child hierarchical chunk tree splits |
| `DOCUMENT_AWARE` | Structured knowledge component boundary splits |

### Database Schema (Migration 015)

| Table | Purpose |
|---|---|
| `embeddings` | Core stored vectors, model metadata, content hashes, active status |
| `embedding_chunks` | Intelligent document chunk vectors, token counts, heading paths |
| `embedding_versions` | Version & upgrade tracking (`v1.0.0` → `v2.0.0`) |
| `embedding_jobs` | Background batch vectorization job log |
| `embedding_providers` | Vault provider configurations & active default models |
| `vector_metadata` | Key-value tags & attributes linked to vectors |

### REST API

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/api/v1/embedding/generate` | Generate embedding for a structured representation |
| `POST` | `/api/v1/embedding/refresh` | Trigger incremental refresh of embeddings |
| `GET` | `/api/v1/embedding/status` | Check embedding pipeline status & active model |
| `GET` | `/api/v1/embedding/stats` | Vault embedding & chunk statistics |
| `GET` | `/api/v1/embedding/history` | Generation job log & model upgrade history |
| `PUT` | `/api/v1/embedding/provider` | Update provider configuration for vault |
| `POST` | `/api/v1/embedding/rebuild` | Upgrade model version and re-embed vault items |


