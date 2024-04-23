package main

import (
	"fmt"
)

func main() {
	var numeros, positivos int

	for i := 0; i < 6; i++ {
		fmt.Scan(&numeros)
		if numeros > 0 {
			positivos++
		}
	}

	fmt.Printf("%d valores positivos\n", positivos)
}
