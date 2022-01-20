
package main

import (
	"fmt"
)

var a bool

func main() {
	// Initialized with zero values
	fmt.Println(a)
	a = true
	fmt.Printf("a: %v and type: %T\n", a, a)

	b := 10 == 10
	c := 10 > 1
	d := 10 == 100
	fmt.Printf("b: %v and type: %T\n", b, b)
	fmt.Printf("c: %v and type: %T\n", c, c)
	fmt.Printf("d: %v and type: %T\n", d, d)
}