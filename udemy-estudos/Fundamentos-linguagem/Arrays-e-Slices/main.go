package main

import (
	"fmt"
)

func main() {
	// A diferença da criação de um array pra um slice é no definição de elementos,
	// pois no slice não se define tamanho e no array sim
	var array1 [5]string
	array1[0] = "Posição 1"

	fmt.Println(array1[0])

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Slice é uma fatia de um array.
	// len(slice) - número de elementos dentro do slice.
	// cap(slice) - capacidade máxima de números dentro do slice, porém com append(slice, ...) você pode aumentar essa capacidade.
	slice := []int{10, 11, 12, 13, 14, 15, 16}
	fmt.Println(slice, len(slice), cap(slice))

	// append(slice, ...) - adiciona mais um ou vários elementos dentro do slice e aumenta sua capacidade.
	slice = append(slice, 18)
	fmt.Println(slice, len(slice), cap(slice))

	slices := append(slice, 189, 151, 132, 53)
	fmt.Println(slices, len(slices), cap(slices))

	slice2 := array2[1:3]
	fmt.Println(slice2)

	// Slice pode ser criado pegando todos os valores de um array e adicionando mais alguns
	slice3 := append(array2[0:5], "Test", "Posição 7")
	fmt.Println(slice3)

	// Slice também funciona como ponteiro
	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	// Arrays Internos
	fmt.Println("--------")
	slice4 := make([]float32, 10, 11)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	// Após a capacidade do slice ser estourada o Go vai aumentar um de tamanho (len - lenght) e dobrar sua capacidade.
	slice4 = append(slice4, 5, 7)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	// Slice pode ser criado com o make sem o parametro de capacidade.
	slice5 := make([]float32, 5)
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))

}
