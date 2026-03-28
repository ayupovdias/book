package handlers

import (
	"net/http"

	"book/models"

	"github.com/gin-gonic/gin"
)

var categories = make(map[int]models.Category)
var categoryID = 1

func GetCategories(c *gin.Context) {
	var result []models.Category
	for _, c := range categories {
		result = append(result, c)
	}
	c.JSON(http.StatusOK, result)
}

func CreateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if category.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name required"})
		return
	}

	category.ID = categoryID
	categoryID++
	categories[category.ID] = category

	c.JSON(http.StatusCreated, category)
}
