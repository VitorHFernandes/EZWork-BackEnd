package todolist

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	handler := &Handler{}

	r.GET("todo", handler.GetAll)
}
