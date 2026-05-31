package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbLoad() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=0909 dbname=Rag-System port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	DB = db
	return db, nil
}
