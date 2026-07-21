# ADR 0001: Modular Monolith Architecture

## Context
Declutr requires strong domain boundaries across authentication, vault encryption, content ingestion, intent processing, and hybrid retrieval. Distributed microservices introduce premature deployment complexity, network overhead, and distributed transaction challenges.

## Decision
Adopt a **Domain-Oriented Modular Monolith** pattern for the Go backend. Each module owns its `domain`, `application`, `repository`, and `transport` packages. Cross-module data access is prohibited except via exposed domain interfaces or application services.

## Status
Accepted

## Consequences
- High cohesion and low coupling across feature modules.
- Simplified single-binary deployment for MVP.
- Clean path for future microservice extraction if specific modules require GPU-backed scale-out.
