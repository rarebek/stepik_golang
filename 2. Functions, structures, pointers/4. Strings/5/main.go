package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Scan(&x)
	res := ""
	for i := 0; i < len(x); i++ {
		if strings.Count(x, string(x[i])) <= 1 {
			res += string(x[i])
		}
	}
	fmt.Println(res)
}
