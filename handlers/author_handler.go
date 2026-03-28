package handlers

import (
	"book/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var authors = make(map[int]models.Author)
var authorID = 1

func GetAuthors(c *gin.Context) {
	var result []models.Author
	for _, a := range authors {
		result = append(result, a)
	}
	c.JSON(http.StatusOK, result)
}

func CreateAuthor(c *gin.Context) {
	var author models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if author.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name required"})
		return
	}

	author.ID = authorID
	authorID++
	authors[author.ID] = author

	c.JSON(http.StatusCreated, author)
}
