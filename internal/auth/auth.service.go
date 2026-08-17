package auth

import (
	"errors"
	"time"

	sessions "github.com/VitorHFernandes/EZWork-BackEnd/internal/auth/sessions"
	"github.com/VitorHFernandes/EZWork-BackEnd/internal/utils"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService struct {
	authRepository    AuthRepository
	sessionRepository sessions.SessionRepository
}

func NewAuthService(authRepository AuthRepository, sessionRepository sessions.SessionRepository) *AuthService {
	return &AuthService{
		authRepository:    authRepository,
		sessionRepository: sessionRepository,
	}
}

func (s *AuthService) Login(email string, pass string) (*LoginResponse, error) {
	user, err := s.authRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(pass, user.Pass) {
		return nil, ErrInvalidCredentials
	}

	token, err := utils.GenerateToken()
	if err != nil {
		return nil, err
	}

	tokenHash := utils.HashToken(token)
	expiresAt := time.Now().Add(24 * time.Hour)

	if err := s.sessionRepository.Create(user.ID, tokenHash, expiresAt); err != nil {
		return nil, err
	}

	return &LoginResponse{
		User:  user,
		Token: token,
	}, nil
}
