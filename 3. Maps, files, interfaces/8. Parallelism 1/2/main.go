package main

import (
	"fmt"
	"sync"
)

func task2(ch chan<- string, n string) {
	for i := 0; i < 5; i++ {
		ch <- n + " "
	}
	close(ch)
}

func main() {
	resultChannel := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			task2(resultChannel, fmt.Sprintf("Routine %d", id))
		}(i)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	for result := range resultChannel {
		fmt.Print(result)
	}
}
