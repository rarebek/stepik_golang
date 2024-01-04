package main

import (
	"fmt"
	"time"
)

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	result := make(chan int)
	sum := 0

	go func() {
		defer close(result)
		for {
			select {
			case a := <-arguments:
				sum += a

			case <-done:
				result <- sum
				return
			}
		}
	}()

	return result
}

func main() {
	arguments := make(chan int)
	done := make(chan struct{})

	go func() {
		arguments <- 1
		time.Sleep(time.Second)
		arguments <- 2
		time.Sleep(time.Second)
		arguments <- 3
		time.Sleep(time.Second)
		close(done)
	}()

	resultChan := calculator(arguments, done)

	result := <-resultChan
	fmt.Println("Result:", result)
}
