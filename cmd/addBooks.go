package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Books []Books `json:"books"`
}

func AddBooks(c *gin.Context) {
	db, err := DbConn(username, password, dbname, port)
	if err != nil {
		logger.Err(err).Msg("error connecting to db. ")
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	smt, err := db.Prepare("insert into books (title,author,publication_year,isbn) values (?, ?, ?, ?)")
	if err != nil {
		logger.Err(err).Msg("error preparing query")
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	req := &Request{}
	err = c.BindJSON(req)
	if err != nil {
		logger.Err(err).Msg("error binding to json object")
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	for _, item := range req.Books {
		_, err := smt.Exec(item.Title, item.Author, item.YOP, item.ISBN)
		if err != nil {
			logger.Err(err).Msg("Error executing query")
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}
	defer smt.Close()
	c.JSON(http.StatusOK, "Books added successfully")
}
