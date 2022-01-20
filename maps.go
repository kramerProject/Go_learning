package main

import (
	"fmt"
	"sort"
)

func main() {
   second()
}

func first() {
	ages := map[string]int{
		"alice": 22,
		"charlie": 23,
	}
	for name, age := range ages {
		fmt.Println(name, age)
	}
}

func second() bool{
	ages := map[string]int{
		"alice": 22,
		"charlie": 23,
	}
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("Nome: %s - Idade: %d\n", name, ages[name])
	}
	kramer, ok := ages["kramer"]
	if !ok {
		fmt.Printf("-----%d", kramer)
		return false
	}
	return true
}



