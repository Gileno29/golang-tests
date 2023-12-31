package soma

import (
	"testing"
)

func TestSoma02(t *testing.T) {
	testes := []struct {
		Valores   []int
		Resultado int
	}{
		{Valores: []int{1, 2, 3}, Resultado: 6},
		{Valores: []int{1, 2, 3, 4}, Resultado: 10},
		{Valores: []int{3, 3, 3, 3}, Resultado: 12},
		{Valores: []int{1, 1, 1, 1}, Resultado: 4},
		{Valores: []int{12, 20, 35}, Resultado: 67},
		{Valores: []int{19, 21, 32}, Resultado: 72},
	}

	for _, teste := range testes {
		total := Soma(teste.Valores...)
		if total != teste.Resultado {
			t.Fatalf("Valor esperado: %d - Valor retornado: %d", teste.Resultado, total)
		}
	}
}
