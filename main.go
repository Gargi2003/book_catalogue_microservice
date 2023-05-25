package main

import (
	"book_catalogue_microservice/cmd"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := gin.Default()

	router.GET("/getBooks", cmd.GetBooks)
	router.Run(":8080")

}
