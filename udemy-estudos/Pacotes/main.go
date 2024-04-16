package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	// go mod init [nome do modulo] - para iniciar o modulo e poder importar bibliotecas
	fmt.Println("test")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)

	erro2 := checkmail.ValidateFormat("3")
	fmt.Println(erro2)
}
