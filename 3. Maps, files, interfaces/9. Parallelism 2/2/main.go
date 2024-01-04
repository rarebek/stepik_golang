package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			work()
			wg.Done()
		}()
	}
	wg.Wait()
}

func work() {
	fmt.Println("work")
}
