package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func GetBooks(c *gin.Context) {
	db, err := DbConn(username, password, dbname, port)
	if err != nil {
		logger.Err(err).Msg("error connecting to db. ")
	}

	rows, _ := db.Query("select * from books")

	var books []Books
	for rows.Next() {
		var book Books
		err := rows.Scan(&book.ID, &book.Title, &book.Author, book.YOP, &book.ISBN)
		if err != nil {
			logger.Err(err).Msg("Error parsing books in database")
		}
		books = append(books, book)
	}
	defer db.Close()
	c.JSON(http.StatusOK, books)
}
