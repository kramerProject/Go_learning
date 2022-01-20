package main

import (
	"fmt"
)

func main() {
	a := 10 == 100
	b := 10 != 10
	c := 10 <= 10
	d := 100 < 1000
	e := 50 >= 49
	f := 100 > 101
	
	fmt.Printf("%v, %v, %v, %v, %v,", a, b, c, d, e, f)
}