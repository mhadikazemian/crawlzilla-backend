package middleware

import (
	"net/http"
	"strings"

	"crawlzilla-backend/internal/models"
	"crawlzilla-backendinternal/db"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			return
		}

		token := strings.TrimPrefix(auth, "Bearer ")

		var user models.User
		err := db.DB.Get(&user, "SELECT * FROM users WHERE api_token = ?", token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid API token"})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
