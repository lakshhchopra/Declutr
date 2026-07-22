# Declutr Official SDK Guide

Declutr provides official SDKs for **TypeScript**, **Go**, and **Python**.

## TypeScript SDK (`@declutr/sdk`)

```typescript
import { DeclutrClient } from '@declutr/sdk';

const client = new DeclutrClient({
  apiKey: "declutr_live_...",
  baseUrl: "http://localhost:8080"
});

const results = await client.search("legal contract");
console.log(results);
```

## Go SDK (`github.com/diablovocado/declutr/sdks/go`)

```go
client := declutr.NewClient(declutr.Config{
    APIKey: "declutr_live_...",
})

res, err := client.Search(ctx, "project plans", nil)
```

## Python SDK (`declutr-sdk`)

```python
from declutr import DeclutrClient

client = DeclutrClient(api_key="declutr_live_...")
response = client.chat(conversation_id="c-1", message="Summarize files")
```
