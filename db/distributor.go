package db

import "github.com/jinzhu/gorm"

type Distributor struct {
	gorm.Model
	Name          string
	PocName       string
	ContactInfo   ContactInfo
	ContactInfoID uint64
	GUID          string
	Password      string
	Username      string
}
