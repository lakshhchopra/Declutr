# Contributing to Declutr

Thank you for contributing to Declutr! This document outlines our monorepo architecture, coding standards, and development guidelines.

## Monorepo Architecture

Declutr is organized as a domain-oriented modular monolith across its backend, web, and mobile codebases:

```
Declutr Monorepo/
├── backend/                  # Go Modular Monolith Backend
│   ├── cmd/server/           # Application entrypoint
│   ├── modules/              # Feature modules (auth, vault, file, search, persona, behavior)
│   ├── shared/               # Cross-cutting concerns (crypto, database, middleware)
│   └── pkg/                  # Shared public utilities
├── frontend/                 # Next.js Web Client (TypeScript)
│   ├── app/                  # App router pages
│   ├── features/             # Feature modules (auth, vault, search)
│   ├── shared/               # UI components, hooks, providers, API services
│   └── declutr-mobile/       # React Native / Expo Mobile Application
│       ├── app/              # Expo router screens
│       ├── features/         # Mobile feature modules
│       ├── shared/           # Native UI components, providers, services
│       └── navigation/       # Navigation configurations
├── database/                 # Database migrations, seeds, and scripts
├── docs/                     # Documentation (architecture, api, adr, development, references)
├── infrastructure/           # Docker, Compose, K8s, Terraform, Monitoring
├── scripts/                  # Development, build, and maintenance scripts
├── security/                 # Security policies and audit documentation
└── tests/                    # Unit, Integration, and E2E test suites
```

## Module Contract & Dependency Rules

1. **Domain Ownership:** Every feature module must own its domain models, application services, repositories, and transport handlers.
2. **Strict Module Boundaries:** No feature module may directly import another feature module's private repository implementation. Communication must occur via explicit interfaces or application services.
3. **Clean Handler Layers:** HTTP handlers contain only request parsing and response rendering logic; business rules reside in application services.

## Development Workflow

1. Format code before committing:
   - Backend: `go fmt ./...`
   - Frontend: `npx prettier --write .`
2. Run test suites to ensure zero regressions:
   - Backend: `go test ./...`
   - Frontend: `npm run build`
