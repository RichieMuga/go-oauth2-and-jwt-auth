package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
    if err != nil {
        panic(fmt.Sprintf("Failed to connect the database: %v", err))
    }
}
