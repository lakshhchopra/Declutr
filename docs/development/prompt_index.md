# Implementation Prompt Index

> **Source of Truth:** [declutr_architecture_document.html](file:///f:/Github/Declutr/docs/architecture/declutr_architecture_document.html)  
> **Section:** 28. Feature-Wise Implementation Prompt Index

---

## Architecture Validation Matrix

| Decision Point | Chosen Approach | Architectural Justification | MVP Status |
| :--- | :--- | :--- | :--- |
| **System Layout** | Go Modular Monolith | Preserves clean domain boundaries without microservice overhead. | Target |
| **Primary Storage** | PostgreSQL 16 + pgvector | Consolidates relational data, metadata indexes, and 512-dim vectors. | Target |
| **Task Scheduling** | Redis + Asynq / Worker Queue | Enables reliable, asynchronous job processing and retries. | Target |
| **Default Privacy** | Private Mode (Client-Side) | Zero-knowledge isolation by executing encryption & classification locally. | Target |
| **Advanced AI** | Enhanced AI Mode (Opt-In) | Opt-in server processing for advanced multimodal AI features. | Target |

---

## Implementation Prompts Directory

| ID | Module Name | Objective | Dependencies | Deliverable |
| :---: | :--- | :--- | :--- | :--- |
| **01** | Project Foundation | Monorepo scaffolding, config layers, and build pipelines. | None | Modular monorepo scaffold. |
| **02** | Database Setup | Initialize PostgreSQL schema with pgvector extensions. | PROMPT 01 | Migration scripts & models. |
| **03** | SRP-6a Auth | Secure Remote Password registration & login flows. | PROMPT 02 | SRP authentication endpoints. |
| **04** | Vault Scope | Cryptographic workspaces and key wrapping handlers. | PROMPT 03 | Vault CRUD & VK key wrapping. |
| **05** | Object Storage | Chunked, direct-to-S3 uploads with pre-signed URLs. | PROMPT 04 | Pre-signed URL generation API. |
| **06** | Job Infrastructure | Redis background task worker queue setup. | PROMPT 05 | Asynchronous worker pipelines. |
| **07** | Ingestion Pipeline | OCR, text extraction routines, and mime validation. | PROMPT 06 | Content ingestion processors. |
| **08** | Metadata Layer | Sealed metadata store and client-encrypted indexing. | PROMPT 07 | Encrypted metadata API. |
| **09** | Semantic Vector | Embeddings generation and pgvector indexing. | PROMPT 08 | 512-dim vector embedding store. |
| **10** | Intent Engine | Classify file utility, category, and functional intent. | PROMPT 09 | Intent classification engine. |
| **11** | Relationship Links | Map relationships and context groups in PostgreSQL. | PROMPT 10 | Context & relational link APIs. |
| **12** | Hybrid Search | Combine FTS keyword search and pgvector similarity. | PROMPT 11 | Hybrid retrieval endpoint. |
| **13** | Reverse Persona | Process user interaction events for recency decay persona. | PROMPT 12 | Persona profile engine. |
| **14** | Adaptive Risk | Passive behavioral event tracking & session risk scoring. | PROMPT 03 | Session risk scoring engine. |
| **15** | Adaptive Auth | Intercept high-risk requests with MFA challenges. | PROMPT 14 | WebAuthn MFA interceptor. |
| **16** | Audit Logging | Append-only HMAC-chained audit log. | PROMPT 02 | Immutable audit log verification. |
| **17** | Observability | OpenTelemetry structured logging and tracing. | PROMPT 01 | Telemetry instrumentation. |
| **18** | Production Deploy | Docker, Helm charts, and CI/CD pipelines. | All Prompts | Production deploy configs. |
