package main

import (
	"fmt"
)



func main() {
	firstslice := []int{1, 2, 3, 4, 5}
	secondslice := append(firstslice[1:3], firstslice[4:]...)
	fmt.Println(secondslice)
	fmt.Println(firstslice)

}