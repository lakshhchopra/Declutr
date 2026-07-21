package application

import (
	"testing"
	"time"

	"github.com/diablovocado/declutr/modules/auth/domain"
)

func TestChallengeStoreSaveAndGet(t *testing.T) {
	store := NewChallengeStore()

	ch := domain.SRPChallenge{
		ID:        "challenge_123",
		EmailHash: "hash_abc",
		Salt:      "salt_123",
		B:         "pub_B",
		b:         "sec_b",
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	store.Save(ch)

	fetched, exists := store.Get("challenge_123")
	if !exists {
		t.Fatalf("Expected challenge to exist in store")
	}

	if fetched.EmailHash != "hash_abc" || fetched.Salt != "salt_123" {
		t.Errorf("Fetched challenge data mismatch: %+v", fetched)
	}

	// Verify single-use deletion
	store.Delete("challenge_123")
	_, existsAfterDelete := store.Get("challenge_123")
	if existsAfterDelete {
		t.Errorf("Expected challenge to be deleted after single-use consumption")
	}
}

func TestChallengeStoreExpiration(t *testing.T) {
	store := NewChallengeStore()

	ch := domain.SRPChallenge{
		ID:        "expired_ch",
		EmailHash: "hash_abc",
		ExpiresAt: time.Now().Add(-1 * time.Minute), // Expired in past
	}

	store.Save(ch)

	_, exists := store.Get("expired_ch")
	if exists {
		t.Errorf("Expected expired challenge to be rejected by store")
	}
}
