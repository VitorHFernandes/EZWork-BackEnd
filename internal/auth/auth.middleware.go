package auth

import (
	"net/http"

	sessions "github.com/VitorHFernandes/EZWork-BackEnd/internal/auth/sessions"
	"github.com/VitorHFernandes/EZWork-BackEnd/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(SessionRepository sessions.SessionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("session")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
		}

		tokenHash := utils.HashToken(cookie)

		userID, err := SessionRepository.GetUserIDByTokenHash(tokenHash)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
		}

		c.Set("userID", userID)
		c.Next()

	}
}
