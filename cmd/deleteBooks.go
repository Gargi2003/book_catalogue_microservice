package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const deleteQuery = "delete from books where id=?"

// DeleteBook deletes a book by its ID.
// @Summary Delete Book
// @Description Delete a book from the database by its ID
// @Tags Books
// @Accept json
// @Produce json
// @Param id query string true "ID of the book to delete"
// @Success 200 {string} string "Book deleted successfully"
// @Failure 404 {string} string "Book not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /deleteBook [delete]
func DeleteBook(c *gin.Context) {
	db, err := DbConn(username, password, dbname, port)
	if err != nil {
		logger.Err(err).Msg("error connecting to db. ")
	}
	defer db.Close()
	id := c.Query("id")
	result, err := db.Exec(deleteQuery, id)
	if err != nil {
		logger.Err(err).Msgf("Error occurred while executing query: %s", deleteQuery)
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Err(err).Msg("Error occurred while getting the number of affected rows")
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, "Book not found")
		return
	}

	c.JSON(http.StatusOK, "Book deleted successfully!!!")

}
