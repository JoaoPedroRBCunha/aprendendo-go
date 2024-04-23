package main

import (
	"fmt"
)

func verificaDivisao(X, Y int) {
	if Y == 0 {
		fmt.Println("divisao impossivel")
	} else {
		resultado := float32(X) / float32(Y)
		fmt.Printf("%.1f\n", resultado)
	}
}

func main() {
	var (
		n int
		X int
		Y int
	)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&X, &Y)
		verificaDivisao(X, Y)
	}

}
