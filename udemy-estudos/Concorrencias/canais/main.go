package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá mundo!", canal)

	for {
		mensagem, aberto := <-canal // <-canal - esperando que o canal receba um valor
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	// outro método de printar o canal
	for mensagem2 := range canal {
		fmt.Println(mensagem2)
	}

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // canal <- mandando um valor para dentro do canal
		time.Sleep(time.Second)
	}
	close(canal)
}
