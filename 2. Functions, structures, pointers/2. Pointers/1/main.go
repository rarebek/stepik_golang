package main

import "fmt"

func main() {
	a := 2
	a1 := &a
	b := 2
	b1 := &b
	test(a1, b1)
}
func test(x1 *int, x2 *int) {
	fmt.Print(*x1 * *x2)
}
