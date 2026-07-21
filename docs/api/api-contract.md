# Declutr API Contract

## Auth

POST /api/v1/auth/register

Request:
{
  "email": "",
  "encryptedMasterKey": ""
}

Response:
{
  "userId": ""
}

POST /api/v1/auth/login

Request:
{
  "email": ""
}

Response:
{
  "challenge": ""
}

## Vault

POST /api/v1/vault

GET /api/v1/vault/:id

## Files

POST /api/v1/files/upload

GET /api/v1/files/:id

DELETE /api/v1/files/:id

## Search

POST /api/v1/search
