
package main

import (
	"fmt"
)

func main() {
	// this means 2^16
	var a uint16
	a = 65534
	fmt.Println(a)
	a++
	fmt.Printf("%v and %T\n", a, a)
	a++
	fmt.Printf("%v and %T\n", a, a)
}