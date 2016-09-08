package db

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

type Retailer struct {
	gorm.Model
	Name          string
	ContactInfo   *ContactInfo `gorm:"ForeignKey:ContactInfoID"`
	ContactInfoID uint64
	GUID          string
	Password      string
	Username      string
}

func (r *Retailer) Save() error {
	r.GUID = r.MakeGuid()

	if err := r.Valid(); err != nil {
		return err
	}

	hashedPassword, err := HashPassword(r.Password)
	if err != nil {
		return err
	}
	r.Password = hashedPassword

	db, err := OpenDB()
	defer db.Close()
	if err != nil {
		return err
	}

	db.Create(&r)
	return nil
}

func (r *Retailer) Valid() error {
	if r.Name == "" {
		return fmt.Errorf("Retailer Name is required.")
	}
	if r.ContactInfo == nil {
		return fmt.Errorf("Retailer ContactInfo is required.")
	}
	if r.GUID == "" {
		return fmt.Errorf("Retailer GUID is required.")
	}
	if r.Password == "" {
		return fmt.Errorf("Retailer Password is required.")
	}
	if r.Username == "" {
		return fmt.Errorf("Retailer Username is required.")
	}

	db, err := OpenDB()
	defer db.Close()
	if err != nil {
		return err
	}

	if !db.Where("username = ?", r.Username).First(&Retailer{}).RecordNotFound() {
		return fmt.Errorf("Username `" + r.Username + "` is taken. Choose another.")
	}
	return nil
}

func (r *Retailer) MakeGuid() string {
	db, _ := OpenDB()
	guid := strings.ToLower(r.Name)
	if db.Where("guid = ?", guid).First(&Retailer{}).RecordNotFound() {
		return guid
	}

	count := 1
	for !db.Where("guid = ?", guid).First(&Retailer{}).RecordNotFound() {
		guid = strings.ToLower(r.Name) + "-" + strconv.Itoa(count)
		count++
	}

	return guid
}

// For User interface
func (r *Retailer) Login() (*Retailer, error) {
	db, err := OpenDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}

	var dbRetailer Retailer
	err = db.Where("username = ?", r.Username).First(&dbRetailer).Error
	if err != nil {
		return nil, err
	}

	// Check passwords match
	if !DoesPasswordMatch(dbRetailer.Password, r.Password) {
		return nil, fmt.Errorf("Username / password combo is incorrect.")
	}
	// If they match, do not return an error
	return &dbRetailer, nil
}

func (r *Retailer) UpdateRefreshToken(randomString string) {
	db, _ := OpenDB()
	defer db.Close()

	db.Model(&r).Update("refresh_token", randomString)
}
