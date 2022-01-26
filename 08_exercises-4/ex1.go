package main

import (
	"fmt"
)

var numbers [5]int


func main() {
	numbers[0], numbers[1], numbers[2], numbers[3], numbers[4] = 1, 2, 3, 4, 5
	for key, value := range numbers {
		fmt.Printf("Key: %v - Value: %v\n", key, value)
	}

	fmt.Printf("type: %T", numbers)
}