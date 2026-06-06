package repository

import (
	"database/sql"

	"github.com/diablovocado/declutr/internal/models"
)

type PostgresUserRepository struct {
	DB *sql.DB
}

func (r *PostgresUserRepository) CreateUser(user models.User) error {
	_, err := r.DB.Exec(`
		INSERT INTO users (
			id,
			email_hash,
			srp_verifier,
			srp_salt,
			encrypted_mvk_ciphertext,
			encrypted_mvk_nonce,
			encrypted_mvk_version,
			created_at,
			updated_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`,
		user.ID,
		user.EmailHash,
		user.SRPVerifier,
		user.SRPSalt,
		user.EncryptedMVKCiphertext,
		user.EncryptedMVKNonce,
		user.EncryptedMVKVersion,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}
