package main

import (
	"fmt"
)

func main() {
	var local int8
	ddd := map[int8]string{
		61: "Brasilia",
		71: "Salvador",
		11: "São Paulo",
		21: "Rio de Janeiro",
		32: "Juiz de fora",
		19: "Campinas",
		27: "Vitoria",
		31: "Belo Horizonte",
	}

	fmt.Scan(&local)
	cidade, ok := ddd[local]

	if ok == true {
		fmt.Println(cidade)
	} else {
		fmt.Println("DDD nao cadastrado")
	}

}
