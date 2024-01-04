package main

import "fmt"

func main() {
	var a, b, sum, res, temp int
	fmt.Scan(&a)
	for {
		b = a % 10
		sum += b
		a = a / 10

		if a <= 0 {
			break
		}
	}
	for {
		temp = sum % 10
		res += temp
		sum = sum / 10

		if sum <= 0 {
			break
		}
	}

	fmt.Print(res)

}
