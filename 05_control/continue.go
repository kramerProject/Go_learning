package main

import (
	"fmt"
)



func main() {
	// continue stops the iteration inside the loop and continue to next.
	for a := 0; a < 10; a++ {
		if a % 2 != 0 {
			continue
		}
		fmt.Printf("Esse número %v é par \n", a)
	}
}