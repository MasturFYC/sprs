package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func createConnection() *sql.DB {
	// load .env file
	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(25)
	db.Ping()

	//defer db.DB().Close()

	return db
}

func InitDatabase() func() *sql.DB {
	db := createConnection()

	fmt.Println("Connecting to database...")

	return func() *sql.DB {
		err := (*db).Ping()
		if err != nil {
			db = createConnection()
		}
		return db
	}
}
