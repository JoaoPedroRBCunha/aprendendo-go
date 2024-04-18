package main

import (
	"fmt"
)

func main() {
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[int]map[string]string{
		1: {
			"nome":      "Paulo",
			"sobrenome": "Araujo",
		},
		2: {
			"nome":      "José",
			"sobrenome": "Farias",
		},
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2[1]["nome"])
	delete(usuario2, 2)
	fmt.Println(usuario2)

	usuario2[3] = map[string]string{
		"nome": "Douglas",
	}
	fmt.Println(usuario2)

}
