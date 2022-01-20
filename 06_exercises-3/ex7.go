package main

import (
	"fmt"
)



func main() {
	status := "D"
	if status == "Done" {
		fmt.Printf("Congrats")
	} else if status == "WIP" {
		fmt.Printf("WIP")
	} else {
		fmt.Printf("Not started")
	}
}