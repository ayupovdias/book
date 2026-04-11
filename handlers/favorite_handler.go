package handlers

import (
	"book/config"
	"book/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFavoriteBooks(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var favorites []models.FavoriteBook
	db.DB.Where("user_id = ?", userID).
		Preload("Book").
		Limit(limit).Offset(offset).
		Find(&favorites)

	c.JSON(http.StatusOK, favorites)
}

func AddToFavorites(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	bookID, _ := strconv.Atoi(c.Param("id"))

	fav := models.FavoriteBook{
		UserID: userID,
		BookID: uint(bookID),
	}

	if err := db.DB.FirstOrCreate(&fav).Preload("Book").First(&fav).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add to favorites"})
		return
	}

	c.JSON(http.StatusOK, fav)
}

func RemoveFromFavorites(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	result := db.DB.Where("user_id = ? AND book_id = ?", userID.(uint), bookID).Delete(&models.FavoriteBook{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found in favorites"})
		return
	}
	c.Status(http.StatusNoContent)
}
