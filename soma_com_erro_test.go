package soma_test

import (
	"testing"

	"github.com/Gileno29/golang-tests/soma"
)

func TestSomaComErro(t *testing.T) {
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
		{Valores: comNill, Resultado: 20},
	}

	for _, teste := range testes {
		total := soma.Soma(teste.Valores...)
		if total != teste.Resultado {
			t.Fatalf("Valor esperado: %d - Valor retornado: %d", teste.Resultado, total)
		}
	}
}
