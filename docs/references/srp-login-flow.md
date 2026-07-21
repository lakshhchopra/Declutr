# Declutr SRP Login Flow

## Login Start

Client -> Server

{
  email
}

Server:

1. Hash email
2. Find user
3. Load:
   - SRP Salt
   - SRP Verifier
4. Generate server ephemeral secret b
5. Generate server public key B
6. Store challenge state

Response:

{
  salt,
  serverPublicKey
}

## Login Finish

Client -> Server

{
  clientPublicKey,
  clientProof
}

Server:

1. Load challenge state
2. Verify client proof
3. Generate server proof
4. Create authenticated session

Response:

{
  serverProof
}

