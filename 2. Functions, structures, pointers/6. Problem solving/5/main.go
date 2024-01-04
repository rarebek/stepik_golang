package main

import (
	"fmt"
	"math"
)

var k float64 = 10
var p float64 = 10
var v float64 = 10

func main() {
	fmt.Println(T())
	fmt.Println(W())
	fmt.Println(M())
}

func T() float64 {
	w := W()
	result := 6 / w
	return result
}

func W() float64 {
	m := M()
	result := math.Sqrt(k / m)
	return result
}

func M() float64 {
	result := p * v
	return result
}
