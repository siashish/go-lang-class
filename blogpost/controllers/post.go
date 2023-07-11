package controllers

import (
	"net/http"
	"siashish/blogpost/models"

	"github.com/gin-gonic/gin"
)

type CreatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := models.Blogpost{Title: input.Title, Content: input.Content}
	models.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}