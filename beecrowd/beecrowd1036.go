package main

import (
	"fmt"
	"math"
)

func verificaBhaskara(A, Delta float64) bool {
	if A == 0 || Delta < 0 {
		return false
	} else {
		return true
	}
}

func main() {
	var A, B, C float64
	fmt.Scan(&A, &B, &C)
	delta := ((math.Pow(B, 2)) - (4 * A * C))

	if verificaBhaskara(A, delta) == true {
		R1 := (-B + math.Sqrt(delta)) / (2 * A)
		R2 := (-B - math.Sqrt(delta)) / (2 * A)
		fmt.Printf("R1 = %.5f\n", R1)
		fmt.Printf("R2 = %.5f\n", R2)
	} else {
		fmt.Println("Impossivel calcular")
	}

}
