package database

import (
	"fmt"
	"os"
	"time"

	"gootmplhtmx/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// DbLoad initializes the database connection. It reads DB_HOST from the
// environment (falls back to localhost), adds sslmode=disable for local
// development, and retries a few times while Docker/Postgres starts.
func DbLoad() (*gorm.DB, error) {
	var task model.TASK

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	dsn := fmt.Sprintf("host=%s user=postgres password=0909 dbname=serverside port=5432 sslmode=disable", host)

	var db *gorm.DB
	var err error

	// Retry loop: helpful when Postgres is starting in Docker Compose
	maxAttempts := 6
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		if attempt < maxAttempts {
			time.Sleep(2 * time.Second)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxAttempts, err)
	}

	DB = db

	if err := db.AutoMigrate(&task); err != nil {
		return nil, fmt.Errorf("auto-migrate failed: %w", err)
	}

	return db, nil
}
