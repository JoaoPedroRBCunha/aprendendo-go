package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a, area float64
	)
	pi := 3.14159
	fmt.Scan(&a)
	area = pi * math.Pow(a, 2)
	fmt.Printf("A=%.4f\n", area)
}
