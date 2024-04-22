package main

import (
	"fmt"
)

var n int

func init() {
	n = 10
	fmt.Println("Função init")
}

func main() {
	fmt.Println("Função main")
	fmt.Println(n)
}
