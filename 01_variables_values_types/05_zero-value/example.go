package main

import "fmt"

// Declare
var a int
var b float64
var c string
var d bool

func main() {
	// Each type has an initial value
	fmt.Printf("value: %v and type: %T\n", a, a)
	fmt.Printf("value: %v and type: %T\n", b, b)
	fmt.Printf("value: %v and type: %T\n", c, c)
	fmt.Printf("value: %v and type: %T\n", d, d)
	// intializing
	a = 10
	fmt.Println(a)
	// attributtion
	a = 200
	fmt.Println(a)
}

