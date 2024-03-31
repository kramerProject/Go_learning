package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)

	go runProcess(channel)

	fmt.Println(<-channel)
	// go func() {
	// 	channel <- 10
	// }()

}

func runProcess(ch chan int) {
	ch <- 10

}
