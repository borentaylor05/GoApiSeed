package db

import "github.com/jinzhu/gorm"

type AlcoholProduct struct {
	gorm.Model
	Name          string
	Description   string
	Type          string
	Price         uint64
	Distributor   Distributor
	DistributorID uint64
}
