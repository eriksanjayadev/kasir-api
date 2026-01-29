package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func initDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetConnMaxIdleTime(5)

	log.Println("Database connection successfully")
	return db, nil
}
