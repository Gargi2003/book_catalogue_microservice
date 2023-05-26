package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type PatchRequest struct {
	Books []struct {
		Title  string `json:"title" `
		Author string `json:"author" `
		YOP    int    `json:"publication_year"`
		ISBN   string `json:"isbn"`
	} `json:"books" binding:"required"`
}

// UpdateBooks updates multiple books in the database.
// @Summary Update Books
// @Description Update multiple books in the database by their ID
// @Tags Books
// @Accept json
// @Produce json
// @Param id query string true "ID of the books to update"
// @Param books body Request true "Books to update"
// @Success 202 {string} string "Books updated successfully"
// @Failure 500 {string} string "Internal Server Error"
// @Router /updateBook [put]
func UpdateBooks(c *gin.Context) {
	db, err := DbConn(username, password, dbname, port)
	if err != nil {
		logger.Err(err).Msg("error connecting to db. ")
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	req := &Request{}
	if err := c.BindJSON(req); err != nil {
		logger.Err(err).Msg("error binding to json object")
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	id := c.Query("id")
	for _, book := range req.Books {
		if _, err := db.Exec("Update books set title=?, author=?, publication_year=?, isbn=? where id=?", book.Title, book.Author, book.YOP, book.ISBN, id); err != nil {
			logger.Err(err).Msg("Error executing query")
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}
	c.JSON(http.StatusAccepted, "Book updated successfully!!!")
}

// ReplaceBooks replaces the fields of a book in the database.
// @Summary Replace Book
// @Description Replace the fields of a book in the database by its ID
// @Tags Books
// @Accept json
// @Produce json
// @Param id query string true "ID of the book to replace"
// @Param books body PatchRequest true "Book fields to replace"
// @Success 202 {string} string "Book replaced successfully"
// @Failure 500 {string} string "Internal Server Error"
// @Router /replaceBook [patch]
func ReplaceBooks(c *gin.Context) {
	db, err := DbConn(username, password, dbname, port)
	if err != nil {
		logger.Err(err).Msg("error connecting to db. ")
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	req := &PatchRequest{}
	if err := c.BindJSON(req); err != nil {
		logger.Err(err).Msg("error binding to json object")
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	id := c.Query("id")
	query := "update books set "
	params := []interface{}{}

	for _, book := range req.Books {
		if book.Title != "" {
			query += "title=?, "
			params = append(params, book.Title)
		}
		if book.Author != "" {
			query += "author=?, "
			params = append(params, book.Author)
		}
		if book.YOP != 0 {
			query += "publication_year=?, "
			params = append(params, book.YOP)
		}
		if book.ISBN != "" {
			query += "isbn=?, "
			params = append(params, book.ISBN)
		}
	}
	query = strings.TrimSuffix(query, ", ")
	query += " where id=? "
	params = append(params, id)
	fmt.Println(query)
	fmt.Println(params...)
	if _, err := db.Exec(query, params...); err != nil {
		logger.Err(err).Msg("Error executing query")
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusAccepted, "Book replaced successfully!!!")

}
