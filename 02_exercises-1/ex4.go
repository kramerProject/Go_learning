package main

import (
	"fmt"
)

type myvar int

var x myvar

func main() {
	fmt.Printf("value: %v type: %T\n", x, x)
	x = 42
	fmt.Printf("%v and %T", x, x)
}