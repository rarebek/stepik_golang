package main

import "fmt"

func main() {
	var a, b int
	var checker bool
	checker = false
	fmt.Scan(&a)
	fmt.Scan(&b)

	for i := b; i >= a; i-- {
		if i%7 == 0 {
			fmt.Print(i)
			checker = true
			break
		}
	}
	if checker == false {
		fmt.Print("NO")
	}

}
