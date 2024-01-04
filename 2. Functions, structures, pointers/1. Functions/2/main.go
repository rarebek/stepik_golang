package main

import "fmt"

func main() {
	fmt.Print(minimumFromFour())
}
func minimumFromFour() int {
	var a int
	min := 999999999
	for i := 0; i < 4; i++ {
		fmt.Scan(&a)
		if a < min {
			min = a
		}
	}
	return min
}
