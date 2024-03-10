package main

import (
	"fmt"
)

func main() {
	// incialização; condição; final do loop
	// exemplo for como while
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	y := 0
	for {
		y++
		if y == 100 {
			fmt.Println(y)
			break
		}
	}
}
