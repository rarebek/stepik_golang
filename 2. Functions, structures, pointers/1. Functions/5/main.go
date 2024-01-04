package main

import "fmt"

func main() {
	fmt.Println(sumInt(10, 20, 30, 40))
}

func sumInt(nums ...int) (int, int) {
	cnt := 0
	sum := 0
	for _, num := range nums {
		cnt++
		sum += num
	}
	return cnt, sum
}
