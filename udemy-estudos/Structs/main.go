package main

import (
	"fmt"
)

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo struct")

	var u usuario
	u.nome = "Davi"
	u.idade = 5
	fmt.Println(u)
	fmt.Println(u.nome)
	fmt.Println(u.idade)

	enderecoExemplo := endereco{"Rua dos Alpinistas", 2}
	usuario2 := usuario{"José", 2, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 10}
	fmt.Println(usuario3)
}
