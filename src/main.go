package main

import (
	"fmt"
)

func validaSenha(Senha int) (string, bool) {
	if Senha == 2002 {
		return fmt.Sprint("Acesso Permitido"), true
	} else {
		return fmt.Sprintf("Senha Invalida"), false
	}

}

func main() {
	for {
		var Senha int
		fmt.Scan(&Senha)
		texto, ok := validaSenha(Senha)
		if ok == true {
			fmt.Println(texto)
			break
		}
		fmt.Println(texto)
	}
}
