package main

import (
	"library/handlers"
	"library/repository"
	"library/service"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Library Management API
// @version 1.0
// @description This is a library management server.
// @host localhost:8080
// @BasePath /api
func main() {
	// Initialize repository and service
	repo := repository.NewBookRepository()
	bookService := service.NewBookService(repo)
	bookHandler := handlers.NewBookHandler(bookService)

	// Initialize Gin router
	router := gin.Default()

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Book routes
	books := router.Group("/api/books")
	{
		books.POST("", bookHandler.CreateBook)
		books.GET("", bookHandler.ListBooks)
		books.GET("/:id", bookHandler.GetBook)
		books.PUT("/:id", bookHandler.UpdateBook)
		books.DELETE("/:id", bookHandler.DeleteBook)
	}

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
