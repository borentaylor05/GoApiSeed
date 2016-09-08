package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type AlcoholOrder struct {
	gorm.Model
	ItemAmount       int
	NeededBy         time.Time
	Retailer         Retailer
	RetailerID       uint64
	AlcoholProduct   AlcoholProduct
	AlcoholProductID uint64
}
