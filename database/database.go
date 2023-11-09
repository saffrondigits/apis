package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://root:abcd@localhost:1234/goclass?sslmode=disable")
	if err != nil {
		log.Println("Cannot connect to the db")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Cannot ping to the db")
		return nil, err
	}

	return db, nil
}
