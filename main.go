package main

import (
	"log"
	"os"

	"gootmplhtmx/Route"
	"gootmplhtmx/database"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	if _, err := database.DbLoad(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use("/static", static.New("./public"))
	Route.Routing(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "4000"
	}

	log.Fatal(app.Listen(":" + port))
}
