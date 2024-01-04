package main

import "fmt"

func main() {
	fmt.Println(convert(100))
}
func convert(a int64) uint16 {
	var b uint16
	b = uint16(a)
	return b
}
