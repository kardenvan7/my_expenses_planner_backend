package database

import (
	"database/sql"
	"fmt"
	"time"
)

func ConnectToDb() (*sql.DB, error) {
	fmt.Println("Connecting to db")
	db, err := sql.Open("mysql", "user:password@tcp(db:3306)/database")

	if err != nil {
		return nil, err
	}

	for db.Ping() != nil {
		fmt.Println("Attempting to connect to db")
		time.Sleep(5 * time.Second)
	}

	fmt.Println("Connected")

	return db, nil
}
