package main

import (
	"fmt"
)

func main() {
	var renda, taxa, impostoDeRenda float32
	fmt.Scan(&renda)

	switch {
	case renda <= 2000.00:
		fmt.Println("Isento")
	case renda <= 3000.00:
		taxa = 0.08
		impostoDeRenda = renda * taxa
		fmt.Printf("R$ %f\n", impostoDeRenda)
	case renda <= 4500.00:
		taxa = 0.18
		impostoDeRenda = renda * taxa
		fmt.Printf("R$ %f\n", impostoDeRenda)
	case renda > 4500.00:
		taxa = 0.28
		impostoDeRenda = renda * taxa
		fmt.Printf("R$ %f\n", impostoDeRenda)
	}

}
