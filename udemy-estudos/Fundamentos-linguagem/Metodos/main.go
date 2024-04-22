package main

import (
	"fmt"
)

type usuario struct {
	nome  string
	idade int
}

func (u usuario) salvar() {
	fmt.Printf("O usuário %s foi salvo com sucesso\n", u.nome)
}

func (u usuario) verficandoIdade(idade int) string {
	if idade >= 18 {
		return fmt.Sprintln("É maior de idade")
	}
	return fmt.Sprintln("É menor de idade")
}

func (u *usuario) fazendoAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Davi", 22}
	usuario1.salvar()
	fmt.Print(usuario1.verficandoIdade(18))
	usuario1.fazendoAniversario()
	fmt.Println(usuario1.idade)

	usuario2 := usuario{"José", 30}
	usuario2.salvar()
	fmt.Print(usuario2.verficandoIdade(18))
	fmt.Println(usuario2.idade)

}
