package adapter

//import "fmt"

type CocheAdapter struct {
	Cadapter *Coche
}

func (c *CocheAdapter) Transportar() {
	c.Cadapter.Verificar()
	c.Cadapter.Acelerar()


}

