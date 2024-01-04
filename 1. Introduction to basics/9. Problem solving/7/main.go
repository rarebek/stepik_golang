package main

import "fmt"

func main() {
	var a, t, cnt int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&t)
		if t == 0 {
			cnt += 1
		}
	}
	fmt.Print(cnt)
}
