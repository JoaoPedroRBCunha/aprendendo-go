package main

import (
	"fmt"
)

func somar(n1, n2 int8) int8 {
	return n1 + n2
}

// Em go você pode retornar dois valores na mesma função.
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(30, 20)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Função f")
	}

	f()

	var texto = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := texto("O texto foi feito")
	fmt.Println(resultado)

	fmt.Println(calculosMatematicos(10, 20))

	resultadoSoma, resultadoSubtracao := calculosMatematicos(15, 10)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// O _ (underline) vai ignorar o valor, e vai ser inexistente sem receber o valor.
	resultadoSoma2, _ := calculosMatematicos(15, 10)
	fmt.Println(resultadoSoma2)

}
