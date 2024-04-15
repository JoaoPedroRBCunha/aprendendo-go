package main

import (
	"fmt"
)

func main() {
	var (
		salario    float64
		percentual int
	)

	fmt.Scan(&salario)
	switch {
	case salario <= 400.00:
		percentual = 15
	case salario <= 800.00:
		percentual = 12
	case salario <= 1200.00:
		percentual = 10
	case salario <= 2000.00:
		percentual = 7
	default:
		percentual = 4
	}

	novoSalario := salario + (salario * (float64(percentual) / 100))
	reajuste := novoSalario - salario

	fmt.Printf("Novo salario: %.2f\n", novoSalario)
	fmt.Printf("Reajuste ganho: %.2f\n", reajuste)
	fmt.Printf("Em percentual: %d %%\n", percentual)
}
