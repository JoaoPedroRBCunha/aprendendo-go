// TESTE DE UNIDADE
package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTiposDeEndereco(t *testing.T) {
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tiposDeEnderecoRecebido := TiposDeEndereco(cenario.enderecoInserido)
		if tiposDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
				cenario.retornoEsperado,
				tiposDeEnderecoRecebido)
		}
	}

	/*
		enderecoParaTeste := "Avenida Paulista"
		tipoDeEnderecoEsperado := "Avenida"
		tipoDeEnderecoRecebido := TiposDeEndereco(enderecoParaTeste)

		if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
				tipoDeEnderecoEsperado,
				tipoDeEnderecoRecebido,
			)
		}
	*/

}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Error("Teste quebrou!")
	}
}
