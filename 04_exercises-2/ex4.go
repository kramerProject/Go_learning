package main

import (
	"fmt"
)


func main() {
	a := 150
	fmt.Printf("%d,  %b, %#x\n\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\n,  %b\n, %#x", b, b, b)
}

