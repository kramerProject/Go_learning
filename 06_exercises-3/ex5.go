package main

import (
	"fmt"
)



func main() {
	for number := 10; number < 100; number++{
		fmt.Printf("(%v / 4) Remainer: %v \n", number, number % 4)
	}
}