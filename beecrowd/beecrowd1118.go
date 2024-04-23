package main

import (
	"fmt"
)

func notaValida() float32 {
	var (
		nota float32
	)
	fmt.Scan(&nota)
	if nota >= 0 && nota <= 10 {
		return nota
	} else {
		fmt.Println("nota invalida")
		return notaValida()
	}
}

func novoCalculo() bool {
	var novoCalculo int
	fmt.Println("novo calculo (1-sim 2-nao)")
	fmt.Scan(&novoCalculo)
	if novoCalculo != 1 && novoCalculo != 2 {
		for {
			fmt.Println("novo calculo (1-sim 2-nao)")
			fmt.Scan(&novoCalculo)
			if novoCalculo == 1 || novoCalculo == 2 {
				break
			}
		}
	}
	if novoCalculo == 2 {
		return false
	} else if novoCalculo == 1 {
		return true
	}
	return false
}

func main() {
	for {
		nota1 := notaValida()
		nota2 := notaValida()
		media := (nota1 + nota2) / 2
		fmt.Printf("media = %.2f\n", media)
		if calculaNovamente := novoCalculo(); calculaNovamente == false {
			break
		}
	}
}
