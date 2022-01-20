package main

import (
	"fmt"
)



func main() {
	// incialização; condição; final do loop
	for a := 33; a <= 122; a++ {
		fmt.Printf("Number: %v String: %v \n", a, string(a))
	}
}