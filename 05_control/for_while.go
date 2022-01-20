package main

import (
	"fmt"
)



func main() {
	// incialização; condição; final do loop
	// exemplo for como while
	hour := 0
	for hour <= 12 {
		fmt.Println("Horas: ", hour)
		for minutes := 0; minutes <= 60; minutes++ {
			fmt.Println("Minutes", minutes)
		}
		hour++
	}

	// Com break
	number := 0
	for {
		if number == 10 {
			fmt.Println("Stopped")
			break
		} else {
			fmt.Println("keep going")
			number++
		}
	}
}