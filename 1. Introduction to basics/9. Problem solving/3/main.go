package main

import "fmt"

func main() {
	var a, h, m, temp int
	fmt.Scan(&a)
	h = a / 3600
	temp = a % 3600
	m = temp / 60
	fmt.Printf("It is %d hours %d minutes.", h, m)

}
