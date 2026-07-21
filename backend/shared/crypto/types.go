package crypto

type EncryptedMVK struct {
	Ciphertext string `json:"ciphertext"`
	Nonce      string `json:"nonce"`
	Version    int    `json:"version"`
}

type EncryptedVaultKey struct {
	Ciphertext string `json:"ciphertext"`
	Nonce      string `json:"nonce"`
}

type EncryptedFileKey struct {
	Ciphertext string `json:"ciphertext"`
	Nonce      string `json:"nonce"`
}
