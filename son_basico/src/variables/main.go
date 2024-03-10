package main

import "fmt"

var b int
var c, d string = "Var C", "Var D"

const (
	myConst = "myConst"
)

func main() {
	a := 10
	b = 22
	fmt.Printf("a: %v \n b: %v \n", a, b)
	fmt.Printf("c: %v \n d: %v \n", c, d)
	fmt.Printf("const: %v \n", myConst)
}
