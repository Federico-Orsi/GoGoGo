package adapter

import "fmt"

type Bici struct {
	Marca string

}

func (b *Bici) Pedalear() {
	fmt.Println("Aguante la bici papuriiiii!!")
}

