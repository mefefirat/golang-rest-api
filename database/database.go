package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	HOST     = "172.17.0.3"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "test"
	DATABASE = "hr"
)

var DB *sql.DB

func Connect() {
	var connectionString string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DATABASE)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	DB = db

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
