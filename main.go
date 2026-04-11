package main

import (
	"book/config"
	"book/handlers"
	"book/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/books", handlers.GetBooks)
		protected.POST("/books", handlers.AddBook)

		protected.GET("/books/favorites", handlers.GetFavoriteBooks)

		protected.GET("/books/:id", handlers.GetBook)
		protected.PUT("/books/:id", handlers.UpdateBook)
		protected.DELETE("/books/:id", handlers.DeleteBook)

		protected.PUT("/books/:id/favorites", handlers.AddToFavorites)
		protected.DELETE("/books/:id/favorites", handlers.RemoveFromFavorites)

		protected.GET("/authors", handlers.GetAuthors)
		protected.POST("/authors", handlers.AddAuthor)
		protected.GET("/categories", handlers.GetCategories)
		protected.POST("/categories", handlers.AddCategory)
	}

	r.Run(":8081")
}
