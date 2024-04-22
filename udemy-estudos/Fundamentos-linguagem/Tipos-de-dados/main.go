package main

import (
	"errors"
	"fmt"
)

func main() {
	// INT
	var numero int8 = 100
	fmt.Println(numero)

	// var numero2 uint32 = -10
	// uint não pega número negativo

	// alias - termo utilizado
	// int32 == rune
	var numero3 rune = 3
	fmt.Println(numero3)

	// alias - termo utilizado
	// uint8 = byte
	var numero4 byte = 4
	fmt.Println(numero4)

	// FIM INT

	// FLOAT
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1340000.50
	fmt.Println(numeroReal2)

	numeroReal3 := 3000.40
	fmt.Println(numeroReal3)

	// FIM FLOAT

	// STRING
	var str string = "String"
	fmt.Println(str)

	str2 := "texto string 2"
	fmt.Println(str2)

	// aspas simples é o mais próximo de Char
	char := 'B'
	fmt.Println(char)

	// FIM STRING

	var texto string
	fmt.Println(texto)

	var boolean1 bool
	fmt.Println(boolean1)

	var erro error = errors.New("Inválido")
	fmt.Println(erro)
}
