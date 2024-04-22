package main

import (
	"fmt"
)

func main() {
	// ARITIMÉTICOS
	soma := 1 + 3
	subtracao := 5 - 2
	multiplicacao := 5 * 2
	divisao := 10 / 4
	restoDaDivisao := 10 % 4

	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)

	// Não pode somar números que tem o tipo de dado diferente, como exemplo:
	// var numero1 int16 = 10
	// var numero2 int16 = 8
	// soma := numero1 + numero2
	var numero1 int16 = 10
	var numero2 int16 = 8
	var soma2 int16 = numero1 + numero2
	fmt.Println(soma2)
	// Pode usar inferência de tipos com duas variáveis iguais
	soma3 := numero1 + numero2
	fmt.Println(soma3)

	// FIM DOS ARITIMÉTICOS

	// ATRIBUIÇÃO
	fmt.Println("--------------")
	var variavel1 string = "String1"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// FIM DOS OPERADOS DE ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	fmt.Println("--------------")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(2 == 2)
	fmt.Println(3 <= 4)
	fmt.Println(3 < 2)
	fmt.Println(2 != 3)

	// FIM DOS OPERADORES RELACIONAIS

	// OPERADORES LÓGICOS
	fmt.Println("--------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// FIM DOS OPERADORES LÓGICOS

	// OPERADORES UNÁRIOS
	fmt.Println("--------------")
	numero := 10
	numero++     // numero = numero + 1
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--    // numero = numero - 1
	numero -= 5 // numero = numero - 5
	fmt.Println(numero)

	numero *= 3 // numero = numero * 3
	fmt.Println(numero)
	numero /= 2 // numero = numero / 2
	fmt.Println(numero)

	// FIM DOS OPERADORES UNÁRIOS

}
