package auth

import "time"

type LoginResponse struct {
	User  *User
	Token string
}

type User struct {
	ID           uint
	Name         string
	UserJobTitle string
	UserLevel    string
	UserLevelID  uint
	Email        string
	Pass         string
	IsActive     bool
	CreatedAt    time.Time
	LastLogin    *time.Time
}

type UserResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	UserJob     string `json:"userJob"`
	UserLevel   string `json:"userLevel"`
	UserLevelID uint   `json:"userLevelID"`
	Email       string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
