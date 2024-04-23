package main

import (
	"fmt"
)

func coordenadas(X, Y int) (string, bool) {
	if X < Y {
		return fmt.Sprintf("Crescente"), true
	} else if X > Y {
		return fmt.Sprintf("Decrescente"), true
	} else {
		return fmt.Sprintf(""), false
	}
}

func main() {
	for {
		var X, Y int
		fmt.Scan(&X, &Y)
		coordenadas, ok := coordenadas(X, Y)
		if ok == false {
			break
		}
		fmt.Println(coordenadas)

	}
}
