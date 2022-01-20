package main

import (
	"fmt"
)

const (
	_ = iota + 2022
	b
	c
	d
	e
)

func main() {
	fmt.Printf("%v %v %v %v", b, c, d, e)
}