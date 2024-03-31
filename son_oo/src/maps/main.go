package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["a"] = 10
	m["b"] = 11
	m["c"] = 30
	delete(m, "c")
	fmt.Println(m)

	_, ok := m["c"]
	if !ok {
		fmt.Println(ok)
		println(ok)
	}

	// var x = map[string]int{}
	// fmt.Println(x)

	newMap := map[string]int{"x": 5, "y": 10}
	fmt.Println(newMap)

	if val, exists := newMap["c"]; exists {
		fmt.Println(val)
	} else {
		fmt.Println(val)
	}
}
