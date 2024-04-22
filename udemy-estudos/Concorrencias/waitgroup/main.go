package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func(texto string) {
		escrever(texto)
		waitGroup.Done() // -1
	}("Olá mundo!")

	go func(texto string) {
		escrever(texto)
		waitGroup.Done() // -1
	}("Programando em Go!")

	go func(texto string) {
		escrever(texto)
		waitGroup.Done() // -1
	}("GoRoutine 3!")

	go func(texto string) {
		escrever(texto)
		waitGroup.Done() // -1
	}("GoRoutine 4!")

	waitGroup.Wait() // Espera a contagem das goroutines chegar em 0
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
