package Route

import (
	"bytes"
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

func Routing(APP *fiber.App) {
	
	APP.Get("/", func(c fiber.Ctx) error {
		var task []model.TASK
		if err := database.DB.Find(&task).Error; err != nil {
			return c.Status(500).SendString(err.Error())
		}
		var b bytes.Buffer
		html := view.Fulpage("todolist", task)
		if err := html.Render(c.Context(), &b); err != nil {
			return err
		}
		return c.Status(200).Send(b.Bytes())
	})

	// Placeholder endpoints so your structure is ready for server-side HTMX updates.
	APP.Post("/add", func(c fiber.Ctx) error {
		var task model.TASK
		text := c.FormValue("user-input")
		task.Tasktx = text
		if err:= database.DB.Create(task).Error; err != nil{
		return c.SendString("database is not accept ", err.Error())
		}

		// return c.SendStatus(fiber.StatusNotImplemented)

	})

	// APP.Delete("/delete/:id", func(c fiber.Ctx) error {
	// 	return c.SendStatus(fiber.StatusNotImplemented)
	// })

	// APP.Patch("/edit/:id", func(c fiber.Ctx) error {
	// 	return c.SendStatus(fiber.StatusNotImplemented)
	// })
	
}
