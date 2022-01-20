package main

import (
	"fmt"
)



func main() {
	x := 50
	switch {
		case x < 10:
			fmt.Println(x)
		case x < 20:
			fmt.Println(x)
		case x == 50:
			fmt.Println("case 3")
	}

	climate := "R"
	switch climate {
		case "Raining":
			fmt.Println("Raining")
		case "Cloudy":
			fmt.Println(x)
		default:
			fmt.Println("Sunny")
	}

	y := 50
	switch y {
		case 2, 50:
			fmt.Println(y)
			// it will execute next case as well
			fallthrough
		case 3, 4:
			fmt.Println(y)
		default:
			fmt.Println("default")
	}
}