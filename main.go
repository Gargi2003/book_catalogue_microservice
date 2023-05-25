package main

import (
	"database/sql"
	"fmt"
	"log"
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

func main() {
	dbConn(username, password, dbname, port)
}

func dbConn(username string, password string, dbname string, port string) {
	dataSourceName := constructURL(username, password, dbname, port)
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("error connecting to db", err)
	}
	fmt.Println("db connection established!!!")
	defer db.Close()
}

func constructURL(username string, password string, dbname string, port string) string {
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
