package main

import (
	"fmt"
)




func main() {
	slice := []string{"Banana", "Maça", "Chocolate"}
	for key, value := range slice {
		fmt.Printf("Key: %v - Value: %v\n", key, value)
	}

	slice = append(slice, "Melancia")
	for key, value := range slice {
		fmt.Printf("Key: %v - Value: %v\n", key, value)
	}
}