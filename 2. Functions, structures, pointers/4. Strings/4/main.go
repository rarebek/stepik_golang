package main

import (
	"fmt"
)

func main() {
	var x string
	fmt.Scan(&x)
	res := ""
	for i := 1; i < len(x); i += 2 {
		res += string(x[i])
	}
	fmt.Println(res)
}
