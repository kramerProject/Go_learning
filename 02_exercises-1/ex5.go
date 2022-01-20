package main

import (
	"fmt"
)

type myvar int

var x myvar
var y int

func main() {
	fmt.Printf("value: %v type: %T\n", x, x)
	x = 42
	fmt.Printf("value of x is %v\n", x)
	y = int(x)
	fmt.Printf("y is ----> %v and %T", y, y)
}