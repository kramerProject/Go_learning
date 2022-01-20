
package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Initialized with zero values
	a := "e"
	b := "é"
	abyte, _ := fmt.Println(a)
	bbyte, _ := fmt.Println(b)
	fmt.Println(abyte, bbyte)

	c := []byte(a)
	d := []byte(b)
	fmt.Println(c, d)

	cbyte, _ := fmt.Println(c)
	dbyte, _ := fmt.Println(d)
	fmt.Println(cbyte, dbyte)

	fmt.Printf("OS: %v and type: %v\n", runtime.GOOS, runtime.GOARCH)
}