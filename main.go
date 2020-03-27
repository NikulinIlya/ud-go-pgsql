package main

import (
	"log"
	"time"

	db "github.com/NikulinIlya/ud-go-pgsql.git/db"
	"github.com/go-pg/pg"
)

func main() {
	log.Printf("Hello")
	pgDb := db.Connect()
	SaveProduct(pgDb)
	SaveTwoProducts(pgDb)

	db.PlaceHolderDemo(pgDb)

	DeleteItem(pgDb)
	UpdateItemPrice(pgDb)
	GetByID(pgDb)
}

// SaveProduct : save random product
func SaveProduct(dbRef *pg.DB) {
	newPI := &db.ProductItem{
		Name:  "Product 1",
		Desc:  "Product 1 desc",
		Image: "Image Path",
		Price: 4.5,
		Features: struct {
			Name string
			Desc string
			Imp  int
		}{
			Name: "F1",
			Desc: "F1 Desc",
			Imp:  3,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}

	newPI.SaveAndReturn(dbRef)
}

// SaveTwoProducts : save two random products
func SaveTwoProducts(dbRef *pg.DB) {
	newPIOne := &db.ProductItem{
		Name:  "Product 2",
		Desc:  "Product 2 desc",
		Image: "Image Path",
		Price: 4.5,
		Features: struct {
			Name string
			Desc string
			Imp  int
		}{
			Name: "F1",
			Desc: "F1 Desc",
			Imp:  3,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}

	newPITwo := &db.ProductItem{
		Name:  "Product 3",
		Desc:  "Product 3 desc",
		Image: "Image Path",
		Price: 4.5,
		Features: struct {
			Name string
			Desc string
			Imp  int
		}{
			Name: "F1",
			Desc: "F1 Desc",
			Imp:  3,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}

	totalItems := []*db.ProductItem{
		newPIOne,
		newPITwo,
	}

	newPIOne.SaveTwoItems(dbRef, totalItems)
}

// DeleteItem : delete the product item by its name
func DeleteItem(dbRef *pg.DB) {
	newPI := &db.ProductItem{
		Name: "Product 5",
	}

	newPI.DeleteItem(dbRef)
}

// UpdateItemPrice :
func UpdateItemPrice(dbRef *pg.DB) {
	newPI := &db.ProductItem{
		ID:    1,
		Price: 5.0,
	}

	newPI.UpdatePrice(dbRef)
}

func GetByID(dbRef *pg.DB) {
	newPI := &db.ProductItem{
		ID: 1,
	}
	newPI.GetByID(dbRef)
}
