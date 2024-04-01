package middlewares

import (
	"fmt"
	"strconv"

	"github.com/Federico-Orsi/Go/db"
	"github.com/Federico-Orsi/Go/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(c *fiber.Ctx) error {
 //obtengo info real del user loggeado!!
	login_user := c.Locals("user").(*jwt.Token) //"user": is the Context key to store user information from the token into context.
	claims := login_user.Claims.(jwt.MapClaims)
	user_id := claims["user_id"].(string)

	c.Locals("user_Id", user_id) // con c.Locals() en Fiber guardamos datos de clave valor, los cuales luego pueden ser utilizados en el resto de middlewares o controller!!

    var user models.User


	num, err := strconv.Atoi(user_id) // este metodo convierte el string de un numero, en un numero entero!!
	if err != nil {
		fmt.Println("Error al convertir la cadena a entero:", err)

	}
	db.DB.First(&user, num) //Ojo!! el id que le pasamos a Gorm para que busque debe ser un entero!!
    c.Locals("rol", user.Rol)


    if user.Rol == "Premium" || user.Rol == "Admin" {
      return c.Next()
	} else{
		return c.JSON("Upss!! Acceso permitido solo a usuarios Premium y Admin")
	}
}