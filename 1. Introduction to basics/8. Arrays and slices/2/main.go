package main

import "fmt"

func main() {
	var a, b int
	var myslice []int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		myslice = append(myslice, b)
	}

	fmt.Print(myslice[3])
}
