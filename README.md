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

---

## 🔍 Hybrid Knowledge Search Engine

Declutr's Hybrid Knowledge Search Engine is the unified retrieval layer of the platform. Rather than relying on simple keyword or semantic search alone, it dynamically combines 7 parallel retrieval strategies to find relevant knowledge regardless of how the user expresses their intent.

> **Retrieval Pipeline**: `User Query` → `Query Parser` → `Search Planner` → (`Keyword Search`, `Metadata Search`, `Entity Search`, `Relationship Search`, `Context Search`, `Memory Search`, `Vector Search`) → `Result Fusion` → `Ranking Engine` → `Search Results`

### Multi-Strategy Retrieval & Weighted Score Fusion

The engine evaluates candidates across 7 dimensions and combines them using weighted fusion:

$$\text{FinalScore} = w_{\text{kw}} S_{\text{kw}} + w_{\text{vec}} S_{\text{vec}} + w_{\text{ent}} S_{\text{ent}} + w_{\text{ctx}} S_{\text{ctx}} + w_{\text{rel}} S_{\text{rel}} + w_{\text{mem}} S_{\text{mem}} + w_{\text{rec}} S_{\text{rec}}$$

- **Keyword Search**: Full Text Search with prefix and quoted exact matching
- **Vector Search**: Dense vector semantic similarity via the Embedding Engine
- **Entity Search**: Canonical entity and alias matching (`Tokyo`, `PyTorch`, `Dr. Sharma`)
- **Context Search**: Intent & activity matching (`Travel`, `Research`, `Medical`)
- **Relationship Search**: Knowledge graph edge matching
- **Memory Search**: Long-term persistent knowledge & strength scoring
- **Recency Decay**: Exponential recency decay scoring

### Query Understanding & Parsing

Automatically detects query intent and structured constraints:
- Quoted exact terms (`"passport photo"`)
- Excluded terms (`-draft`)
- File types (`pdf`, `docx`, `png`, `mp4`)
- Year-based and relative date ranges (`2025`)
- Entity and location detection (`Tokyo`, `Japan`, `PyTorch`, `Cardiology`)

### Complete Match Explainability

Every search result explains **why** it matched:

```json
{
  "assetId": "asset-passport-001",
  "score": 0.94,
  "whyMatched": "Matched via exact keyword match in title & matched entity (Tokyo, Japan, Passport) & high semantic similarity.",
  "contributingStrategies": ["KEYWORD", "VECTOR", "ENTITY", "CONTEXT", "MEMORY"],
  "matchedEntities": ["Tokyo", "Japan", "Passport"],
  "matchedContexts": ["Japan Vacation"],
  "relatedMemories": ["Japan Vacation 2025"]
}
```

### Database Schema (Migration 016)

| Table | Purpose |
|---|---|
| `search_history` | Audit log of user search queries, latency, and result counts |
| `saved_searches` | Bookmarked search queries with custom filters & pin status |
| `search_statistics` | Vault-level search analytics, top queries, and strategy usage |
| `search_preferences` | Per-vault ranking weights and default search options |
| `search_index_versions` | Index synchronization and version tracking |

