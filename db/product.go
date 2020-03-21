package db

import (
	"time"
	"log"
	orm "github.com/go-pg/pg/orm"
	pg "github.com/go-pg/pg"
)

// ProductItem :
type ProductItem struct {
	RefPointer int `sql:"-"`
	tableName struct{} `sql:"product_items_collection"`
	ID int `sql:"id,pk"`
	Name string `sql:"name,unique"`
	Desc string `sql:"desc"`
	Image string `sql:"image"`
	Price float64 `sql"price,type:real"`
	Features struct {
		Name string
		Desc string
		Imp int
	} `sql:"features,type:jsonb"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive bool `sql:"is_active"`
}

// CreateProductItemsTable : create ProdItems database table
func CreateProductItemsTable(db *pg.DB) error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	crateErr := db.CreateTable(&ProductItem{}, options)

	if crateErr != nil {
		log.Printf("Error while creating table productItems, reason: %v\n", crateErr)
		return crateErr
	}

	log.Printf("Table ProductItems created successfully.\n")

	return nil
}