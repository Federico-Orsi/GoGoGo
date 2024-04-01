package routes

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func TestRoute(app *fiber.App) {

    test := app.Group("/test")

	test.Use(jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{Key: []byte("s3cr3t")},
    }))


    test.Get("/", func(c *fiber.Ctx) error {
     //c.Locals() se utiliza para almacenar info en el contexto local de la peticion!! Esa info luego puede ser enviada con Next() al proximo middleware o controller!!

	    user := c.Locals("user").(*jwt.Token)  //"user": is the Context key to store user information from the token into context.
		claims := user.Claims.(jwt.MapClaims)
		user_id := claims["user_id"].(string)

		return c.SendString("Welcome User under id: " + user_id )

	})

	test.Get("/balcon", func(c *fiber.Ctx) error {
		return c.SendString("Hello 43!! Let's fucking Goooooooooooo!!")
	})


}