package main

import "fmt"

func main() {
	var a, e, min, cnt int
	var sl []int

	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&e)
		sl = append(sl, e)
	}
	min = sl[0]

	for i := 0; i < a; i++ {
		if sl[i] < min {
			min = sl[i]
		}
	}

	for i := 0; i < a; i++ {
		if min == sl[i] {
			cnt += 1
		}
	}

	fmt.Print(cnt)

}
