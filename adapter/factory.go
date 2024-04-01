package adapter

import "errors"

func Factory(s string) (Adapter, error) {
	switch s {
	case "coche":
		c := &Coche{Vtv: true}
		return &CocheAdapter{c}, nil
	case "bici":
		return &BiciAdapter{&Bici{}}, nil
	}

	return nil, errors.New("upsss opción no válida") // si el cliente no introduce ninguna de las opciones indicadas en los case, retorna nil
}