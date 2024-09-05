package config

import (
	"log"

	"klinik-hewan/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection

func InitDB(dsn string) *gorm.DB {
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
    return DB
}
// MigrateSchema performs database schema migration
func MigrateSchema() {
	err := DB.AutoMigrate(&models.Hewan{}, &models.Produk{})
	if err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}
}
