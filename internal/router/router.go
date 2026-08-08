package router

import (
	todolist "github.com/VitorHFernandes/EZWork-BackEnd/internal/todo-list"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	todolist.RegisterRoutes(r)

	return r
}