### REST API

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/api/v1/search/query` | Execute multi-strategy hybrid search query |
| `POST` | `/api/v1/search/saved` | Save / bookmark a search query with filters |
| `GET` | `/api/v1/search/saved` | List saved searches for a vault |
| `DELETE` | `/api/v1/search/saved` | Delete a saved search query |
| `GET` | `/api/v1/search/history` | Get recent query execution history |
| `GET` | `/api/v1/search/suggestions` | Search-as-you-type autocomplete suggestions |
| `GET` | `/api/v1/search/stats` | Get search engine metrics & top queries |
| `GET` / `PUT` | `/api/v1/search/preferences` | Get / update vault ranking weights |

---

## 📈 Knowledge Insights & Timeline Intelligence Engine

Declutr's Knowledge Insights & Timeline Engine proactively organizes and presents a user's digital life without waiting for manual searches. It transforms static stored documents into living chronological timelines, milestone alerts, and proactive pattern insights.

> **Architecture**: `Assets` → `Metadata` → `Entities` → `Relationships` → `Contexts` → `Memory` → `Hybrid Search` → **`Knowledge Insights Engine`** → `Timeline & Insights`

### Timeline Engine
Automatically organizes vault documents and activities into chronological timelines across 10 categories:
- ✈️ **Travel Events** (`Japan Vacation 2025 Flight Booking`)
- 🎓 **Education Events** (`PhD Thesis Chapter 4 Finalized`)
- 🏥 **Medical History** (`Cardiology Visit with Dr. Sharma`)
- 💼 **Financial & Tax Events** (`Annual Tax Return 2024 Filed`)
- 📁 **Projects, Legal, Purchases, Subscriptions & Custom Contexts**

### Proactive Insights & Pattern Detection
The engine background scanner continuously evaluates vault knowledge to generate actionable insights:
- **Upcoming Expirations**: Passport, Visa, Insurance renewal warnings
- **Recurring Expenses**: 30-day medication cycles, monthly subscriptions
- **Frequent Places**: Top referenced travel destinations (`Tokyo, Japan`)
- **Important & Missing Documents**: Document completeness checks
- **Trending Interests & Knowledge Growth**: Interest evolution analytics

### Milestone Detection
Tracks critical life, medical, travel, and project milestones:
- `US Passport Renewal Due` (`UPCOMING`)
- `Tax Return Filing 2024` (`COMPLETED`)
- `PhD Thesis Final Submission` (`PENDING`)

### Complete Insight Explainability
Every proactive insight provides explicit rationale and supporting evidence:
```json
{
  "insightId": "ins-001",
  "insightType": "EXPIRATION_WARNING",
  "title": "Passport Renewal Needed Soon",
  "summary": "Your US Passport expires in 65 days. Renewal recommended before upcoming travel.",
  "whyGenerated": "Passport expiration date detected in document 'Japanese Visa & Passport Scan' (expires Sep 2025).",
  "evidence": ["Asset: Japanese Visa & Passport Scan", "Expiration Date: 2025-09-25"],
  "importance": 0.95,
  "confidence": 0.98
}
```

### Database Schema (Migration 017)

| Table | Purpose |
|---|---|
| `timeline_events` | Core chronological event feed across life, projects, travel, medical |
| `timeline_groups` | Sequence groups linking related timeline events into context streams |
| `knowledge_insights` | Proactive automated insights with evidence and dismissal state |
| `insight_history` | Audit log of generated, dismissed, and actioned insights |
| `insight_preferences` | Per-vault enabled insight types and min confidence settings |
| `milestones` | Detected major milestones with status and due dates |

### REST API

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/api/v1/insights/timeline` | Get chronological timeline events (filtered by `eventType`) |
| `GET` | `/api/v1/insights` | Get active non-dismissed proactive insights |
| `GET` | `/api/v1/insights/milestones` | Get detected vault milestones |
| `POST` | `/api/v1/insights/dismiss` | Dismiss an insight by `insightId` |
| `POST` | `/api/v1/insights/refresh` | Trigger incremental background intelligence refresh |
| `GET` | `/api/v1/insights/stats` | Vault insight and timeline metrics |
| `GET` / `PUT` | `/api/v1/insights/preferences` | Get / update insight preferences |

---

## 🤖 Declutr AI Copilot (Grounded RAG & Personal Intelligence)

Declutr AI Copilot is the personal intelligence layer built on top of the user's digital vault. Unlike generic chatbots, every AI response is strictly grounded in retrieved vault documents, memories, context, and timeline events.

> **RAG Pipeline**: `User Question` → `Intent Parser` → `Hybrid Search Integration` → `Context Builder` → `Prompt Builder` → `LLM Synthesis` → `Grounded Response` → `Citations & Evidence`

### Grounded Answer Synthesis & Zero-Hallucination Policy
- Answers are synthesized **only** when supporting vault evidence is retrieved.
- If evidence is absent, the AI explicitly states: *"I searched your vault for relevant records, but could not find sufficient grounded evidence..."*
- Every assistant message is paired with verifiable document citations, snippets, confidence scores, and reasoning overviews.

```json
{
  "messageId": "msg-asst-1",
  "role": "ASSISTANT",
  "content": "Based on your vault document 'Japanese Visa & Passport Scan' (PDF), your passport expires on September 25, 2025 (in 65 days).",
  "confidence": 0.96,
  "reasoningOverview": "Grounded via exact entity match (Tokyo, Japan, Passport) and semantic vector search over document asset-passport-001.",
  "citations": [
    {
      "assetId": "asset-passport-001",
      "title": "Japanese Visa & Passport Scan",
      "assetType": "PDF",
      "snippet": "Passport number A987654321. Expiration Date: 2025-09-25.",
      "confidence": 0.98,
      "matchedEntities": ["Tokyo", "Japan", "Passport"]
    }
  ]
}
```

