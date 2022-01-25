package main

import (
	"fmt"
)



func main() {
	// make array de inteiros, tamanho e capacidade
	slice := make([]int, 5, 10)
	// not possible to add the sixth element without an append
	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5
	// the capacity will be 10 until the tenth element
	// After the data a new array will be created with capacity 20
	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)
	slice = append(slice, 11)
	fmt.Println(slice, len(slice), cap(slice))
}