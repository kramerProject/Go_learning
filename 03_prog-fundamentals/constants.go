
package main

import (
	"fmt"
)

// a defined as int
var a = 10

// const types are not defined when declared
const b = 10
func main() {
	// can attritbute b to c but not a to c
	c := 10
	c = b
	fmt.Printf("value: %v type: %T", c, c)
}