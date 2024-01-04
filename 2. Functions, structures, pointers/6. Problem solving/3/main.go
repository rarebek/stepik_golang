package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	max := 0
	b := []string{}
	b = strings.Split(a, "")
	c := []int{}
	for i := 0; i < len(b); i++ {
		num, _ := strconv.Atoi(b[i])
		c = append(c, num)
	}
	for i := 0; i < len(c); i++ {
		if c[i] > max {
			max = c[i]
		}
	}
	fmt.Println(max)
}
