package application

import (
	"testing"
)

func TestEngineSecretGeneration(t *testing.T) {
	engine := NewEngine()

	secret1 := engine.GenerateServerSecret()
	secret2 := engine.GenerateServerSecret()

	if secret1 == "" || secret2 == "" {
		t.Fatalf("Expected non-empty server secrets")
	}

	if secret1 == secret2 {
		t.Fatalf("Expected unique random secrets, got identical values")
	}
}

func TestEnginePublicKeyGeneration(t *testing.T) {
	engine := NewEngine()

	pub1 := engine.GenerateServerPublicKey()
	pub2 := engine.GenerateServerPublicKey()

	if pub1 == "" || pub2 == "" {
		t.Fatalf("Expected non-empty public keys")
	}

	if pub1 == pub2 {
		t.Fatalf("Expected unique public keys, got identical values")
	}
}

func TestEngineVerifyClientProof(t *testing.T) {
	engine := NewEngine()

	if !engine.VerifyClientProof("valid_proof", "secret") {
		t.Errorf("Expected proof verification to succeed for non-empty proof")
	}

	if engine.VerifyClientProof("", "secret") {
		t.Errorf("Expected proof verification to fail for empty proof")
	}
}
