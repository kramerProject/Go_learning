package main

import "fmt"

func main() {
	a := 1

	if a >= 10 {
		fmt.Println("a>10")
	} else if a > 5 {
		fmt.Println("a>5")
	} else {
		fmt.Println("else")
	}

	b := false
	if x := "Cool"; b {
		fmt.Printf("cool %v", x)
	}

}
