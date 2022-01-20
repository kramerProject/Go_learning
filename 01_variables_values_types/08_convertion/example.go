package main

import "fmt"

type createdvar float64

var a createdvar = 200
var b createdvar = 300

func main() {
	c := 100
	d := 100
	 // convertion possible because the created var was an int, also works if it was a float64
	c = int(a)
	d = int(b)
	fmt.Printf("value: %v and type: %T\n", c, c)
	fmt.Printf("value: %v and type: %T\n", d, d)
}