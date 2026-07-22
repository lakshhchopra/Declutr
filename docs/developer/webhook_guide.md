# Declutr Webhook Platform Guide

Declutr's Webhook Engine delivers real-time event notifications to third-party endpoints.

## HMAC-SHA256 Signature Verification

Every HTTP POST payload from Declutr contains an `X-Declutr-Signature` header computed as:

`HMAC-SHA256(webhook_secret, raw_payload_bytes)`

```python
import hmac, hashlib

def verify_signature(secret, payload_bytes, signature_header):
    expected = "sha256=" + hmac.new(secret.encode(), payload_bytes, hashlib.sha256).hexdigest()
    return hmac.compare_digest(expected, signature_header)
```

## Dead Letter Queue (DLQ)

If a webhook endpoint returns a non-2xx HTTP code or times out, Declutr retries delivery up to 3 times with exponential backoff. If all attempts fail, the event is moved to the **Dead Letter Queue (DLQ)** for inspection and re-triggering.
