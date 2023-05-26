package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const deleteQuery = "delete from books where id=?"

func DeleteBook(c *gin.Context) {
	db, err := DbConn(username, password, dbname, port)
	if err != nil {
		logger.Err(err).Msg("error connecting to db. ")
	}
	defer db.Close()
	id := c.Query("id")
	rows, err1 := db.Query(deleteQuery, id)
	if err1 != nil {
		logger.Err(err).Msgf("Error occurred while executing query %s", deleteQuery)
	}
	if !rows.Next() {
		c.JSON(http.StatusNotFound, "No matching elements found witht the given id")
		return
	}
	c.JSON(http.StatusOK, "Books deleted successfully!!!")
}
