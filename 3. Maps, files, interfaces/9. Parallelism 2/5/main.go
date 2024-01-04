package main

import (
	"sync"
)

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	wg := new(sync.WaitGroup)
	ar1 := make([]int, n)
	ar2 := make([]int, n)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				val := fn(x1)
				ar1[i] = val
			}(i)
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				val := fn(x2)
				ar2[i] = val
			}(i)
		}
	}()

	go func() {
		wg.Wait()
		for i := 0; i < n; i++ {
			out <- ar1[i] + ar2[i]
		}
	}()
}
