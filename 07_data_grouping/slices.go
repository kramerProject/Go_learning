package main

import (
	"fmt"
)



func main() {
	flavors := []string{"chocolate", "morango", "flocos"}
	chocolate := flavors[0]
	fmt.Println(chocolate)
	strawberry := flavors[1]
	flakes := flavors[2]
	fmt.Printf("%v -- %v -- %v\n", chocolate, strawberry, flakes)
	fmt.Println(flavors[:])
	// excluding strawberry
	flavors = append(flavors[0:1], flavors[2:]...)
	fmt.Println(flavors)
}