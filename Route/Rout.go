package Route

import (
	"gootmplhtmx/database"
	"gootmplhtmx/model"
	"gootmplhtmx/view"

	"github.com/gofiber/fiber/v3"
)

/*
get / is to take that mean to  display all todod
post /add is to add and to take full update list
path /delet is to delete and to updte list that mean by remove that delete
PATCH /edit is to edit already todo list
*/

var Db = database.DB
func Routing(APP *fiber.App) {

	APP.Get("/", func(c fiber.Ctx) error {
		var tasks []model.TASK
		if err := database.DB.Find(&tasks).Error; err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return view.Fulpage("Todo List", tasks).Render(c.Context(), c.Response().Writer())
	})

	// Placeholder endpoints so your structure is ready for server-side HTMX updates.
	APP.Post("/add", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotImplemented)
	})

	APP.Delete("/delete/:id", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotImplemented)
	})

	APP.Patch("/edit/:id", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotImplemented)
	})
	
}
