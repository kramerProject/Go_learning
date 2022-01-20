package main

import (
	"fmt"
)



func main() {
	// incialização; condição; final do loop
	for hour := 0; hour <= 12; hour++ {
		fmt.Println("Horas: ", hour)
		for minutes := 0; minutes <= 60; minutes++ {
			fmt.Println("Minutes", minutes)
		}
	}
}