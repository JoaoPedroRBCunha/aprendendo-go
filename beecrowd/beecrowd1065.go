package main

import (
	"fmt"
)

func main() {
	var numeros, pares int

	for i := 0; i < 5; i++ {
		fmt.Scan(&numeros)
		if numeros%2 == 0 {
			pares++
		}
	}

	fmt.Printf("%d valores pares\n", pares)
}
