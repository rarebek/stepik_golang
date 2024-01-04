package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	first_three_sum := (a / 100000) + (a / 10000 % 10) + (a / 1000 % 10)
	second_three_sum := a%10 + (a%100)/10 + (a%1000)/100
	if first_three_sum == second_three_sum {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
