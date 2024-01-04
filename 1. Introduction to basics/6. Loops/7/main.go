package main

import "fmt"

func main() {
	var x, p, y, i int
	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)
	curX := x
	i = 0

	for {
		curX = curX + curX*p/100
		i += 1
		if curX >= y {
			break
		}
	}
	fmt.Print(i)

}
