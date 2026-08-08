package todolist

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "lista de tarefas",
	})
}
