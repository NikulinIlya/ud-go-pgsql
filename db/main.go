package db

import (
	"log"
	"os"
	"time"

	pg "github.com/go-pg/pg"
)

// Connect : make connection to database
func Connect() *pg.DB {
	options := &pg.Options{
		User:         "user",
		Password:     "password",
		Addr:         "localhost:5432",
		Database:     "mydb",
		DialTimeout:  30 * time.Second,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		IdleTimeout:  30 * time.Minute,
		MaxConnAge:   1 * time.Minute,
		PoolSize:     20,
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
