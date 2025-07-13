package handlers

import (
	"net/http"

	"crawlzilla-backend/internal/db"
	"crawlzilla-backend/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateURLRequest struct {
	URL string `json:"url" binding:"required,url"`
}

func CreateURLHandler(c *gin.Context) {
	var req CreateURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	user := c.MustGet("user").(models.User)

	result, err := db.DB.Exec(`
		INSERT INTO urls (user_id, url, status)
		VALUES (?, ?, 'queued')
	`, user.ID, req.URL)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create URL entry"})
		return
	}

	id, _ := result.LastInsertId()

	c.JSON(http.StatusCreated, gin.H{
		"id":     id,
		"status": "queued",
		"url":    req.URL,
	})
}
