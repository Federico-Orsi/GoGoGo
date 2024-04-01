package routes

import (
	//"fmt"
	//"log"

	//"github.com/Federico-Orsi/Go/controllers"
	"github.com/Federico-Orsi/Go/db"
	//"github.com/Federico-Orsi/Go/jwt_local"
	//"github.com/Federico-Orsi/Go/middlewares"
	"github.com/Federico-Orsi/Go/models"
	"github.com/gofiber/fiber/v2"
)




func Products_Route(app *fiber.App) {


	app.Post("/create_product", func(c *fiber.Ctx) error {

		// product := models.Product{}

		// if err := c.BodyParser(&product); err != nil {
		// 	return err
		// }

        // db.DB.Create(&product)
		// if product.ID == 0 {

		// 	return c.Status(fiber.StatusBadRequest).JSON("Some error")
		// }

		// return c.JSON(product)
//--------------------------------------------------------------------------
		// Iniciar una transacción
		tx := db.DB.Begin()

		// Deferir la función para que la transacción se confirme o revierta al salir de la función
		defer func() {
			if r := recover(); r != nil {
				// Si hubo un pánico, revertir la transacción
				tx.Rollback()
			}
		}()

		// Crear el producto dentro de la transacción
		product := models.Product{}
		if err := c.BodyParser(&product); err != nil {
			// Si hay un error al analizar el cuerpo de la solicitud, revertir la transacción y devolver el error
			tx.Rollback()
			return err
		}

		if err := tx.Create(&product).Error; err != nil {
			// Si hay un error al crear el producto, revertir la transacción y devolver el error
			tx.Rollback()
			return err
		}

		// Commit la transacción si todo ha ido bien
		tx.Commit()

		// Si el producto fue creado exitosamente, devolverlo como JSON
		return c.JSON(product)

	})


	app.Get("/findProduct/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")
        var product models.Product

		db.DB.First(&product, id)

		if product.ID == 0 {

			return c.Status(fiber.StatusNotFound).JSON("Producto no encontrado!!")
		}

		return c.JSON(product)
	})






	app.Get("/products", func(c *fiber.Ctx) error {

		var products []models.Product

        db.DB.Find(&products)

		return c.JSON(products)//Importante saber que Go es distinto con respecto a la respuesta que se manda en el JSON!! Se coloca directamente la variable creada (var). En caso de querer almacenar el resultado de la query en memoria, solamente esto lo hariamos para chequear y manejar un posible error en la consulta a la DB!!
	})




}