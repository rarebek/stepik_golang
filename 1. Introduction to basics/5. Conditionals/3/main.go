package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a/10000 != 0 {
		fmt.Print(a / 10000)
	} else if a/1000 != 0 {
		fmt.Print(a / 1000)
	} else if a/100 != 0 {
		fmt.Print(a / 100)
	} else if a/10 != 0 {
		fmt.Print(a / 10)
	} else {
		fmt.Print(a)
	}
}
