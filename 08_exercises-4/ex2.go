package main

import (
	"fmt"
)

var numbers [5]int


func main() {
	slice := []int{1,2,3,4,5,6,7,8,9,10}
	for _, value := range slice {
		fmt.Printf("Value: %v\n", value)
	}

	fmt.Printf("type: %T", slice)
}