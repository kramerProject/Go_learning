package main

import (
	"fmt"
)



func main() {
	//if  initializing; conditional;
	if hour := 0; !(hour > 1) {
		fmt.Println(hour)
	}

	if n := 10; n > 1 {
		fmt.Println(n)
	}

	if y:= -1; y > 10 {
		fmt.Printf("%v greater than 10", y)
	} else if y > 0 {
		fmt.Printf("%v smaller or equal 10 and greater than zero", y)
	} else {
		fmt.Printf("%v is a negative number", y)
	}



	if z:= 499; (z > 100 && z < 500) {
		fmt.Printf("%v is number is greater than 100 and less than 500", z)
	}
}