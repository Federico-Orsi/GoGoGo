package adapter

//import "fmt"

type BiciAdapter struct {
	Badapter *Bici
}

func (b *BiciAdapter) Transportar() {
	b.Badapter.Pedalear()

}

