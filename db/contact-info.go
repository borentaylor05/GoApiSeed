package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type ContactInfo struct {
	gorm.Model
	PhoneNumber string
	Address     string
	Email       string
	City        string
	State       string
	Zip         string
	Type        string
	ContactID   uint64 // TODO figure out how to implement
}

func (c *ContactInfo) Save() error {
	if err := c.Valid(); err != nil {
		return err
	}

	db, err := OpenDB()
	defer db.Close()
	if err != nil {
		return err
	}

	db.Create(&c)
	return nil
}

func (c *ContactInfo) Valid() error {
	if c.Email == "" {
		return fmt.Errorf("Email is required.")
	}
	if c.PhoneNumber == "" {
		return fmt.Errorf("Phone number is required.")
	}
	if c.Address == "" {
		return fmt.Errorf("Address is required.")
	}
	if c.City == "" {
		return fmt.Errorf("City is required.")
	}
	if c.State == "" {
		return fmt.Errorf("State is required.")
	}
	if c.Zip == "" {
		return fmt.Errorf("Zip is required.")
	}
	if c.Type == "" {
		return fmt.Errorf("Type is required.")
	}
	return nil
}
