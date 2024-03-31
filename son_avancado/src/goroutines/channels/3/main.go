package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	channel := make(chan int)
	wg.Add(2)

	go func() {

		for i := 10; i < 20; i++ {
			channel <- i
		}
		wg.Done()
	}()

	go func() {

		for i := 21; i < 30; i++ {
			channel <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(channel)
	}()

	for number := range channel {
		fmt.Println(number)
	}

}
