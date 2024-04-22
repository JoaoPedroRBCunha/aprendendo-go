package main

import (
	"fmt"
)

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel2++
	fmt.Println(variavel1, variavel2)

	// Ponteiro é uma referência de memória
	var variavel3 int = 10
	var variavel4 *int = &variavel3 // *int (pode ser int, float, string) - identifica que é uma criação de um ponteiro
	variavel5 := &variavel3         // &variavel3 - utilizar esse "&" faz com que pegue a referência daquela variável

	fmt.Println(variavel3, variavel4, variavel5)

	variavel3++
	fmt.Println(variavel3, *variavel4, *variavel5)
	// *variavel4 - utilizar esse "*" (asterisco) é uma desreferênciação,
	// pois ele irá pegar o valor da variável e não o local onde está alocada
}
