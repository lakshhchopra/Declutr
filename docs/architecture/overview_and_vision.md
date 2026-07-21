# Declutr Architecture — Vision & System Overview

> **Source of Truth:** [declutr_architecture_document.html](file:///f:/Github/Declutr/docs/architecture/declutr_architecture_document.html)  
> **Version:** 2.0  
> **Classification:** Technical Architecture Specification

---

## 1. Executive Summary

Declutr is an **AI-powered digital life vault** engineered to help users securely store, organize, connect, and retrieve the massive, fragmented volume of information they accumulate throughout their digital lives.

Traditional storage relies on explicit folder paths, rigid filenames, file extensions, and manually managed tags. Human memory, however, is contextual, associative, and narrative-driven. Users recall items not by explicit paths, but by real-world contexts and relationships (e.g., *"that hotel booking receipt I saved before my Mumbai trip"* or *"the document related to my summer internship"*).

Declutr shifts the digital storage paradigm from an encrypted file database to an **Intelligent Personal Digital Memory System**. It answers seven core semantic questions for every stored item:

1. **WHAT** the item contains and represents.
2. **WHY** the item matters to the user (its utility and value).
3. **WHAT** broader context or project the item connects to.
4. **WHAT** other files, events, or entities the item relates to.
5. **HOW** the user interacts with this information over time.
6. **HOW** the user is likely to remember and search for information.
7. **WHETHER** the current session suggests a security anomaly.

---

## 2. Core Product Intelligence Pillars

Declutr's architecture is built around five core intelligence pillars and a supporting behavioral security layer:

```
  +-----------------------------------------------------------------------+
  |                          DECLUTR VAULT CORE                           |
  +-----------------------------------------------------------------------+
      │                │                 │                │            │
      ▼                ▼                 ▼                ▼            ▼
 🧠 Content       🎯 Intent         🔗 Relationship  👤 Persona    🔍 Retrieval
 Intelligence   Intelligence        Intelligence   Intelligence   Intelligence
 (OCR/Vectors)   (Utility Tagging)  (Relational)    (User Model)   (Hybrid Search)
      │                │                 │                │            │
      +----------------+-----------------+----------------+------------+
                                       │
                                       ▼
                     🛡 Behavioral Authentication Security
```

1. **Content Intelligence:** Extracts structured semantic features, OCR text, document layout concepts, visual objects, and audio segments. Automatically categorizes and summarizes content upon ingestion.
2. **Intent Intelligence:** Determines the utility of an item (reference material, active planning document, legal archive, transaction record) to predict when, why, and how a user will retrieve it.
3. **Relationship Intelligence:** Dynamically associates independent digital artifacts. Detects temporal proximity, entity matching, and contextual links (e.g., linking a boarding pass, hotel receipt, and restaurant screenshot into a unified *"Trip"* context).
4. **Persona Intelligence:** Compiles behavior signals, search choices, and content patterns into a private, probabilistic **Reverse Persona** that guides contextual ranking and intent inference.
5. **Retrieval Intelligence:** Synthesizes traditional keyword search, metadata filters, semantic vectors, and relationship graphs to support natural-language query resolution and temporal reasoning.
6. **Behavioral Security:** Constructs passive behavioral baselines (session properties, access volume, download rates) to evaluate risk dynamically and trigger adaptive authentication challenges without disrupting user flow.

---

## 3. High-Level Data Flow

```
     [User / Client App]
            │
            ▼ (HTTPS / SRP Auth)
     [Authentication & Session Layer]
            │
            ▼
     [Digital Vault Scope (AES-GCM Encryption)]
            │
            ▼
     [Direct-to-S3 Chunked Upload Pipeline]
            │
            ▼ (Redis / Worker Task Queue)
     [Content Ingestion & Extraction Engine] (OCR / Text Parsing / Audio Whisper)
            │
            ▼
     [Contextual & Relationship Engine] (Entity Extraction / Intent Tagging)
            │
            ▼
     [Hybrid Vector + Keyword Indexer] (PostgreSQL 16 + pgvector)
            │
            ▼
     [Persona & Retrieval Intelligence Engine]
```
