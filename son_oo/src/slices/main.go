package main

import "fmt"

func main() {
	slice := make([]int, 2, 5)

	slice = append(slice, 10, 2, 5, 12)

	fmt.Println(slice)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// for i := 0; i < 20; i++ {
	// 	slice = append(slice, i)
	// 	fmt.Println("Len: ", len(slice), "Cap: ", cap(slice))
	// }

	// cria apontamento
	// sliceTest := slice
	// slice = append(slice, 1, 2, 3, 4, 5)
	// slice[0] = 10
	// fmt.Println(slice)
	// fmt.Println(sliceTest)

	sliceStr := []string{"Hello", "World", "Much", "Better"}

	fmt.Println(sliceStr)
	fmt.Println(sliceStr[:2])
	fmt.Println(sliceStr[2:4])

}
