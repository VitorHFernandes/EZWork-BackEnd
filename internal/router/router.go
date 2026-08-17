package router

import (
	"github.com/VitorHFernandes/EZWork-BackEnd/internal/auth"
	todolist "github.com/VitorHFernandes/EZWork-BackEnd/internal/todo-list"
	"github.com/gin-gonic/gin"
)

func New(
	todoHandler *todolist.Handler,
	authHandler *auth.Handler,
) *gin.Engine {
	r := gin.Default()

	authRoutes := r.Group("/auth")
	auth.RegisterRoutes(authRoutes, authHandler)

	todoRoutes := r.Group("/todos")
	todolist.RegisterRoutes(todoRoutes, todoHandler)

	return r
}
