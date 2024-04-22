package main

import (
	"fmt"
)

func main() {
	a := func(texto string) string {
		return fmt.Sprintf("3 %s", texto)
	}("testando")
	fmt.Println(a)
}
