package cmd

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var logger zerolog.Logger

// @Summary Get all books
// @Description Get all books from the database
// @Produce json
// @Tags Books
// @Success 200 {array} Books
// @Router /getBooks [get]
func GetBooks(c *gin.Context) {
	db, err := DbConn(username, password, dbname, port)
	if err != nil {
		logger.Err(err).Msg("error connecting to db. ")
	}
	defer db.Close()
	rows, _ := db.Query("select * from books")
	books := parseRows(rows)
	c.JSON(http.StatusOK, books)
}

// @Summary Get books by author
// @Description Get books by author from the database
// @Produce json
// @Tags Books
// @Param author query string true "Author name"
// @Success 200 {array} Books
// @Router /getBooksByAuthor [get]
func GetBookbyAuthor(c *gin.Context) {
	db, err := DbConn(username, password, dbname, port)
	if err != nil {
		logger.Err(err).Msg("error connecting to db. ")
	}
	defer db.Close()
	author := c.Query("author")
	rows, _ := db.Query("select * from books where author=?", author)
	books := parseRows(rows)
	c.JSON(http.StatusOK, books)
}

// @Summary Get book by ID
// @Description Get a book by ID from the database
// @Produce json
// @Tags Books
// @Param id query string true "Book ID"
// @Success 200 {array} Books
// @Router /getBooksById [get]
func GetBookById(c *gin.Context) {
	db, err := DbConn(username, password, dbname, port)
	if err != nil {
		logger.Err(err).Msg("error connecting to db. ")
	}
	defer db.Close()
	id := c.Query("id")
	rows, _ := db.Query("select * from books where id=?", id)
	books := parseRows(rows)
	c.JSON(http.StatusOK, books)
}

func parseRows(rows *sql.Rows) *[]Books {
	var books []Books
	for rows.Next() {
		var book Books
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.YOP, &book.ISBN)
		if err != nil {
			logger.Err(err).Msg("Error parsing books in database")
		}
		books = append(books, book)
	}
	return &books
}
