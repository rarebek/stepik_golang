package main

import "fmt"

func main() {
	var a, f, s, t, res int
	fmt.Scan(&a)
	f = a / 100
	s = a / 10 % 10
	t = a % 10

	res = t*100 + s*10 + f
	fmt.Print(res)
}
