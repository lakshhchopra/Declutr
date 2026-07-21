# Declutr API Specification

> **Source of Truth:** [declutr_architecture_document.html](file:///f:/Github/Declutr/docs/architecture/declutr_architecture_document.html)  
> **Section:** 19. API Architecture

---

## Overview

Declutr services expose modular endpoints using standard REST over HTTPS with JSON payloads. Authentication uses bearer tokens issued upon successful SRP-6a verification.

---

## Endpoints Manifest

### 1. Authentication & Session

#### `POST /v1/auth/srp/initiate`
* **Auth Required:** None (Anonymous)
* **Description:** Initiates the zero-knowledge SRP-6a login challenge.
* **Request Payload:**
```json
{
  "email": "user@example.com"
}
```
* **Response Payload:**
```json
{
  "challengeId": "ch_9f8d1c",
  "salt": "hex_salt_string",
  "serverPublicKey": "hex_server_public_key"
}
```

#### `POST /v1/auth/srp/verify`
* **Auth Required:** None (Anonymous)
* **Description:** Submits client proof `M1` for validation and retrieves session credentials.
* **Request Payload:**
```json
{
  "challengeId": "ch_9f8d1c",
  "email": "user@example.com",
  "clientPublicKey": "hex_client_public_key",
  "clientProof": "hex_client_proof_M1"
}
```
* **Response Payload:**
```json
{
  "serverProof": "hex_server_proof_M2",
  "accessToken": "ey...jwt_access_token"
}
```

---

### 2. Ingestion & File Uploads

#### `POST /v1/files/upload/initiate`
* **Auth Required:** Bearer Token (`JWT`)
* **Description:** Starts a direct-to-S3 chunked upload session and returns pre-signed URLs.
* **Request Payload:**
```json
{
  "vaultId": "v_123",
  "fileSize": 10485760,
  "chunkCount": 3
}
```
* **Response Payload:**
```json
{
  "uploadId": "up_789",
  "urls": [
    "https://storage.cloudflare.com/r2/chunk_1?signature=...",
    "https://storage.cloudflare.com/r2/chunk_2?signature=...",
    "https://storage.cloudflare.com/r2/chunk_3?signature=..."
  ]
}
```

#### `POST /v1/files/upload/commit`
* **Auth Required:** Bearer Token (`JWT`)
* **Description:** Finalizes S3 transfers, commits record state, and enqueues background AI ingestion.
* **Request Payload:**
```json
{
  "uploadId": "up_789",
  "encryptedFileKey": "base64_efk",
  "encryptedMetadata": "base64_meta"
}
```
* **Response Payload:**
```json
{
  "itemId": "item_abc123",
  "status": "QUEUED"
}
```

---

### 3. Retrieval & Hybrid Search

#### `POST /v1/search/query`
* **Auth Required:** Bearer Token (`JWT`)
* **Description:** Resolves hybrid natural-language keyword and semantic vector queries.
* **Request Payload:**
```json
{
  "queryText": "hotel booking receipt Mumbai trip",
  "limit": 10,
  "filters": {
    "vaultId": "v_123"
  }
}
```
* **Response Payload:**
```json
{
  "results": [
    {
      "itemId": "item_abc123",
      "score": 0.92,
      "matchedContext": "Mumbai Business Trip",
      "encryptedMetadata": "base64_payload"
    }
  ]
}
```

---

### 4. Relationships & Feedback

#### `POST /v1/relationships/link`
* **Auth Required:** Bearer Token (`JWT`)
* **Description:** Connects two digital items in the relationship graph.
* **Request Payload:**
```json
{
  "sourceItemId": "item_1",
  "targetItemId": "item_2",
  "relationshipType": "PART_OF"
}
```

#### `POST /v1/feedback/verify`
* **Auth Required:** Bearer Token (`JWT`)
* **Description:** Submits user confirmation, correction, or rejection of AI-generated tags/intents.
* **Request Payload:**
```json
{
  "aiMetaId": "aim_456",
  "status": "accepted"
}
```
