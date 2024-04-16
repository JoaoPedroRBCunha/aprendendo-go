package main

import (
	"fmt"
	"math"
)

func triangulo(A, B, C float64) string {
	if A > B+C && B > A+C && C > A+B {
		return fmt.Sprintln("NAO FORMA TRIANGULO")
	}
	if math.Pow(A, 2) == math.Pow(B, 2)+math.Pow(C, 2) {
		return fmt.Sprintln("TRIANGULO RETANGULO")
	}
	if math.Pow(A, 2) > math.Pow(B, 2)+math.Pow(C, 2) {
		return fmt.Sprintln("TRIANGULO OBTUSANGULO")
	}
	if math.Pow(A, 2) < math.Pow(B, 2)+math.Pow(C, 2) {
		return fmt.Sprintln("TRIANGULO ACUTANGULO")
	}
	if A == B && A == C && B == C {
		return fmt.Sprintln("TRIANGULO EQUILATERO")
	}
	if A == B || A == C || B == C {
		return fmt.Sprintln("TRIANGULO ISOSCELES")
	}
	return ""
}

func main() {
	var A, B, C float64
	fmt.Scan(&A, &B, &C)

	fmt.Println(triangulo(A, B, C))

}
