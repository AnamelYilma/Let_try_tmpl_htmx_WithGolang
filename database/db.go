package database

import (
	"fmt"
	// "gootmplhtmx/Route"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gootmplhtmx/model"
)

var DB *gorm.DB
var task model.TASK

func DbLoad() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=0909 dbname=serverside port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	
	DB = db
	db.AutoMigrate(&task)

	return db, nil
}
