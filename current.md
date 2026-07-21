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

