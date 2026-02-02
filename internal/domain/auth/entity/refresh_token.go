package entity

import "time"

type RefreshToken struct {
	ID string
	UserID string
	ExpiresAt time.Time
	CreatedAt time.Time
}
