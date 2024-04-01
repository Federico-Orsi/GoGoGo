package adapter

import "fmt"

type Coche struct {
	Marca string
	Modelo uint64 // todos los uint se refieren a numeros enteros positivos (es decir mayores a cero). En este caso el limite máximo es altisimo. Hay otros uint que tienen un techo mucho mas bajo!!
	Vtv bool // los datos boolean por defecto se inicializar en false!! Salvo que se declaren en true explicitamente!!
}

func (c *Coche) Acelerar() {
	fmt.Println("Acelerando babyyy")
}

func (c *Coche) Verificar() {
	if c.Vtv {
	fmt.Println("vtv OK")
   } else {
	fmt.Println("Debes realizar la vtv")
   }

}