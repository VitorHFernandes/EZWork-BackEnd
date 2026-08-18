package router

import (
	"github.com/VitorHFernandes/EZWork-BackEnd/internal/auth"
	"github.com/VitorHFernandes/EZWork-BackEnd/internal/auth/sessions"
	todolist "github.com/VitorHFernandes/EZWork-BackEnd/internal/todo-list"
	"github.com/gin-gonic/gin"
)

func New(
	todoHandler *todolist.Handler,
	authHandler *auth.Handler,
	sessionRepository sessions.SessionRepository,
) *gin.Engine {
	r := gin.Default()

	authMiddleware := auth.AuthMiddleware(sessionRepository)

	authRoutes := r.Group("/auth")
	auth.RegisterRoutes(authRoutes, authHandler)

	todoRoutes := r.Group("/todos")
	todolist.RegisterRoutes(
		todoRoutes,
		todoHandler,
		authMiddleware,
	)

	return r
}
