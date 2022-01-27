package main

import (
	"fmt"
)

func main() {
	people := [][]string{
		[]string{"Kramer", "Silva", "Surf"},
		[]string{"Jo", "Checoli", "Play"},
		[]string{"Nin", "Checoli", "Sing"},
	}

	for _, person := range people {
		fmt.Printf("person: %v lastname: %v hobby: %v\n", person[0], person[1], person[2])
	}
	fmt.Println(people)
}