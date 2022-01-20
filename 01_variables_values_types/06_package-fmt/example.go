package main

import "fmt"

func main() {
	// interpreted
	x := "Good morning \n how are yoou \n"
	// literal
	y := `this is a literal string this \n is not a enter`
	
	// Sprint does not print anything, it returns a string
	z := fmt.Sprint(x, y)
	word := z
	fmt.Println(word)
}