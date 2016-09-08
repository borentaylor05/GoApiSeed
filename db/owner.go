package db

import "github.com/jinzhu/gorm"

type Owner struct {
	gorm.Model
	FirstName     string
	LastName      string
	ContactInfo   ContactInfo
	ContactInfoID uint64
}
