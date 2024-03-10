package main

import "fmt"

func xpto(a *int) int {
	*a = 100
	return *a
}

func main() {

	b := 1
	fmt.Println(xpto(&b))

	fmt.Println(b)
	//guarde na variavel, aponta pra memoria da máquina
	x := 11
	fmt.Println(&x)

	y := &x

	fmt.Println(y)
	fmt.Println(*y)

	*y = 20
	fmt.Println(x)

	var z *int = &x

	fmt.Println(z)
}
