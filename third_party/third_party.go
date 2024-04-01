package third_party

import "log"

type Goku interface {
	Level() string
	Technique() string
}

type BasicGoku struct{}

func (bg BasicGoku) Level() string {
	return "Normal"
}

func (bg BasicGoku) Technique() string {
	return "Kaioken aumentado 20 veces"
}

func ShowGoku(goku Goku) {
	log.Printf("El nivel actual de Gokú es: %v", goku.Level())
	log.Printf("Las técnicas con este nivel de Gokú son: %v", goku.Technique())

}