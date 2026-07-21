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

- **Shared Design System Foundation (Issue #001)**:
  - Established centralized theme system (`ThemeProvider`) supporting Dark Mode (default), Light Mode, System Theme detection, and persistent `localStorage` preference.
  - Implemented semantic CSS design tokens in `globals.css` (Background, Surface, Card, Border, Primary, Secondary, Accent, Success, Warning, Danger, Info, Text tokens).
  - Built reusable `shadcn/ui` & Radix UI component primitives under `frontend/shared/components/`:
    - Buttons (Primary, Secondary, Ghost, Danger, Outline, Icon, Loading)
    - Inputs (TextInput, PasswordInput, SearchInput, Textarea)
    - Cards & Badges (Card, CardHeader, CardContent, Badge variants)
    - Feedback (Alerts, Spinner, Skeleton, EmptyState, ErrorState)
    - Overlay (Dialog / Modal)
    - Navigation (SidebarShell, TopNavigation, Breadcrumb, Tabs)
    - Layout (Container, Grid, Stack, Section)
  - Created interactive Design System Showcase page at `/design-system`.

