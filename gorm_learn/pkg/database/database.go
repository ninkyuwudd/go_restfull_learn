package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitializeDb() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/goNewTable")

	if err != nil {
		message := fmt.Sprintf("error open the database %d", err)
		panic(message)
	}

	return db
}
