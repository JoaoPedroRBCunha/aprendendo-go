package main

import "fmt"

func main() {
	var X, Y int
	fmt.Scanln(&X, &Y)

	count := 1
	for i := 1; i <= Y; i++ {
		fmt.Print(i)
		if count < X && i != Y {
			fmt.Print(" ")
			count++
		} else {
			fmt.Println()
			count = 1
		}
	}
}
