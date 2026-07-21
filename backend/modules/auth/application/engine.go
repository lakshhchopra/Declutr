package application

import (
	"crypto/rand"
	"encoding/hex"
)

type Engine struct{}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) GenerateServerSecret() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (e *Engine) GenerateServerPublicKey() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (e *Engine) VerifyClientProof(
	clientProof string,
	serverSecret string,
) bool {
	return clientProof != ""
}
