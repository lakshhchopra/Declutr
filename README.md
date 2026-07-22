# Declutr

> Your personal knowledge operating system — private, AI-powered, and built for the way you actually think.

Declutr lets you store every document, note, and file in an end-to-end encrypted Vault, then gives you an AI copilot that can search, summarize, and reason over everything you've ever saved — across every Life Area, Project, and Goal in your life.

---

## Features

| Capability | Description |
|---|---|
| 🔒 **Zero-Knowledge Vault** | SRP-based auth, client-side encrypted storage |
| 🤖 **Grounded AI Copilot** | RAG-powered chat with citations from your own files |
| 🧠 **Semantic Search** | Hybrid keyword + vector search across all assets |
| 🗂️ **Life Operating System** | Life Areas → Projects → Goals instead of plain folders |
| 🔮 **Predictive Intelligence** | Proactive suggestions before you ask |
| 🤝 **Multi-Agent Platform** | Specialist AI agents that collaborate on complex tasks |
| 🔗 **Integrations** | Google Drive, Notion, GitHub, Dropbox connectors |
| 🏢 **Organizations** | Multi-tenant workspaces with role-based access |
| 🔧 **Developer Platform** | REST API, Webhooks, OAuth2, and SDK |
| 🧩 **Extension Marketplace** | Sandboxed third-party extensions |

---

## Tech Stack

| Layer | Technology |
|---|---|
| Frontend | Next.js 14 (App Router), TypeScript, Tailwind CSS, shadcn/ui |
| Mobile | React Native + Expo |
| Backend | Go (net/http), Clean Architecture |
| Database | PostgreSQL 16 |
| Cache | Redis 7 |
| AI | OpenAI / Gemini API (pluggable providers) |
| Storage | Cloudflare R2 / S3-compatible |
| Auth | SRP (Secure Remote Password) — zero-knowledge |

---

## Repository Structure

```
declutr/
├── backend/              # Go REST API server
│   ├── cmd/server/       # Entry point (main.go)
│   ├── modules/          # Feature modules (auth, vault, ai, search, lifeos, …)
│   ├── shared/           # Config, middleware, DB helpers, observability
│   └── tests/            # Integration & unit tests
├── database/
│   └── migrations/       # 10 grouped SQL migrations (001–010)
├── docs/
│   └── declutr_architecture_document.html  # Interactive architecture docs
├── frontend/             # Next.js web app
│   ├── app/              # App Router pages
│   ├── features/         # Feature components (auth, vault, copilot, lifeos, …)
│   ├── components/       # Shared UI components
│   ├── hooks/            # Custom React hooks
│   ├── services/         # API client functions
│   ├── stores/           # Zustand state stores
│   └── styles/           # Global styles
├── frontend/declutr-mobile/  # React Native + Expo mobile app
├── docker-compose.yml    # One-command local dev environment
├── .env.example          # All environment variables documented
└── README.md
```

---

## Quick Start

### Prerequisites

- **Go** 1.22+
- **Node.js** 18+
- **Docker** (for local Postgres + Redis)

### 1. Clone

```bash
git clone https://github.com/lakshhchopra/declutr.git
cd declutr
```

### 2. Set up environment

```bash
cp .env.example .env
# Edit .env with your API keys
```

### 3. Start database & cache

```bash
docker compose up -d db redis
```

### 4. Run backend

```bash
cd backend
go run cmd/server/main.go
# API available at http://localhost:8080
```

### 5. Run frontend

```bash
cd frontend
npm install
npm run dev
# App available at http://localhost:3000
```

### 6. Run mobile (optional)

```bash
cd frontend/declutr-mobile
npm install
npx expo start
```

### One-command (Docker)

```bash
docker compose up
# Frontend → http://localhost:3000
# Backend  → http://localhost:8080
```

---

## Environment Variables

Copy `.env.example` to `.env`. All variables are documented in [`docs/declutr_architecture_document.html`](docs/declutr_architecture_document.html).

