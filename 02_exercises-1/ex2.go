
package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	// Initialized with zero values
	fmt.Println(x, y, z)
	fmt.Printf("x: %v and type: %T\n", x, x)
	fmt.Printf("x: %v and type: %T\n", y, y)
	fmt.Printf("x: %v and type: %T\n", z, z)

	
}

