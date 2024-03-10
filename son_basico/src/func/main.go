package main

import "fmt"

func namedReturn(a string) (x string) {
	x = a
	return
}

func moreReturns(a, b string) (string, string) {
	return a, b
}

func variadicFunc(x ...int) int {
	var res int
	for _, v := range x {
		res += v
	}
	return res
}

func funcInsideFunc() func() int {
	x := 10

	return func() int {
		return x * x
	}
}

func main() {
	fmt.Println(namedReturn("Wesley"))

	nm, ls := moreReturns("Wes", "Saf")
	fmt.Println(nm)
	fmt.Println(ls)

	fmt.Println(variadicFunc(1, 2, 5, 10))

	z := 0

	add := func() int {
		z += 20
		return z
	}
	fmt.Println(add())

	fmt.Println(funcInsideFunc())
}
