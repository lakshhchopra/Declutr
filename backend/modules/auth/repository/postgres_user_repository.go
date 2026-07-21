package repository

import (
	"database/sql"

	"github.com/diablovocado/declutr/modules/auth/domain"
)

type PostgresUserRepository struct {
	DB *sql.DB
}

func (r *PostgresUserRepository) CreateUser(user domain.User) error {
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

func (r *PostgresUserRepository) CreateSession(session domain.Session) error {
	_, err := r.DB.Exec(`
		INSERT INTO sessions (
			id,
			user_id,
			access_token,
			created_at,
			expires_at
		)
		VALUES ($1,$2,$3,$4,$5)
	`,
		session.ID,
		session.UserID,
		session.AccessToken,
		session.CreatedAt,
		session.ExpiresAt,
	)

	return err
}

func (r *PostgresUserRepository) GetSessionByToken(token string) (*domain.Session, error) {
	session := &domain.Session{}

	err := r.DB.QueryRow(`
		SELECT
			id,
			user_id,
			access_token,
			created_at,
			expires_at
		FROM sessions
		WHERE access_token = $1
	`,
		token,
	).Scan(
		&session.ID,
		&session.UserID,
		&session.AccessToken,
		&session.CreatedAt,
		&session.ExpiresAt,
	)

	if err != nil {
		return nil, err
	}

	return session, nil
}

func (r *PostgresUserRepository) DeleteSession(token string) error {
	_, err := r.DB.Exec(`
		DELETE FROM sessions
		WHERE access_token = $1
	`, token)

	return err
}

func (r *PostgresUserRepository) GetUserByEmailHash(emailHash string) (*domain.User, error) {
	user := &domain.User{}

	err := r.DB.QueryRow(`
		SELECT
			id,
			email_hash,
			srp_verifier,
			srp_salt,
			encrypted_mvk_ciphertext,
			encrypted_mvk_nonce,
			encrypted_mvk_version,
			created_at,
			updated_at
		FROM users
		WHERE email_hash = $1
	`,
		emailHash,
	).Scan(
		&user.ID,
		&user.EmailHash,
		&user.SRPVerifier,
		&user.SRPSalt,
		&user.EncryptedMVKCiphertext,
		&user.EncryptedMVKNonce,
		&user.EncryptedMVKVersion,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
