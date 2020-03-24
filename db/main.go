package db

import (
	"log"
	"os"

	pg "github.com/go-pg/pg"
)

// Connect : make connection to database
func Connect() *pg.DB {
	options := &pg.Options{
		User:     "user",
		Password: "password",
		Addr:     "localhost:5432",
		Database: "mydb",
	}

	var db *pg.DB = pg.Connect(options)

	if db == nil {
		log.Printf("Failed to connect to database.\n")
		os.Exit(100)
	}

	log.Printf("Connection to database is successful.\n")

	CreateProductItemsTable(db)

	return db
}
