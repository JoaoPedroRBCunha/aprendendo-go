package main

import (
	"fmt"
)

func main() {
	var variavel1 string = "Variável 1"
	// inferência de tipos (variavel2 := "Variável 2)
	variavel2 := "Variável 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		// O tanto de espaço que você colocar antes ou após uma aspa ele irá obdecer e será utilizado.
		variavel3 string = "É só um "
		variavel4 string = "\n teste"
		test      string
	)
	fmt.Println(test, variavel3, variavel4)

	variavel5, variavel6 := "Teste bem aplicado", "Teste mal aplicado"
	// fmt.Println(variavel5 + variavel6) = irá concatenar (juntar) e apresentar as duas variaveis
	// fmt.Println(variavel5, variavel6) = irá separar por um espaço e apresentar as duas variaveis
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	// A constante não precisa necessariamente ser utilizada igual uma variável, posso somente cria-la e não utiliza-la,
	// porém é uma boa prática manter apenas as que estão e irão ser utilizadas
	const (
		constante2 string = ""
		constante3 int    = 3
	)

	// Para trocar os números de lugar é bem simples
	numero1, numero2 := 2, 3
	numero1, numero2 = numero2, numero1
	fmt.Println(numero1, numero2)

}
