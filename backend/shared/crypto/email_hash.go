package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"golang.org/x/crypto/argon2"
)

func HashEmail(email string) string {
	normalized := strings.ToLower(strings.TrimSpace(email))

	salt := []byte("declutr-global-email-salt")

	hash := argon2.IDKey(
		[]byte(normalized),
		salt,
		1,
		64*1024,
		4,
		32,
	)

	sum := sha256.Sum256(hash)

	return hex.EncodeToString(sum[:])
}
