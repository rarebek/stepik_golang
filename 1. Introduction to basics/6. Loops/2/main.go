package main

import "fmt"

func main() {
	var a, b int
	var i int
	sum := 0
	fmt.Scan(&a, &b)
	for i = a; i <= b; i++ {
		sum += i
	}
	fmt.Print(sum)

}
