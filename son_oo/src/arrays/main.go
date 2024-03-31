package main

import "fmt"

func main() {
	var x [10]int

	fmt.Println(len(x))
	x[0] = 11

	fmt.Println(x)

	var y [10]string
	fmt.Println(y)

	arr := [3]int{1, 2, 3}
	fmt.Println(arr[1])
}
