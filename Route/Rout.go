package Route

import (
	"bytes"
	// "fmt"
	"log"
	"gootmplhtmx/database"
	"gootmplhtmx/model"
	"gootmplhtmx/view"
	"html"

	"github.com/gofiber/fiber/v3"
	// "golang.org/x/net/html"
)

func Routing(APP *fiber.App) {

	APP.Get("/", func(c fiber.Ctx) error {
		var tasks []model.TASK
		if err := database.DB.Find(&tasks).Error; err != nil {
			log.Printf("DB Error: %v", err)
			return c.Status(500).SendString("Database error")
		}  // ← add this closing brace

		var b bytes.Buffer
		component := view.Fulpage("Todo List", tasks)

		if err := component.Render(c.Context(), &b); err != nil {
			log.Printf("Render Error: %v", err)
			return c.Status(500).SendString("Error rendering template")
		}
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.Status(200).SendString(b.String())
	})

	// Placeholder endpoints so your structure is ready for server-side HTMX updates.

	APP.Post("/add", func(c fiber.Ctx) error {
		var task model.TASK
		text := c.FormValue("user-input")
		task.Tasktx = text
		task.Status = false
		if err := database.DB.Create(&task).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		var tasks []model.TASK
		if err := database.DB.Find(&tasks).Error; err != nil {
			return c.Status(500).SendString(err.Error())
		}
		var b bytes.Buffer
		html := view.Listing(tasks)
		if err := html.Render(c.Context(), &b); err != nil {
			return err
		}
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.Status(fiber.StatusOK).Send(b.Bytes())
	})

	APP.Delete("/delete/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		var task model.TASK
		if err := database.DB.Delete(&task, id).Error; err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Status(200).SendString("")
	})

		APP.Patch("/edit/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		var task model.TASK
		if err := database.DB.First(&task, "id = ?", id).Error; err != nil {
			return c.Status(404).SendString("not found")
		}

		
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.Status(200).SendString(`
			<div class="todoList todoList--edit">
				<div class="todoListI todoListI--edit">
					<div class="taskMeta">Editing task #` + id + `</div>
					<input type="text" class="editInput" name="user-input" value="` + html.EscapeString(task.Tasktx) + `" />
				</div>
				<div class="todoListO todoListO--edit">
					<button class="btnO" hx-get="/" hx-target="#list-container" hx-swap="innerHTML">Cancel</button>
				</div>
			</div>
		`)
	})

}
