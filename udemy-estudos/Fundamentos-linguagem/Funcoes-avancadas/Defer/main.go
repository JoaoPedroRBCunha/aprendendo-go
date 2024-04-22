package main

import (
	"fmt"
)

func funcao1() {
	fmt.Println("Função 1")
}

func funcao2() {
	fmt.Println("Função 2")
}

func alunosAprovados(n1, n2 float32) string {
	defer fmt.Print("Status: ")
	fmt.Println("Verificando status")
	media := float32((n1 + n2) / 2)
	if media >= 6 {
		return fmt.Sprintf("Aprovado")
	}
	return fmt.Sprintf("Reprovado")
}

func main() {
	defer funcao1() // defer == adiar
	funcao2()
	fmt.Println(alunosAprovados(7, 8))
}
