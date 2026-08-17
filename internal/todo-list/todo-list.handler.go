package todolist

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *TodoService
}

func NewHandler(service *TodoService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetAll(c *gin.Context) {
	userID := uint(1)

	todos, err := h.service.GetTodoList(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todos)
}
