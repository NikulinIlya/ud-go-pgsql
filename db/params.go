package db

import (
	"log"

	pg "github.com/go-pg/pg"
)

// Params : special type for demo
type Params struct {
	Param1 string
	Param2 string
}

// PlaceHolderDemo :
func PlaceHolderDemo(db *pg.DB) error {
	var value string
	params := Params{
		Param1: "This is param 1",
		Param2: "This is param 2",
	}
	var query string = "SELECT ?param1"
	_, selectErr := db.Query(pg.Scan(&value), query, params)

	if selectErr != nil {
		log.Printf("Error while running the select query, reason: %v\n", selectErr)
		return selectErr
	}

	log.Printf("Scan successful, scanned value: %s\n", value)

	return nil
}
