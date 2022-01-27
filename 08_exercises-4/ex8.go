package main

import (
	"fmt"
)

func main() {
	people := map[string][]string{
		"Kramer": []string{"Surf", "Cook"},
		"Jo": []string{"Play", "Soccer"},
		"Nin": []string{"Sing", "Dance"},
	}

	
	for key, hobbies := range people {
		fmt.Printf("Person: %v\n", key)
		for _, value := range hobbies{
			fmt.Printf("\tHobby: %v\n", value)
		}
	}
}