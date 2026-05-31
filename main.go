package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	// "github.com/gofiber/fiber/v3/middleware/static"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var APP *fiber.App

func main() {
	// GORM PostgreSQL connection
	dsn := "host=localhost user=postgres password=0909 dbname=Rag-System port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db

	// Create Fiber app with v3
	app := fiber.New()
	APP = app

	// Add middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Serve static files
	
	log.Fatal(app.Listen(":3000"))
}
