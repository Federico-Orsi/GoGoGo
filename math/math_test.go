package math

import "testing"

func TestSum(t *testing.T) {
    result := Sum(5, 5)
    expected := 10

    if result != expected {
        t.Errorf("Resultado obtenido: %d; Valor esperado: %d", result, expected)
    }
}