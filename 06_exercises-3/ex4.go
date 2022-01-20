package main

import (
	"fmt"
)



func main() {
	age := 0
	year := 1983
	for {
		if year > 2022 {
			break
		}
		fmt.Printf("Year: %v Age: %v\n", year, age)
		age++
		year++
		
	}
}