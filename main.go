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
	APP := fiber.New()
	APP.Use(logger.New())
	APP.Use(recover.New())
	APP.Use("/static", static.New("./public"))
	Route.Routing(APP)
	// Print all registered routes
	for _, route := range APP.GetRoutes(true) {
		log.Printf("Route: %s %s", route.Method, route.Path)
	}
	log.Fatal(APP.Listen(":4000"))
}
