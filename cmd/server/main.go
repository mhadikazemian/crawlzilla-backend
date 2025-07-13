package main

import (
	"crawlzilla-backend/internal/db"
	"crawlzilla-backend/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mhadikazemian/crawlzilla-backend/internal/middleware"
)

func main() {
	db.Connect("root", "yourpass", "127.0.0.1:3306", "page_analyzer")

	r := gin.Default()

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/urls", handlers.CreateURLHandler)
	}

	log.Println("🚀 Listening on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
