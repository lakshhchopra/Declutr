package application

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateAccessToken() string {
	b := make([]byte, 32)
	rand.Read(b)

	return hex.EncodeToString(b)
}
