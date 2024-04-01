package routes

import (
	"github.com/Federico-Orsi/Go/observer"
	"github.com/gofiber/fiber/v2"
)

func Observer_Route(app *fiber.App) {

	app.Get("observer", func(c *fiber.Ctx) error {

       m := observer.Message{}
	   m.Msg = "Come onnnn Goooo!! We fucking made it brooooo!!"
	   email := observer.Email{}
	   wsap := observer.WhatsApp{}

	   m.Add_Observer("email", &email)
	   m.Add_Observer("wasaaa", &wsap)
	   m.Remove_Observer("email")

	   m2 := observer.Message{}
	   m2.Msg = "Come onnnn Goooo!! We fucking made it Mannnn!!"
	   email2 := observer.Email{}
	   m2.Add_Observer("email2", &email2)

	   m.Notify_Observers()
	   m2.Notify_Observers()


		return c.JSON("Trying Observer pattern")
	})

}
