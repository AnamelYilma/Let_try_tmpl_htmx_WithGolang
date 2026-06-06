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
func DbLoad() (*gorm.DB, error) {
	var task model.TASK
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	dsn := fmt.Sprintf("host=%s user=postgres password=0909 dbname=serverside port=5432 sslmode=disable", host)
	var db *gorm.DB
	var err error
	maxAttempts := 6
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Printf("Attempt %d: Could not connect to database: %v\n", attempt, err)
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
	fmt.Println("✅ Connected to database successfully!")
	return db, nil
}