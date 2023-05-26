package main

import (
	handlers "book_catalogue_microservice/cmd"
	_ "book_catalogue_microservice/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title   Book Catalogue APIs
// @version 1.0
// @description A book catalogue service using gin-framework
// @host localhost:8080
func main() {

	router := gin.Default()
	// Serve the Swagger UI
	router.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/getBooks", handlers.GetBooks)
	router.GET("/getBooksByAuthor", handlers.GetBookbyAuthor)
	router.GET("/getBooksById", handlers.GetBookById)
	router.DELETE("/deleteBook", handlers.DeleteBook)
	router.POST("/addBook", handlers.AddBooks)
	router.PUT("/updateBook", handlers.UpdateBooks)
	router.PATCH("/replaceBook", handlers.ReplaceBooks)
	router.Run(":8080")

}
