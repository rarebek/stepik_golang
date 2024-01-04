package main

import "fmt"

func main() {
	var n, c, d int
	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&d)

	var i int
	for i = 0; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Print(i)
			break
		}
	}

}
