
package main

import (
	"fmt"
)

const (
	_ = iota
	KB = 1 << (iota * 10)
	MB
)
func main() {

	fmt.Printf("%b\n\n", KB)
	fmt.Printf("%b\n\n\n", MB)
}