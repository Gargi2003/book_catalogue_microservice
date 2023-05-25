package cmd

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username   = "root"
	password   = "Dell0SS!"
	dbname     = "book_catalogue"
	topology   = "tcp"
	port       = "localhost:3306"
	driverName = "mysql"
)

type Books struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	YOP    int64  `json:"publication_year"`
	ISBN   string `json:"isbn"`
}

func ConstructURL(username string, password string, dbname string, port string) string {
	var sb strings.Builder
	sb.WriteString(username)
	sb.WriteString(":")
	sb.WriteString(password)
	sb.WriteString("@")
	sb.WriteString(topology)
	sb.WriteString("(")
	sb.WriteString(port)
	sb.WriteString(")")
	sb.WriteString("/")
	sb.WriteString(dbname)

	return sb.String()
}
func DbConn(username string, password string, dbname string, port string) (*sql.DB, error) {
	dataSourceName := ConstructURL(username, password, dbname, port)
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("couldn't establish db connection. Error: %v", err)
	}
	return db, nil
}
