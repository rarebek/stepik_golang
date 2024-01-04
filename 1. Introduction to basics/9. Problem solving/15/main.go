package main

import "fmt"

func main() {
	var a []int
	var num, temp, res, con int
	fmt.Scan(&num)
	fmt.Scan(&con)
	for num > 0 {
		temp = num % 10
		a = append(a, temp)
		num /= 10
	}

	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != con {
			res = res*10 + a[i]
		}
	}
	fmt.Print(res)
}
