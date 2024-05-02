package main

import (
	"fmt"
	"reflect"
)

func main() {
	test := 30
	var hello string = "Hello World!"
	fmt.Println(hello)
	fmt.Println("Esta variável é do tipo ", reflect.TypeOf(test))
}
