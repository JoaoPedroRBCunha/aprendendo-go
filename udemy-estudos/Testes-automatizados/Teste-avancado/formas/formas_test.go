package formas_test

import (
	"math"
	. "teste-avancado/formas"
	"testing"
)

func TestArea(t *testing.T) {
	// TDD - Test Drive Development
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da área esperada %f\n", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da área esperada %f\n", areaRecebida, areaEsperada)
		}
	})
}
