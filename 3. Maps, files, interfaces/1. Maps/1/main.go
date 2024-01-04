package main

import "fmt"

func main() {
	var a int
	var dict = make(map[int]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Scan(&a)
		if m, ok := dict[a]; ok {
			fmt.Print(m, " ")
		} else {
			dict[a] = work(a)
			fmt.Print(dict[a], " ")
		}
	}
}
func work(n int) int {
	if n >= 4 {
		return n + 1
	} else {
		return n - 1
	}
}
