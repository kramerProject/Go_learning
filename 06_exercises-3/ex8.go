package main

import (
	"fmt"
)



func main() {
	x := 123
	switch {
		case x == 2:
			fmt.Printf("Case %v", x)
		case x == 3:
			fmt.Printf("Case %v", x)
		case x == 4:
			fmt.Printf("Case %v", x)
		default:
			fmt.Printf("Case Nothing %v", x)
	}
}