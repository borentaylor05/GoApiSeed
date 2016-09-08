package main

import (
	"database/sql"

	"github.com/borentaylor05/streamline/db"
)

// Up is executed when this migration is applied
func Up_20160827135528(txn *sql.Tx) {
	database, _ := db.Open()
	defer database.Close()
	database.CreateTable(&db.Owner{}, &db.AlcoholProduct{}, &db.AlcoholOrder{}, &db.Retailer{}, &db.Distributor{}, &db.ContactInfo{})
}

// Down is executed when this migration is rolled back
func Down_20160827135528(txn *sql.Tx) {
	database, _ := db.Open()
	defer database.Close()
	database.DropTable(&db.Owner{}, &db.AlcoholProduct{}, &db.AlcoholOrder{}, &db.Retailer{}, &db.Distributor{}, &db.ContactInfo{})
}
