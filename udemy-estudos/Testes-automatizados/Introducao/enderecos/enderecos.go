package enderecos

import (
	"strings"
)

// TipoDeEndereco verifica se o tipo de endereço tem um tipo válido e o retorna
func TiposDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecoLetraMinuscula, " ")[0]
	/* strings.Split(enderecoLetraMinuscula, " ")[0]
	RUA DOS BOBOS = ["RUA", "DOS", "BOBOS"]
	*/

	enderecoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoValido = true
		}
	}

	if enderecoValido {
		return strings.Title(primeiraPalavraEndereco)
	}
	return "Tipo Inválido"

}
