package main

import "fmt"

func main() {
	a := 2
	b := 3
	a1 := &a
	b1 := &b
	test(a1, b1)
}
func test(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Print(*x1, " ", *x2)
}
