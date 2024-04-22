package main

import (
	"fmt"
)

func main() {
	tarefas, resultados := make(chan int, 45), make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

	/* ou poderia ser feito dessa forma:

	for i := 0; i < 20; i++ {
		tarefas <- i
		resultado := <-resultados
		fmt.Println(resultado)
	}
	close(tarefas)

	*/

}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
