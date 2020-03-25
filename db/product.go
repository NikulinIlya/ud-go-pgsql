package db

import (
	"log"
	"time"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

// ProductItem :
type ProductItem struct {
	RefPointer int      `sql:"-"`
	tableName  struct{} `sql:"product_items_collection"`
	ID         int      `sql:"id,pk"`
	Name       string   `sql:"name,unique"`
	Desc       string   `sql:"desc"`
	Image      string   `sql:"image"`
	Price      float64  `sql"price,type:real"`
	Features   struct {
		Name string
		Desc string
		Imp  int
	} `sql:"features,type:jsonb"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive  bool      `sql:"is_active"`
}

// Save : save ProdItem to database
func (pi *ProductItem) Save(db *pg.DB) error {
	insertErr := db.Insert(pi)

	if insertErr != nil {
		log.Printf("Error while inserting new item into DB, reason: %v\n", insertErr)
		return insertErr
	}
	log.Printf("ProductItem %s inserted successfully.\n", pi.Name)
	return nil
}

// SaveAndReturn :
func (pi *ProductItem) SaveAndReturn(db *pg.DB) (*ProductItem, error) {
	InsertResult, insertErr := db.Model(pi).Returning("*").Insert()
	if insertErr != nil {
		log.Printf("Error while inserting new item, reason: %v\n", insertErr)
		return nil, insertErr
	}

	log.Printf("ProductItem inserted successfully")
	log.Printf("Received new value result is: $v\n", InsertResult.RowsAffected)

	return pi, nil
}

// SaveTwoItems :
func (pi *ProductItem) SaveTwoItems(db *pg.DB, items []*ProductItem) error {
	_, insertErr := db.Model(items[0], items[1]).Insert()

	if insertErr != nil {
		log.Printf("Error while inserting bulk items, reason: %v\n", insertErr)
		return insertErr
	}
	log.Printf("Bulk insert successful\n")
	return nil
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

// DeleteItem : delete the product item selected by its name
func (pi *ProductItem) DeleteItem(db *pg.DB) error {
	_, deleteErr := db.Model(pi).Where("name = ?name").Delete()
	if deleteErr != nil {
		log.Printf("Error while deleting item, reason: $v\n", deleteErr)
	}
	log.Printf("Delete successful for %s, item\n", pi.Name)
	return nil
}
