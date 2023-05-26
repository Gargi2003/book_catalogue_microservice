package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

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

func ReplaceBooks(c *gin.Context) {
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