### Database Schema (Migration 018)

| Table | Purpose |
|---|---|
| `conversations` | Multi-turn user RAG conversation session tracking |
| `messages` | Grounded user and assistant messages with token usage & confidence |
| `conversation_context` | Auditable RAG context snapshots passed into prompt builder |
| `conversation_feedback` | User upvote/downvote ratings on response quality |
| `prompt_versions` | Versioned system & context prompts |
| `response_history` | Audit log of latency, model versions, and response metrics |

### REST API

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/api/v1/copilot/conversations` | Start a new RAG conversation session |
| `GET` | `/api/v1/copilot/conversations` | List conversation history sessions for vault |
| `DELETE` | `/api/v1/copilot/conversations` | Delete a conversation session |
| `POST` | `/api/v1/copilot/messages` | Send message & receive grounded answer with citations |
| `GET` | `/api/v1/copilot/messages` | Get message history for a conversation session |
| `POST` | `/api/v1/copilot/feedback` | Submit user rating on AI grounding quality |
| `GET` | `/api/v1/copilot/messages/stream` | Stream response via Server-Sent Events (SSE) |

---

## ⚡ Workflow Automation & Intelligent Actions Engine

Declutr's Workflow Automation Engine allows users to construct intelligent, event-driven rules within their vault. It automatically triggers actions based on internal vault events (uploads, AI analysis, entity discovery, expirations, schedules) without requiring external dependencies.

> **Architecture**: `Event` → `Trigger Engine` → `Condition Engine` → `Rule Engine` → `Action Engine` → `Execution` → `History & Observability`

### Triggers, Conditions & Actions

- **12 Internal Event Triggers**: `ASSET_UPLOADED`, `ASSET_UPDATED`, `ASSET_DELETED`, `CONTEXT_CREATED`, `MEMORY_CREATED`, `ENTITY_FOUND`, `RELATIONSHIP_CREATED`, `DOCUMENT_EXPIRING`, `DAILY_SCHEDULE`, `MANUAL_TRIGGER`, `AI_INSIGHT_CREATED`, `TIMELINE_EVENT`
- **Rule Evaluator**: Evaluates AND, OR, NOT condition rules across `fileType`, `entity`, `context`, `confidence`, `date`, and `storageSize`.
- **Executable Actions**: `APPLY_TAGS`, `CREATE_COLLECTION`, `MOVE_ASSET`, `GENERATE_SUMMARY`, `ARCHIVE_ASSET`, `CREATE_REMINDER`, `PIN_MEMORY`, `REFRESH_SEARCH_INDEX`, `NOTIFY_USER`.

### Database Schema (Migration 019)

| Table | Purpose |
|---|---|
| `workflows` | Core workflow definition, enabled status, and aggregate statistics |
| `workflow_triggers` | Event trigger configuration linked to workflows |
| `workflow_conditions` | Rule condition evaluation criteria (`field`, `operator`, `value`, `combinator`) |
| `workflow_actions` | Sequenced executable actions (`action_type`, `config`, `execution_order`) |
| `workflow_runs` | Individual execution run records (`status`, `duration_ms`, `error_message`) |
| `workflow_logs` | Granular step-by-step execution log entries |
| `workflow_history` | Vault-level historical execution statistics and success rates |

### REST API

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/api/v1/workflows` | Create a new workflow rule definition |
| `GET` | `/api/v1/workflows` | List all workflows for a vault |
| `PUT` | `/api/v1/workflows` | Update workflow definition |
| `DELETE` | `/api/v1/workflows` | Delete workflow definition |
| `POST` | `/api/v1/workflows/toggle` | Enable or disable workflow |
| `POST` | `/api/v1/workflows/run` | Manually trigger workflow execution |
| `GET` | `/api/v1/workflows/history` | Get run history and step logs |
| `GET` | `/api/v1/workflows/stats` | Vault workflow observability metrics (success rate, avg duration) |

---

## 🔔 Notification Center & Proactive Intelligence

Declutr's Notification Center & Proactive Intelligence system delivers contextual, explainable, and actionable alerts to users across expirations, workflows, security events, AI insights, and memory discoveries.

> **Architecture**: `Domain Event` → `Notification Rules` → `Priority Engine` → `Delivery Scheduler` → `Notification Center` → `User Action`

### Priority Engine & Actionable Alerts

