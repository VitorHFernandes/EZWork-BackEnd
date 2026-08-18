package sessions

import "time"

type SessionRepository interface {
	Create(userID uint, tokenHash string, expiresAt time.Time) error
	GetUserIDByTokenHash(tokenHash string) (uint, error)
	DeleteSessionToken(tokenHash string) error
}
