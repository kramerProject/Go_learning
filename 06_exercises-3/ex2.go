package main

import (
	"fmt"
)



func main() {

	for x := 65; x <= 90; x++ {
		fmt.Printf("Number: %v\n", x)
		for y := 1; y <= 3; y++ {
			fmt.Printf("\tUnicode is: %#U\n", x)
		}
	}
}