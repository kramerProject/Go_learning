package main

import (
	"fmt"
)

func main() {
	// operador curto := apenas dentro de funções
	x := 2
	y := "hello"
	z := 10 == 10
	w := 10 > 20
	fmt.Printf("x: %v and type: %T\n", x, x)
	fmt.Printf("x: %v and type: %T\n", y, y)
	fmt.Printf("x: %v and type: %T\n", z, z)
	fmt.Printf("x: %v and type: %T\n", w, w)
	z = 10 > 100
	fmt.Println(z)
}