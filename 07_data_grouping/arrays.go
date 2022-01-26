package main

import (
	"fmt"
)


var x [5]int

func main() {
	x[0], x[1], x[2], x[3], x[4] = 10, 20, 30, 40, 60
	fmt.Println(cap(x))
}