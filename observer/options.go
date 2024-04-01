package observer

import "fmt"

type Email struct{}

func (e *Email) Notify(data string) {
	fmt.Println("Se ha enviado un e-mail a federicoantonio.orsi@gmai.com con el siguiente mensaje: ", data)
}

type WhatsApp struct{}

func (e *WhatsApp) Notify(data string) {
	fmt.Println("Se ha enviado un WhatsApp al movil +54(911)5815-7310 con el siguiente mensaje: ", data)
}