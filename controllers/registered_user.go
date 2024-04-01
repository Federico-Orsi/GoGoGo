package controllers

import (
	"github.com/gofiber/fiber/v2"
	//"github.com/golang-jwt/jwt/v5"
)

func Registered_user(c *fiber.Ctx) error {

	user := c.Locals("user_Id").(string)  //para buscar en c.Locals() se pone la clave con la cual se guardó el dato y el tipo de dato guardado!!
	rol := c.Locals("rol").(string)


	return c.JSON("Welcome User: " + user + ", your actual rol is " + rol + ".")
}