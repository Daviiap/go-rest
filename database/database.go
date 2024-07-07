package database

import (
	"go_rest_api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDatabaseConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})

	if err == nil {
		db.AutoMigrate(&models.User{})
	}

	return db, err
}
