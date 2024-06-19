package db

import (
	"log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"app/models"

	"gorm.io/gorm"
)

// github.com/mattn/go-sqlite3
func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Services{})
	db.AutoMigrate(&models.LogChecked{})
	return db
}
