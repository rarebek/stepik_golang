package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	h := a / 30
	m := 2 * (a % 30)

	fmt.Println("It is", h, "hours", m, "minutes.")

}