- **Priority Levels**: `LOW`, `MEDIUM`, `HIGH`, `URGENT` calculated dynamically using event type, importance, urgency, and recency.
- **Actionable Buttons**: Open Asset, View Context, Run Workflow, Retry Job, Dismiss, Pin, Archive, Snooze.
- **Proactive Digests**: Automated Daily Summaries and Weekly Recaps tracking knowledge growth, new memories, document expirations, and workflow activity.

### Database Schema (Migration 020)

| Table | Purpose |
|---|---|
| `notifications` | Core notification alerts, types, priority levels, read/dismiss status |
| `notification_rules` | Custom alert rules matching domain event triggers |
| `notification_preferences` | Vault channel settings (`IN_APP`, `EMAIL`, `PUSH`, `DESKTOP`) and digest frequency |
| `notification_delivery` | Log of channel delivery states |
| `notification_history` | Audit log of user interactions (`READ`, `DISMISSED`, `ACTIONED`, `SNOOZED`) |
| `digest_reports` | Daily and Weekly proactive intelligence summary reports |

### REST API

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/api/v1/notifications` | List notifications for vault (filtered by status, priority) |
| `POST` | `/api/v1/notifications/read` | Mark specific notification(s) or all as read |
| `POST` | `/api/v1/notifications/dismiss` | Dismiss a notification |
| `POST` | `/api/v1/notifications/action` | Execute actionable notification step |
| `GET` | `/api/v1/notifications/digests` | Get Daily and Weekly digest reports |
| `GET` / `PUT` | `/api/v1/notifications/preferences` | Get / update channel preferences and digest frequency |
| `GET` | `/api/v1/notifications/stats` | Vault notification stats (unread count, urgent count, read rate) |

---

## 🔒 Secure Sharing & Collaboration Platform

Declutr's Secure Sharing & Collaboration Platform enables privacy-first, granular, revocable, and auditable resource sharing across Assets, Folders, Collections, Context Streams, Projects, Timeline Views, and Search Results.

> **Architecture**: `User` → `Permission Engine` → `Share Manager` → `Access Validation` → `Audit System` → `Shared Resource`

### Role Hierarchy & Permissions

- **Member Roles**: `READ_ONLY` (Viewer), `COMMENT_ONLY` (Commenter), `EDIT` (Editor), `OWNER` / `CO_OWNER` (Full Control).
- **Access Modes**: `PRIVATE`, `INVITE_ONLY`, `LINK_SHARING` (password-protected, expiration, download limits), `TEMPORARY_ACCESS`.
- **Threaded Discussions**: Inline comments, replies, mentions, and resolution tracking.
- **Auditable History**: Activity log tracking views, downloads, edits, comments, shares, permission changes, and revocations.

### Database Schema (Migration 021)

| Table | Purpose |
|---|---|
| `shares` | Core shared resource container and access settings |
| `share_permissions` | Role permission matrices (`can_view`, `can_edit`, `can_comment`, `can_share`, `can_manage_members`) |
| `share_members` | Explicit member email list, roles, and membership expiration |
| `share_links` | Link sharing tokens, password hashes, view & download limits |
| `share_comments` | Threaded discussion comments, replies, and resolution state |
| `share_activity` | Auditable collaboration history log |
| `share_invitations` | Pending email invitations with accept/reject tokens |

### REST API

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/api/v1/shares` | Create a shared resource container |
| `GET` | `/api/v1/shares` | List all shares for a vault |
| `DELETE` | `/api/v1/shares` | Revoke a share and all access |
| `POST` | `/api/v1/shares/invite` | Send invitation to user or email |
| `POST` | `/api/v1/shares/invite/accept` | Accept an invitation token |
| `POST` | `/api/v1/shares/links` | Create password-protected share link |
| `POST` | `/api/v1/shares/links/revoke` | Revoke a share link token |
| `POST` | `/api/v1/shares/comments` | Add threaded comment or reply |
| `GET` | `/api/v1/shares/comments` | List comments for a share |
| `GET` | `/api/v1/shares/activity` | Get audit activity history for share |
| `GET` | `/api/v1/shares/stats` | Vault collaboration stats (active shares, members, comments) |

---

## 🕒 Version History, Recovery & Time Machine

Declutr's Version History & Recovery System ("Time Machine") captures immutable and incremental snapshots of assets, metadata, AI analysis, contexts, relationships, collections, memories, workflows, and preferences, allowing safe field-level diff inspection, point-in-time restoration, and soft-delete recovery.

> **Architecture**: `Asset Change` → `Version Manager` → `Snapshot Generator` → `Version Store` → `Recovery Engine` → `Restore`

