package main

import (
	"fmt"
	"time"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	result := make(chan int)

	go func() {
		defer close(result)

		select {
		case v := <-firstChan:
			result <- v * v

		case v := <-secondChan:
			result <- v * 3

		case <-stopChan:
			return

		default:
			return
		}
	}()

	return result
}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	go func() {
		firstChan <- 2
		time.Sleep(time.Second)
		secondChan <- 5
		time.Sleep(time.Second)
		close(stopChan)
	}()

	resultChan := calculator(firstChan, secondChan, stopChan)

	result := <-resultChan
	fmt.Println("Result:", result)
}
