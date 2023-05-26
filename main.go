package main

import (
	handlers "book_catalogue_microservice/cmd"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := gin.Default()

	router.GET("/getBooks", handlers.GetBooks)
	router.GET("/getBooksByAuthor", handlers.GetBookbyAuthor)
	router.GET("/getBooksById", handlers.GetBookById)
	router.DELETE("/deleteBook", handlers.DeleteBook)
	router.POST("/addBook", handlers.AddBooks)
	router.Run(":8080")

}
