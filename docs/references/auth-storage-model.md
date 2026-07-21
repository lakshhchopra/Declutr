# Declutr Auth Storage Model

## User Record

Stored on server:

- user_id
- email
- srp_verifier
- srp_salt
- encrypted_mvk_ciphertext
- encrypted_mvk_nonce
- encrypted_mvk_version
- created_at
- updated_at

Server never stores:

- password
- MEK
- MVK plaintext
- vault keys plaintext
- file keys plaintext

## Authentication

Protocol:
- SRP-6a

Server stores:
- verifier
- salt

Server never receives:
- password
- MEK

## Recovery

TBD (architecture decision required)
