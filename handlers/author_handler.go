package handlers

import (
	"book/config"
	"book/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAuthors(c *gin.Context) {
	var authorList []models.Author
	db.DB.Find(&authorList)
	c.JSON(http.StatusOK, authorList)
}

func AddAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save author"})
		return
	}
	c.JSON(http.StatusCreated, author)
}
