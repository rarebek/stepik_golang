package main

import "fmt"

func main() {
	var a int
	res := 2
	fmt.Scan(&a)
	fmt.Print(1, " ")
	for {
		if res <= a {
			fmt.Print(res, " ")
			res *= 2
		} else {
			break
		}
	}
}
