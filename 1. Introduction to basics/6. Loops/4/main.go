package main

import "fmt"

func main() {
	var a int
	max := 0
	cnt := 0
	for {
		fmt.Scan(&a)
		if a != 0 {
			if a > max {
				cnt = 1
				max = a
			} else if a == max {
				cnt += 1
			}
		} else {
			break
		}
	}
	fmt.Print(cnt)
}
