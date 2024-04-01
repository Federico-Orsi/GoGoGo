package builder

import (
	"errors"
)


func Factory(s string) (MessageBuilder, error) {
	switch s {
	case "json":
		return &JSON_message{}, nil
	case "xml":
		return &XML_message{}, nil
	}

	return nil, errors.New("upsss opción no válida") // si el cliente no introduce ninguna de las opciones indicadas en los case, retorna nil
}