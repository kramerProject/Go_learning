package main

import (
	"fmt"
)


func main() {
	literal := `this is string \n literal       `
	raw := "this is not \n\n continue in the next line"
	fmt.Printf("%v %v", literal, raw)
}