package main

import "fmt"

func main() {

	var a, t int
	fmt.Scan(&a)
	var arr []int

	for i := 0; i < a; i++ {
		fmt.Scan(&t)
		arr = append(arr, t)
	}

	for i := 0; i < a; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}

}
