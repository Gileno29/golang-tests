package soma

import (
	"testing"
)

func TestSoma(t *testing.T) {
	test := []int{1, 2, 3, 4}
	total := Soma(test...)

	if total != 10 {
		t.Fatalf("Valor esperado: 10 - Valor retornado: %d", total)
	}
}
