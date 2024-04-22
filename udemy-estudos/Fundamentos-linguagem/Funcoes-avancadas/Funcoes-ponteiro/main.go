package main

import (
	"fmt"
)

func inverterNumero(numero int) int {
	return numero * -1
}

func inverterNumeroComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := -20
	fmt.Println(inverterNumero(numero))
	fmt.Println(numero)

	numeroNovo := -100
	fmt.Println(numeroNovo)
	inverterNumeroComPonteiro(&numeroNovo)
	fmt.Println(numeroNovo)

}
