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
