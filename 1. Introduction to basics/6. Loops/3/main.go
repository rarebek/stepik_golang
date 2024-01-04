package main

import "fmt"

func main() {
	var i, a, sum, l int
	fmt.Scan(&l)
	for i = 0; i < l; i++ {
		fmt.Scan(&a)
		if a < 100 && a > 9 && a%8 == 0 {
			sum += a
		}
	}
	fmt.Print(sum)
}
