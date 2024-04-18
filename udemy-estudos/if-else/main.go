package main

import (
	"fmt"
)

func main() {
	numero := 15

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else if numero == 15 {
		fmt.Println("Igual a 15")
	} else {
		fmt.Println("Menor que 15")
	}

	if outroNumero := numero; outroNumero > 0 { // if init
		fmt.Println("Maior que 0")
	} else {
		fmt.Println("Menor ou igual a 0")
	}

	// fmt.Println(outroNumero) a variável criada no if init não pode ser acessada fora do seu if criado.

}
