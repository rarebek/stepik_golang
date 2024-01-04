package main

import "fmt"

func main() {

	var a, t, cnt int
	fmt.Scan(&a)
	var arr []int

	for i := 0; i < a; i++ {
		fmt.Scan(&t)
		arr = append(arr, t)
	}

	for i := 0; i < a; i++ {
		if arr[i] >= 0 {
			cnt += 1
		}
	}
	fmt.Print(cnt)

}
