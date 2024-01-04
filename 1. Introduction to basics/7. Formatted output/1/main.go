package main

import (
	"fmt"
)

func main() {
	var inputNumber float64
	fmt.Scan(&inputNumber)

	if inputNumber <= 0 {
		fmt.Printf("число %2.2f не подходит\n", inputNumber)
	} else if inputNumber > 10000 {
		fmt.Printf("%e\n", inputNumber)
	} else {
		squared := inputNumber * inputNumber
		fmt.Printf("%.4f\n", squared)
	}
}
