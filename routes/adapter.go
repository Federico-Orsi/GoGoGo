package routes

import (
	"fmt"

	"github.com/Federico-Orsi/Go/adapter"

	"github.com/gofiber/fiber/v2"
)



func Adapter_Route(app *fiber.App) {


	app.Get("adapter", func(c *fiber.Ctx) error {

		var transporte string
		fmt.Print("Por favor introduzca el transporte que utilizará: ")
		fmt.Scan(&transporte)

		transporter, err := adapter.Factory(transporte)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
		transporter.Transportar()

		}


		return c.JSON("Trying adapter")
	})

}

