package db

import (
	"GoApiSeed/util"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Open() (*gorm.DB, error) {
	config, err := util.ReadConfig()
	if err != nil {
		return nil, err
	}
	host := "host=" + config.DB.Host
	user := " user=" + config.DB.User
	dbname := " dbname=" + config.DB.Name
	sslmode := " sslmode=disable"
	password := " password=" + config.DB.Password
	db, err := gorm.Open("postgres", host+user+dbname+sslmode+password)
	if err != nil {
		return nil, err
	}
	return db, nil
}
