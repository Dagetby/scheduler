package routers

import (
	"database/sql"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *sql.DB //database

func init() {

	connStr := "user=postgres password=example dbname=postgres sslmode=disable"

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print(err)
	}
	db = conn
}

//returns a handle to the DB object
func GetDB() *sql.DB {
	return db
}
