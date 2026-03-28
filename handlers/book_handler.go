package handlers

import (
	"net/http"
	"strconv"

	"book/models"

	"github.com/gin-gonic/gin"
)

var books = make(map[int]models.Book)
var bookID = 1

func GetBooks(c *gin.Context) {
	var result []models.Book

	category := c.Query("category")

	for _, book := range books {
		if category != "" {
			if strconv.Itoa(book.CategoryID) != category {
				continue
			}
		}
		result = append(result, book)
	}

	c.JSON(http.StatusOK, result)
}

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if book.Title == "" || book.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	book.ID = bookID
	bookID++
	books[book.ID] = book

	c.JSON(http.StatusCreated, book)
}

func GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	book, exists := books[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, exists := books[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	var updated models.Book
	c.ShouldBindJSON(&updated)

	updated.ID = id
	books[id] = updated

	c.JSON(http.StatusOK, updated)
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	delete(books, id)

	c.Status(http.StatusNoContent)
}
