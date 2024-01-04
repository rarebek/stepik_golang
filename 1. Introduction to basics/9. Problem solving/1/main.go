package main

import "fmt"

func main() {
	var a, f, s, t int
	fmt.Scan(&a)
	f = a / 100
	s = a / 10 % 10
	t = a % 10
	fmt.Print(f + s + t)

}
