package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Connecting to db")
	db, err := sql.Open("mysql", "user:password@tcp(db:3306)/database")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	for db.Ping() != nil {
		fmt.Println("Attempting to connect to db")
		time.Sleep(5 * time.Second)
	}

	fmt.Println("Connected")

	fmt.Println("Dropping tables")
	_, err = db.Exec(`DROP TABLE IF EXISTS transactions, categories;`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Creating tables")

	_, err = db.Exec(`
		CREATE TABLE categories (
			ID int AUTO_INCREMENT PRIMARY KEY,
			UUID bigint NOT NULL,
			Name varchar(255) NOT NULL,
			Color varchar(8) NOT NULL
		);
	`)

	if err != nil {
		log.Fatal("categories: ", err)
	}

	_, err = db.Exec(`
		CREATE TABLE transactions (
			ID int AUTO_INCREMENT PRIMARY KEY,
			UUID bigint NOT NULL,
			Name varchar(255) NOT NULL,
			Amount float NOT NULL,
			Date bigint NOT NULL,
			Type varchar(255) NOT NULL,
			CategoryUuid bigint REFERENCES categories(UUID)
		);
	`)

	if err != nil {
		log.Fatal("transactions: ", err)
	}

	fmt.Println("Filling tables")

	_, err = db.Exec(`
		INSERT INTO categories (UUID, Name, Color) VALUES (1, 'Test category', 'FFFFFFFF');
	`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		INSERT INTO transactions (UUID, Name, Amount, Date, Type, CategoryUuid) VALUES (1, 'Test transaction', 255, 1653544523, 'expense', 1);
	`)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Getting category")

	result, err := db.Query("SELECT * FROM categories;")

	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	var category Category

	if result.Next() {
		err = result.Scan(
			&category.ID,
			&category.UUID,
			&category.Name,
			&category.Color,
		)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v", category)
	}

	fmt.Println("Finished")
}

type Category struct {
	ID    int
	UUID  int
	Name  string
	Color string
}

type Transaction struct {
	ID           int
	UUID         int
	Name         string
	Amount       float32
	Date         int
	Type         string
	CategoryUuid sql.NullInt32
}
