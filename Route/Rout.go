package Route

import (
	"bytes"
	"gootmplhtmx/database"
	"gootmplhtmx/model"
	"gootmplhtmx/view"

	"github.com/gofiber/fiber/v3"
	// "golang.org/x/net/html"
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
		task.Status = false
		if err:= database.DB.Create(&task).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		var tasks []model.TASK
		if err:= database.DB.Find(&tasks).Error; err != nil{
			return c.Status(500).SendString(err.Error())
		}
		var b bytes.Buffer
		html:= view.Listing(tasks)
		if err := html.Render(c.Context(), &b); err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).Send(b.Bytes())
	})



	APP.Delete("/delete/:id", func(c fiber.Ctx) error {
		var task model.TASK
		id := c.Params("id")
		database.DB.Delete(&task,id)
		return c.Status(fiber.StatusAccepted).SendString("Deleted")
		

	})

	APP.Patch("/edit/:task", func(c fiber.Ctx) error {
		var task model.TASK
		dbtx := c.Params("tasktx")
		if err:= database.DB.First(&task,"id=?",dbtx).Error; err!=nil{
			return c.Status(fiber.StatusNotFound).SendString("not found for edit")
		}
		if err:= database.DB.Delete(&task).Error; err != nil{
			return c.Status(fiber.StatusNotFound).SendString("Can't update")
		}

		return c.SendString(`
				<input type="text" placeholder="Type here..." id="inpt" name="user-input" value="`+task.Tasktx+`"  />
		
		`)


	})
	
}
