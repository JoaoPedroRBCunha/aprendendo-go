package main

import (
	"fmt"
)

func main() {
	var (
		ddd     int
		destino string
	)

	fmt.Scan(&ddd)
	switch {
	case ddd == 61:
		destino = "Brasilia"
	case ddd == 71:
		destino = "Salvador"
	case ddd == 11:
		destino = "Sao Paulo"
	case ddd == 21:
		destino = "Rio de Janeiro"
	case ddd == 32:
		destino = "Juiz de Fora"
	case ddd == 19:
		destino = "Campinas"
	case ddd == 27:
		destino = "Vitoria"
	case ddd == 31:
		destino = "Belo Horizonte"
	default:
		destino = "DDD nao cadastrado"
	}

	fmt.Printf("%s\n", destino)

}
