package main

import "fmt"

func main() {
	channel := make(chan int)
	ok := make(chan bool)

	go func() {
		for i := 10; i < 20; i++ {
			channel <- i
		}
		ok <- true
	}()

	go func() {
		for i := 10; i < 20; i++ {
			channel <- i
		}
		ok <- true
	}()

	go func() {
		<-ok
		<-ok
		close(channel)
	}()

	for number := range channel {
		fmt.Println("Number", number)
	}
}
