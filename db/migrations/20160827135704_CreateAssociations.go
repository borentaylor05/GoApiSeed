package main

import (
	"database/sql"

	"github.com/borentaylor05/streamline/db"
)

// Up is executed when this migration is applied
func Up_20160827135704(txn *sql.Tx) {
	dbase, _ := db.Open()
	defer dbase.Close()

	// Alcohol Product
	dbase.Model(&db.AlcoholProduct{}).AddForeignKey("distributor_id", "distributors(id)", "RESTRICT", "RESTRICT")
	// Alcohol Orders
	dbase.Model(&db.AlcoholOrder{}).AddForeignKey("alcohol_product_id", "alcohol_products(id)", "RESTRICT", "RESTRICT")
	dbase.Model(&db.AlcoholOrder{}).AddForeignKey("retailer_id", "retailers(id)", "RESTRICT", "RESTRICT")
	// Distriutors
	dbase.Model(&db.Distributor{}).AddForeignKey("contact_info_id", "contact_infos(id)", "RESTRICT", "RESTRICT")
	dbase.Model(&db.Owner{}).AddForeignKey("contact_info_id", "contact_infos(id)", "RESTRICT", "RESTRICT")
	dbase.Model(&db.Retailer{}).AddForeignKey("contact_info_id", "contact_infos(id)", "RESTRICT", "RESTRICT")
}

// Down is executed when this migration is rolled back
func Down_20160827135704(txn *sql.Tx) {

}
