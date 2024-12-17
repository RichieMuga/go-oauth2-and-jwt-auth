// Package database contains the db connection
package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB varible for sqlirte
var DB *gorm.DB

// ConnectDatabase establishes a connection to the SQLite database
func ConnectDatabase() (*gorm.DB, error) {
	// Use a default database path if not specified in env
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./app.db"
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database. \n%v", err)
		return nil, err
	}

	// Set the global DB variable
	DB = db

	// Auto migrate your models here
	// For example:
	// err = db.AutoMigrate(&models.User{})
	// if err != nil {
	// 	log.Fatalf("Failed to auto migrate database. \n%v", err)
	// 	return nil, err
	// }

	return db, nil
}
