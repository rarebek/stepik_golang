package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	// package main уже объявлен, нужные импорты уже есть
	x, err := divide(a, b)
	if err != nil {
		fmt.Println("ошибка")
	} else {
		fmt.Println(x)
	}
}
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("ошибка")
	}
	return a / b, nil
}
