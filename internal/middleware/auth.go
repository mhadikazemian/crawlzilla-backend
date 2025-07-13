package middleware

import (
	"net/http"
	"strings"

	"crawlzilla-backend/internal/model"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			return
		}

		token := strings.TrimPrefix(auth, "Bearer ")
		user, err := model.GetUserByToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Set user in context
		c.Set("user", user)
		c.Next()
	}
}
