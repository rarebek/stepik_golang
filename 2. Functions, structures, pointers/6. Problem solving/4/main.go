package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	var res string
	fmt.Scan(&a)
	for i := 0; i < len(a); i++ {
		num, _ := strconv.Atoi(string(a[i]))
		res += strconv.Itoa(num * num)
	}

	fmt.Println(res)
}
