package routes

import (
	"github.com/Federico-Orsi/Go/db"

	"github.com/Federico-Orsi/Go/models"
	"github.com/gofiber/fiber/v2"
)




func Tasks_Route(app *fiber.App) {


	app.Post("/create_task", func(c *fiber.Ctx) error {

		task := models.Task{}

		if err := c.BodyParser(&task); err != nil {
			return err
		}

        db.DB.Create(&task)
		if task.ID == 0 {

			return c.Status(fiber.StatusBadRequest).JSON("Por favor debe indicar el Id de Usuario para la tarea solicitada!!")
		}

		return c.JSON(task)

	})



}