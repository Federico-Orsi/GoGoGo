package routes

import (
	"fmt"

	"github.com/Federico-Orsi/Go/builder"

	"github.com/gofiber/fiber/v2"
)



func Builder_Route(app *fiber.App) {


	app.Get("builder", func(c *fiber.Ctx) error {

		// json := &builder.JSON_message{}  // cualquiera de estos 2 los usaríamos en caso que el formato del mensaje lo quisiéramos mandar directo sin que la elección la haga el usuario.
		//xml := &builder.XML_message{}     // cualquiera de estos 2 los usaríamos en caso que el formato del mensaje lo quisiéramos mandar directo sin que la elección la haga el usuario.

		var msg_format string
		fmt.Print("Por favor introduzca si el archivo lo desea en formato json o xml: ")
		fmt.Scan(&msg_format) // esto mismo de pedirle al usuario que ingrese un valor por consola se puede lograr tmb usando el metodo: bufio.NewReader(os.Stdin), es un poco mas complejo pero en algunos casos que necesites datos mas específicos puede sernos útil!!

		format, err := builder.Factory(msg_format)
		if err != nil {
			fmt.Println("Error:", err)
		} else {

			sender := &builder.Sender{}

			sender.SetBuilder(format)
			msg, err := sender.BuildMessage("federicoantonio.orsi@gmail.com", "Hello Golang!!")

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(msg.Body))
		}




		return c.JSON("Trying Builder pattern")
	})

}

