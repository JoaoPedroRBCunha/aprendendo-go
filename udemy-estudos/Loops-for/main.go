package main

import (
	"fmt"
)

func main() {
	i := 0

	for i < 10 {
		i++
		// time.Sleep(time.Second)
		fmt.Println(i)
		fmt.Println("Incrementando i")
	}

	fmt.Println("--------")
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
		fmt.Println("Incrementando j")
	}

	fmt.Println("--------")
	nomes := []string{"João", "Davi", "José"}
	for _, nomes := range nomes {
		fmt.Println(nomes)
	}

	fmt.Println("--------")
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Leoncio",
	}
	fmt.Println("--------")
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	fmt.Println("--------")
	var k int
	for {
		if k > 4 {
			break
		}
		fmt.Println(k)
		k++

	}

}
