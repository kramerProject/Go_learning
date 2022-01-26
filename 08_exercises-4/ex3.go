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

	beforelast := len(slice) - 1

	fmt.Printf("1 to 3----%v\n", slice[:3])
	fmt.Printf("5 to last----%v\n", slice[4:])
	fmt.Printf("2 to 7----%v\n", slice[1:7])
	fmt.Printf("3 to 9----%v\n", slice[2:9])
	fmt.Printf("3 to 9----%v\n", slice[2:beforelast])
}