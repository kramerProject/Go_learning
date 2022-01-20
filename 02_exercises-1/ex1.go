package main

import (
	"fmt"
)

func main() {
	x := 42
	y := "James bond"
	z := true
	fmt.Println(x, y, z)
	fmt.Printf("x: %v and type: %T\n", x, x)
	fmt.Printf("x: %v and type: %T\n", y, y)
	fmt.Printf("x: %v and type: %T\n", z, z)
	
}