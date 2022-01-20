package main

import (
	"fmt"
)

func main() {
	// this means 2^16
	a := "a bunch of strinh $$777#"
	for _, v := range a {
		fmt.Printf("%b ---- %T ---- %#U ---- %#x\n", v, v, v, v)
	}
	fmt.Printf("---------------\n")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%b ---- %T ---- %#U ---- %#x\n", a[i], a[i], a[i], a[i])
	}
}

// proximos fundamentos da programação 6