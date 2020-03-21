package main

import (
	"log"

	db "github.com/NikulinIlya/ud-go-pgsql.git/db"
)

func main() {
	log.Printf("Hello")
	db.Connect()
}
