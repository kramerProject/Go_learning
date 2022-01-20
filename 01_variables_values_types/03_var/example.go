package main

import "fmt"

var y = 100

func main() {
	z := 10
	fmt.Println(someother(z))
}

func someother(x int) int {
	return y + x
}