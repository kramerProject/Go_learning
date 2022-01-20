package main

import (
	"fmt"
)



func main() {
	age := 0
	for x := 1983; x <= 2022; x++ {
		fmt.Printf("Year: %v Age: %v\n", x, age)
		age++
	}
}