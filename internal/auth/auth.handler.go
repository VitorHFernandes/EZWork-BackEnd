package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *AuthService
}

func NewHandler(service *AuthService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Login(c *gin.Context) {
	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request",
		})
		return
	}

	res, err := h.service.Login(request.Email, request.Password)
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Credentials",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	userResponse := UserResponse{
		ID:          res.User.ID,
		Name:        res.User.Name,
		UserJob:     res.User.UserJobTitle,
		UserLevel:   res.User.UserLevel,
		UserLevelID: res.User.UserLevelID,
		Email:       res.User.Email,
	}

	c.SetCookie(
		"session",
		res.Token,
		86400,
		"/",
		"",
		false, //TODO => Colocar em true para https.
		true,
	)

	c.JSON(http.StatusOK, userResponse)
}
