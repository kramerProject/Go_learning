package main

import "fmt"

var y int

func main() {
	// once you declare a type you can not change it
	y = 10
	fmt.Println(someother(y))
}

func someother(x int) int {
	return x + y
}