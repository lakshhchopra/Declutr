package domain

import "time"

type Session struct {
	ID string

	UserID string

	AccessToken string

	CreatedAt time.Time
	ExpiresAt time.Time
}
