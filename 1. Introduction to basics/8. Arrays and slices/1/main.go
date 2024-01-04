package main

import "fmt"

func main() {
	var a, x, y int
	var workArray [10]uint8
	var temp uint8
	for i := 0; i < 10; i++ {
		fmt.Scan(&a)
		workArray[i] = uint8(a)
	}
	for i := 0; i < 3; i++ {
		fmt.Scan(&x, &y)
		temp = workArray[x]
		workArray[x] = workArray[y]
		workArray[y] = temp
	}
	for j := 0; j < 10; j++ {
		fmt.Print(workArray[j], " ")
	}
}
