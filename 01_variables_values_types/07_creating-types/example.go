package main

import "fmt"

type createdvar int
var b createdvar
func main() {
	x := 100
	fmt.Printf("value: %v and type: %T\n", b, b)
	fmt.Printf("value: %v and type: %T\n", x, x)
}