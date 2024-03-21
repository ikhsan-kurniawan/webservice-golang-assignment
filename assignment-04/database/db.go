package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "wawan123"
	dbPort   = "5432"
	dbName   = "mygram"
	db       *gorm.DB
	err      error
)

func StartDB() (*gorm.DB, error) {
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, dbPort, user, password, dbName)
	
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	
	return db, err
}