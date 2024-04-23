package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		fmt.Printf("%d %d %d\n", i, i*i, i*i*i)
		fmt.Printf("%d %d %d\n", i, i*i+1, i*i*i+1)
	}
}
