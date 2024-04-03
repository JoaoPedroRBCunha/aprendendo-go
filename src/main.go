package main

import (
	"fmt"
	"log"
	"os"
)

var (
	cara, coroa int
)

func lancarMoeda(lado string) {
	switch lado {
	case "cara":
		cara++

	case "coroa":
		coroa++

	default:
		fmt.Println("Caiu em pé")
	}
}

func main() {
	a, b := 10, 10
	if a > b {
		fmt.Println("a é maior que b")
	} else if a < b {
		fmt.Println("b é maior que a")
	} else {
		fmt.Println("a é igual a b")
	}

	switch {
	case a > b:
		fmt.Println("a é maior que b")

	case a < b:
		fmt.Println("b é maior que a")

	default:
		fmt.Println("a é igual a b")
	}

	file, err := os.Open("Hello.txt")
	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)
	if _, erro := file.Read(data); erro != nil {
		log.Panic(erro)
	}

	fmt.Println(string(data))

}
