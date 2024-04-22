package main

import (
	"fmt"
)

func soma(numeros ...int) int {
	total := 0
	for _, numeros := range numeros {
		total += numeros
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	valorDaSoma := soma(1, 2, 3, 4, 5, 6, 200, 102, 12, 13)
	fmt.Println(valorDaSoma)

	escrever("Olá mundo", 10, 20, 30, 40, 50)

}
