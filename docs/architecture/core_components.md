# Declutr Architecture — Core Components & Module Contracts

> **Source of Truth:** [declutr_architecture_document.html](file:///f:/Github/Declutr/docs/architecture/declutr_architecture_document.html)  
> **Section:** 4. Core System Components

---

## 1. Modular Monolith Design

Declutr is engineered as a clean **Modular Monolith** backend in Go, preserving strict domain boundaries while avoiding premature microservice distribution. Each module enforces strict boundary contracts and interacts through in-process service injections.

```
backend/modules/
├── auth/          # SRP-6a verification, JWT creation, session lifecycles
├── vault/         # Vault workspace scopes, key wrapping (VK with MVK)
├── file/          # Digital item metadata, version pointers, direct S3 pre-signed links
├── search/        # Hybrid keyword (Postgres FTS) + semantic (pgvector) query execution
├── persona/       # Probabilistic Reverse Persona weights & recency decay evaluator
└── behavior/      # Session anomaly detection, risk scoring, adaptive auth triggers
```

---

## 2. Component Responsibility Matrix

| Module Name | Domain Responsibility | Key Interface Points |
| :--- | :--- | :--- |
| **Auth Module** | Handles SRP-6a credential checks, JWT generation, passkey setups, and session lifetimes. | Vault Module, Behavioral Risk Module |
| **Vault Module** | Manages physical vault scopes, shared crypt-keys, and folder hierarchies. | Auth Module, File Module |
| **File / Upload Module** | Coordinates chunk allocations, S3 pre-signed upload URLs, and commit checks. | Vault Module, Background Jobs |
| **Content Module** | Ingests digital items, detects file boundaries, and triggers the intelligence pipeline. | Upload Module, AI Processing |
| **AI Processing Module** | Orchestrates feature extraction, OCR execution, embeddings, and classification. | Content Module, Metadata Module |
| **Metadata Module** | Stores and manages client-encrypted, system, and AI-generated metadata. | Content Module, Search Module |
| **Intent Module** | Analyzes items to classify their context-utility and functional intent. | Metadata Module, Persona Module |
| **Relationship Module** | Links items via relational references and contexts. | Intent Module, Search Module |
| **Persona Module** | Constructs the Reverse Persona by mapping user interactions and actions. | Auth Module, Intent Module, Search Module |
| **Behavioral Risk Module** | Monitors traffic profiles and session metrics to calculate adaptive threat scores. | Auth Module, Audit Module |
| **Search Module** | Executes hybrid query parsing, vector matching, and metadata filtering. | Metadata Module, Relationship Module |
| **Audit Module** | Appends immutable transaction actions to the HMAC chain. | All Modules |

---

## 3. Strict Module Boundaries Rules

1. **No Direct Database Access Across Modules:** Module `A` must never execute raw SQL against tables owned by Module `B`.
2. **Interface Ingestion:** Interactions must occur via domain interfaces or exported application services.
3. **DTO Isolation:** Transport DTOs must remain in the transport layer; domain models represent internal business state.
