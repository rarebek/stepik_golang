package main

import "fmt"

func main() {
	fmt.Println(fibonacci(50))
}
func fibonacci(n int) int {
	var a, b int
	index := 2
	a, b = 0, 1
	for index <= n {
		a, b = b, a+b
		index++
	}
	return b

}
