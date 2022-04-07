package database

import (
	"atm-api/models"
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Printf("We are connected to the database")

	db, _ := DB.DB()
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return migrateTables(DB)

}

func migrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Account{},
		&models.User{},
	)
}
