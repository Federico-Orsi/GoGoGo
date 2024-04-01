package routes

import (
	"fmt"
	"strconv"

	//"log"

	//"github.com/Federico-Orsi/Go/controllers"
	"github.com/Federico-Orsi/Go/db"
	//"github.com/Federico-Orsi/Go/jwt_local"
	//"github.com/Federico-Orsi/Go/middlewares"
	"github.com/Federico-Orsi/Go/models"
	"github.com/gofiber/fiber/v2"
)




func Orders_Route(app *fiber.App) {


	app.Post("/add_to_cart/:cart_id", func(c *fiber.Ctx) error {

		order := models.OrderProduct{}

		if err := c.BodyParser(&order); err != nil {
			return err
		}

		cart_id := c.Params("cart_id")
		fmt.Println(cart_id)
        var cart models.Cart

		db.DB.First(&cart, cart_id)

		if cart.ID == 0 {

			return c.Status(fiber.StatusNotFound).JSON("Carrito no encontrado!!")
		}

		var product models.Product
        db.DB.First(&product, order.Product_id)
        if product.ID == 0 {

			return c.Status(fiber.StatusNotFound).JSON("Producto no encontrado!!")
		}

		cart_id_int, err := strconv.Atoi(cart_id) // este metodo convierte el string de un numero, en un numero entero!!
	if err != nil {
		fmt.Println("Error al convertir la cadena a entero:", err)

	}
		order.CartID = uint(cart_id_int) // esto convierte un entero en uint

        db.DB.Create(&order)
		if order.ID == 0 {

			return c.Status(fiber.StatusBadRequest).JSON("Some error with the order")
		}

		return c.JSON(order)

	})








}