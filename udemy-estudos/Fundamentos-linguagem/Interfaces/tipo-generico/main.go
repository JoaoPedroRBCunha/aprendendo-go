package main

import (
	"fmt"
)

func generico(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generico("String")
	generico(1)
	generico(true)
}
