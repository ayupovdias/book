package handlers

import (
	"book/config"
	"book/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categoryList []models.Category
	if err := db.DB.Find(&categoryList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch categories"})
		return
	}
	c.JSON(http.StatusOK, categoryList)
}

func AddCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save category"})
		return
	}

	c.JSON(http.StatusCreated, category)
}
