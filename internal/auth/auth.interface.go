package auth

type AuthRepository interface {
	GetByEmail(email string) (*User, error)
}