| Variable | Description |
|---|---|
| `DATABASE_URL` | PostgreSQL connection string |
| `REDIS_URL` | Redis connection string |
| `OPENAI_API_KEY` | OpenAI API key for AI features |
| `GEMINI_API_KEY` | Google Gemini API key (alternative) |
| `R2_BUCKET` | Cloudflare R2 bucket name |
| `R2_ACCESS_KEY_ID` | Cloudflare R2 access key |
| `R2_SECRET_ACCESS_KEY` | Cloudflare R2 secret |
| `JWT_SECRET` | JWT signing secret |
| `PORT` | Backend server port (default: 8080) |

---

## Database Migrations

Migrations are in `database/migrations/` and run automatically in Docker. They are grouped into 10 logical domains:

| File | Domain |
|---|---|
| `001_users_and_auth.sql` | Users, SRP credentials |
| `002_sessions_and_security.sql` | Sessions, audit events, trusted devices |
| `003_vaults_and_assets.sql` | Vaults, assets, versions, recycle bin |
| `004_processing_and_embeddings.sql` | OCR/extraction jobs, vector embeddings |
| `005_ai_copilot_and_memory.sql` | AI conversations, episodic memory, persona |
| `006_search_and_insights.sql` | Saved searches, search history, knowledge insights |
| `007_workflows_and_notifications.sql` | Automations, workflow executions, notifications |
| `008_organizations_and_sharing.sql` | Orgs, members, shared assets |
| `009_developer_and_extensions.sql` | API keys, webhooks, marketplace extensions |
| `010_intelligence_agents_predictive_lifeos.sql` | Agents, predictions, Life Areas, Projects, Goals |

---

## Deployment

### Backend → [Railway](https://railway.app) / [Render](https://render.com) / [Fly.io](https://fly.io)

```bash
# Railway example
railway login
railway up
```

### Frontend → [Vercel](https://vercel.com) / [Netlify](https://netlify.com)

```bash
# Vercel example
cd frontend
vercel deploy
```

### Database → [Supabase](https://supabase.com) / [Neon](https://neon.tech)

Create a project and copy the connection string into `DATABASE_URL`.

### Storage → [Cloudflare R2](https://developers.cloudflare.com/r2/)

Create a bucket and set `R2_BUCKET`, `R2_ACCESS_KEY_ID`, `R2_SECRET_ACCESS_KEY`.

### Cache → [Upstash Redis](https://upstash.com)

Create a Redis database and copy the URL into `REDIS_URL`.

---

## API

The REST API is available at `http://localhost:8080/api/v1/`. Full documentation is in [`docs/declutr_architecture_document.html`](docs/declutr_architecture_document.html).

Key endpoint groups:

- `/api/v1/auth/*` — Authentication (register, login, session)
- `/api/v1/vaults/*` — Vault and asset management
- `/api/v1/copilot/*` — AI Copilot conversations
- `/api/v1/search/*` — Hybrid knowledge search
- `/api/v1/lifeos/*` — Life Operating System (areas, projects, goals)
- `/api/v1/predictive/*` — Predictive intelligence feed
- `/api/v1/agents/*` — Autonomous agent platform
- `/api/v1/multiagent/*` — Multi-agent coordinator
- `/api/v1/developer/*` — Public developer platform (API keys, webhooks)
- `/api/v1/organizations/*` — Enterprise multi-tenancy
- `/api/v1/extensions/*` — Extension marketplace

---

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feat/my-feature`
3. Commit: `git commit -m 'feat: add my feature'`
4. Push: `git push origin feat/my-feature`
5. Open a Pull Request

See [CONTRIBUTING.md](CONTRIBUTING.md) for full guidelines.

---

## Architecture

Full interactive architecture documentation is available at [`docs/declutr_architecture_document.html`](docs/declutr_architecture_document.html).

It covers: Architecture overview, folder structure, database schema, authentication flow, upload pipeline, AI pipeline, search pipeline, deployment guide, environment variables, roadmap, and current progress.

---

## License

MIT — see [LICENSE](LICENSE).
