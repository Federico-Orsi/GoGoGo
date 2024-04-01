package routes

import (
	"fmt"

	"github.com/Federico-Orsi/Go/adapter"
	"github.com/Federico-Orsi/Go/third_party"
	"github.com/gofiber/fiber/v2"
)

// ---------------------------------------------------------------------------
// En este ejemplo se están sobreescribiendo los valores del struct original
type GokuSuperSaiyan struct{}

func (gsuper GokuSuperSaiyan) Level() string {
	return "Super Saiyan"
}

func (gsuper GokuSuperSaiyan) Technique() string {
	return "Kaioken aumentado 20 veces, el super Hamehamehá y la HenkiDama"
}
//----------------------------------------------------------------------------
// este ejemplo mantiene los valores del struct original tal cual estaban y solamente agrega mas valores:
// Esta es la forma mas utilizada, ya que lo que se busca en un Decorador es que no modifique los objetos previos, sino que solamente agregue funcionalidades extras!!
type Decorator2 struct{
	goku third_party.Goku
}

func (d2 Decorator2) Level() string {
	fmt.Println("agregado de decorator2")
	return d2.goku.Level()
}

func (d2 Decorator2) Technique() string {

	return d2.goku.Technique() + " agregado de decorator2 en el return"
}



func Decorator_Route(app *fiber.App) {

	app.Get("superSaiyan", func(c *fiber.Ctx) error {


		super_goku := GokuSuperSaiyan{} //struct
		third_party.ShowGoku(super_goku) // Se le pasa el struct que implementa la interfaz!!

		return c.JSON("Gokú Super Saiyan bitchhh")
	})


	app.Get("decorator", func(c *fiber.Ctx) error {


		decorator_2 := Decorator2{goku: third_party.BasicGoku{}} //struct
		third_party.ShowGoku(decorator_2) // Se le pasa el struct que implementa la interfaz!!

		return c.JSON("Gokú con decorator2")
	})

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

