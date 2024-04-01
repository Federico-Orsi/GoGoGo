package routes

import (
	//"fmt"
	//"log"

	//"github.com/Federico-Orsi/Go/controllers"

	"fmt"

	"github.com/Federico-Orsi/Go/db"

	//"github.com/Federico-Orsi/Go/jwt_local"
	//"github.com/Federico-Orsi/Go/middlewares"
	"github.com/Federico-Orsi/Go/models"
	"github.com/gofiber/fiber/v2"
)




func Carts_Route(app *fiber.App) {


	app.Get("/findCart/:cid", func(c *fiber.Ctx) error {

		cid := c.Params("cid")
		var cart models.Cart
		db.DB.First(&cart, cid)

		db.DB.Preload("User").Preload("User.Task").Preload("OrderProducts.Product").Preload("OrderProducts").First(&cart, cid)

		if cart.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON("Carrito no encontrado!!")
		}

		var user models.User
		db.DB.First(&user, cart.User_id)
		fmt.Println(user)

		//var prod models.Product
		//db.DB.First(&prod, cart.OrderProducts[1].Product_id)
		//fmt.Println(prod)


		return c.JSON(cart)
	})





}