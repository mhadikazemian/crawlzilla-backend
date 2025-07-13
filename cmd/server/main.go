package main

import (
	"crawlzilla-backend/internal/db"
	"crawlzilla-backend/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect("root", "my-secret-pw", "127.0.0.1:3306", "crawlzilla")

	router := gin.Default()

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/me", func(c *gin.Context) {
		user, _ := c.Get("user")
		c.JSON(200, user)
	})

	log.Println("Server running on :8080")
	router.Run(":8080")
}
