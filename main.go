package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/static"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

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

	// Add middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Serve static files
	app.Use("/css", static.New(static.Config{
		Root: "./css",
	}))
	app.Use("/js", static.New(static.Config{
		Root: "./node_modules/htmx.org/dist",
	}))

	// Routes
	app.Get("/", func(c fiber.Ctx) error {
		return Page().Render(c.Context(), c.Response().BodyWriter())
	})

	app.Post("/clicked", func(c fiber.Ctx) error {
		return ClickedResponse().Render(c.Context(), c.Response().BodyWriter())
	})

	log.Fatal(app.Listen(":3000"))
}