### Versioning & Snapshot Strategy

- **Versioned Resources**: Assets, Metadata, AI Analysis, Entity Extraction, Relationships, Contexts, Collections, Workflows, Vault Settings, Preferences, Search Index Metadata.
- **Snapshot Modes**: `FULL`, `INCREMENTAL`, `DELTA`, `COMPRESSED`, `IMMUTABLE`.
- **Field Diff Engine**: Computes added, removed, and modified key-value pairs between any two version snapshots.
- **Recycle Bin (Soft Delete)**: Manages soft-deleted vault items with retention expiration, bulk restoration, and permanent purge actions.

### Database Schema (Migration 022)

| Table | Purpose |
|---|---|
| `resource_versions` | Core version records, resource type, version numbers, checksums |
| `version_snapshots` | Snapshot payloads (`FULL`, `DELTA`, `COMPRESSED`, `IMMUTABLE`) |
| `change_history` | Audit log of field-level modifications and actors |
| `recycle_bin` | Soft-deleted items, original path, retention expiration |
| `restore_jobs` | Tracking restoration job executions and status |
| `version_diffs` | Cached version comparison diff payloads |

### REST API

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/api/v1/versions` | List version timeline for resource or vault |
| `POST` | `/api/v1/versions/snapshot` | Capture a new version snapshot |
| `POST` | `/api/v1/versions/compare` | Compare two version snapshots and generate field-level diff |
| `POST` | `/api/v1/versions/restore` | Restore resource state to target version |
| `GET` | `/api/v1/recyclebin` | List soft-deleted items in Recycle Bin |
| `POST` | `/api/v1/recyclebin/restore` | Restore soft-deleted item(s) |
| `DELETE` | `/api/v1/recyclebin/purge` | Permanently purge soft-deleted item(s) |
| `GET` | `/api/v1/versions/stats` | Vault versioning & time machine metrics |

---

## 📦 Backup, Disaster Recovery & Business Continuity

Declutr's Backup & Disaster Recovery System protects the entire Vault from catastrophic data loss across personal and enterprise deployments using encrypted full/incremental snapshots, automated retention schedules, SHA-256 integrity validation, and catastrophe recovery.

> **Architecture**: `Vault` → `Backup Scheduler` → `Snapshot Engine` → `Backup Storage` → `Integrity Verification` → `Recovery Manager` → `Restore`

### Backup Types & Disaster Recovery Modes

- **Backup Types**: `MANUAL`, `SCHEDULED`, `INCREMENTAL`, `FULL`, `ENCRYPTED`, `OFFLINE`, `COLD_STORAGE`.
- **Content Scope**: Assets, Metadata, AI Analysis, Entities, Relationships, Contexts, Memories, Embeddings, Workflows, Notifications, Preferences, Vault Settings, Version History, Audit Logs.
- **Recovery Scopes**: `FULL_VAULT`, `SELECTIVE`, `ASSETS_ONLY`, `METADATA_ONLY`, `AI_STATE_ONLY`, `WORKFLOWS_ONLY`, `SETTINGS_ONLY`.
- **Merge Strategies**: `OVERWRITE_EXISTING`, `RESTORE_AS_NEW_VAULT`, `MERGE_RESTORE`, `DRY_RUN`.

### Database Schema (Migration 023)

| Table | Purpose |
|---|---|
| `backups` | Core backup package records, types, size, checksums |
| `backup_jobs` | Tracking background backup, restore, and integrity validation jobs |
| `backup_files` | Individual files inside backup payload |
| `backup_manifests` | Manifest data tracking vault contents, total assets, memories, workflows |
| `backup_history` | Audit log of backup creation, verification, and purging events |
| `restore_jobs` | Tracking disaster recovery execution jobs and recovery mode |
| `restore_history` | Audit log of disaster recovery restores |

### REST API

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/api/v1/backups` | Create manual snapshot backup |
| `GET` | `/api/v1/backups` | List backups for vault |
| `GET` | `/api/v1/backups/detail` | Get backup details and manifest |
| `POST` | `/api/v1/backups/schedule` | Configure automated backup schedule & retention policy |
| `POST` | `/api/v1/backups/restore` | Trigger disaster recovery vault restoration |
| `POST` | `/api/v1/backups/verify` | Verify backup SHA-256 checksum & manifest integrity |
| `POST` | `/api/v1/backups/cancel` | Cancel active backup or restore job |
| `GET` | `/api/v1/backups/stats` | Vault backup & disaster recovery stats |










