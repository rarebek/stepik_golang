package main

import "fmt"

func task(channel chan int, n int) {
	channel <- n + 1
}

func main() {
	resultChannel := make(chan int)

	go task(resultChannel, 5)

	result := <-resultChannel

	fmt.Println(result)
}
