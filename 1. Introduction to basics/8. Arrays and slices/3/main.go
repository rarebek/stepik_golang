package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	var max int = array[0]
	for i := 0; i < 5; i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	fmt.Print(max)

}
