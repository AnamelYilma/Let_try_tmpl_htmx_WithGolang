package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/static"

	"gootmplhtmx/Route"
	"gootmplhtmx/database"
)

func main() {
	if _, err := database.DbLoad(); err != nil {
		log.Fatal(err)
	}

	// Create Fiber app with v3
	APP := fiber.New()

	// Add middleware
	APP.Use(logger.New())
	APP.Use(recover.New())
	APP.Use("/static", static.New("./public"))
	Route.Routing(APP)

	log.Fatal(APP.Listen(":3000"))
	
}
