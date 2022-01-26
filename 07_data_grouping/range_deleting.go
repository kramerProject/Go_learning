package main

import (
	"fmt"
)




func main() {
	people := map[int]string{
		1: "Roger",
		2: "Kramer",
		3: "Joao",
	}

	for k, v := range people {
		fmt.Printf("Key: %v - Value: %v\n", k, v)
	}
	delete(people, 2)
	fmt.Println(people)
}