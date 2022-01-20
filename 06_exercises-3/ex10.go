package main

import (
	"fmt"
)



func main() {
	x := -1
	if (x > 20 && x < 50) {
		fmt.Printf("1 - True %v", x)
	}
	if (x > 1 && x == 20) {
		fmt.Printf("2 - False %v", x)
	}
	if (x == 1 || x == 15) {
		fmt.Printf("3 - True %v", x)
	}
	if (x == 20000 || x == 10) {
		fmt.Printf("4 - True %v", x)
	}
	if x != -3 {
		fmt.Printf("5 - True %v", x)
	}
	
}