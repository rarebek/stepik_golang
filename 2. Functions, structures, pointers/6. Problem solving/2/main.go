package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	b := strings.Split(a, "")
	a = strings.Join(b, "*")
	fmt.Println(a)
}
