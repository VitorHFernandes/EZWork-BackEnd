package todolist

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, handler *Handler, authMiddleware gin.HandlerFunc) {
	r.Use(authMiddleware)
	r.GET("/", handler.GetAll)
}
