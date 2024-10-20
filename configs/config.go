package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // ou _ "github.com/lib/pq" para PostgreSQL
)

func DBConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:rootpassword@tcp(127.0.0.1:3306)/metadata")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
